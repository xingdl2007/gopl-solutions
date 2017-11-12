// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// simple FTP server
package main

import (
	"net"
	"log"
	"bufio"
	"io"
	"os/exec"
	"strings"
	"os"
)

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func exeCmd(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":
			exeCmd(c, cmds[0], cmds[1:]...)
		case "cd":
			// one client change work dir, all other client will see the same effect
			if err := os.Chdir(cmds[1]); err != nil {
				log.Print(err)
			}
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("file %s: %v", cmds[1], err)
				continue
			}
			mustCopy(c, file)
		case "close":
			return
		default:
			help := "ls: list content\ncd: change directory\nget: get file content\nclose: close connection\n"
			mustCopy(c, strings.NewReader(help))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
