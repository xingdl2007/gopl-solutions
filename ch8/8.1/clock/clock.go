// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
)

var port = flag.String("port", "8000", "port number")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
