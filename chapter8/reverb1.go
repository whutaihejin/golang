package main

import (
	"net"
	"io"
	"log"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

func main() {
	listener, err := net.Listen(TCP, "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	io.Copy(c, c)
	c.Close()
}