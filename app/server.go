package main

import (
	"fmt"
	"net"
	"os"
)

func handle(conn net.Conn) {
	var buffer []byte

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Could not read from connection: ", err.Error())
	}

	conn.Write([]byte("+PONG\r\n"))
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379: ", err.Error())
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handle(conn)
}
