package main

import (
	"net"
	"io"
	"time"
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
		handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}