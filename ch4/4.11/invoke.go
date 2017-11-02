// call vim/code in go
package main

import (
	"os/exec"
	"log"
	"fmt"
	"io"
	"bufio"
	"os"
	"io/ioutil"
)

func logger(pipe io.ReadCloser) {
	reader := bufio.NewReader(pipe)

	for {
		output, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Print(output)
	}
}

func main() {
	file, err := ioutil.TempFile(os.TempDir(), "trailsman-")
	if err != nil {
		log.Fatal(err)
	}
	content := "This is a test"
	file.WriteString(content)

	log.Println(file.Name())
	file.Close()

	// template file
	cmd := exec.Command("vim", file.Name())

	//------------------------------------------------
	// You can find the reason when bad things happen
	// a stderrPipe to get error message
	//pipe, err := cmd.StderrPipe()

	// excellent goroutine
	//go logger(pipe)
	//------------------------------------------------

	//----------------------------
	// solution
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	//----------------------------

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}
