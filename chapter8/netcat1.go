package main

import (
	"io"
	"log"
	"net"
	"os"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

func main() {
	conn, err := net.Dial(TCP, "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}