package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, error := conn.Read(buffer)

	if error != nil {
		log.Println(error)
	}

	fmt.Printf("Client: %s\n", buffer[:n])
	fmt.Fprintf(conn, "echo - %s\n", string(buffer[:n]))
}

func handleConnectionAsync(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(time.Second))
		buffer := make([]byte, 1024)

		n, error := conn.Read(buffer)

		if error != nil {
			if netError, ok := error.(net.Error); ok && netError.Timeout() {
				continue
			} else {
				log.Println(error)
				break
			}
		}

		fmt.Printf("Client: %s\n", buffer[:n])

		conn.SetWriteDeadline(time.Now().Add(time.Second))
		fmt.Fprintf(conn, "echo - %s\n", string(buffer[:n]))
	}
}

func tcpServer() {
	listener, error := net.Listen("tcp", "localhost:8080")

	if error != nil {
		log.Fatalln(error)
	}

	defer listener.Close()

	for {
		conn, error := listener.Accept()

		if error != nil {
			log.Println(error)
		}

		go handleConnectionAsync(conn)
	}
}
