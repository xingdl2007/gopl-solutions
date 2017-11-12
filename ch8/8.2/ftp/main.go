package main

import (
	"net"
	"log"
	"io"
	"os"
)

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
