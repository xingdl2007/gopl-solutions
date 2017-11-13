// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

// ref: https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
// aggregate chan array to one share chan

// aggregated message
type Message struct {
	filesize int64
	id       int
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	size := len(roots)
	if size == 0 {
		roots = []string{"."}
		size = 1
	}

	//!+
	var message = make(chan Message)
	var wg sync.WaitGroup

	// Traverse each root of the file tree in parallel.
	var fileSizes = make([]chan int64, size)

	var n = make([]*sync.WaitGroup, size)
	// separate statistic
	for i := 0; i < size; i++ {
		fileSizes[i] = make(chan int64)
		n[i] = new(sync.WaitGroup)
	}

	for i, root := range roots {
		n[i].Add(1)
		go walkDir(root, n[i], fileSizes[i])

		go func(i int) {
			n[i].Wait()
			close(fileSizes[i])
		}(i)

		// gather all filesize message, with extra id info
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var msg Message
			for fs := range fileSizes[i] {
				msg.filesize = fs
				msg.id = i
				message <- msg
			}
		}(i)
	}

	// message closer
	go func() {
		wg.Wait()
		close(message)
	}()

	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes []int64
	nfiles = make([]int64, size)
	nbytes = make([]int64, size)

	// totals
	var tfiles, tbytes int64

loop:
	for {
		// separate
		select {
		case msg, ok := <-message:
			if !ok {
				break loop
			}
			nfiles[msg.id]++
			nbytes[msg.id] += msg.filesize

			tfiles++
			tbytes += msg.filesize
		case <-tick:
			printAllDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(tfiles, tbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printAllDiskUsage(file []string, nfiles, nbytes []int64) {
	for i := range nfiles {
		fmt.Printf("%s: %d files  %.1f GB; ", file[i], nfiles[i], float64(nbytes[i])/1e9)
	}
	fmt.Println()
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
