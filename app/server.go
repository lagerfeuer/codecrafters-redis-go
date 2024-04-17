package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	fmt.Println("$$$ handle $$$")

	for {
		buffer := make([]byte, 128)

		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Could not read from connection: ", err.Error())
			return
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}
