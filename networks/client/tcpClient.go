package main

import (
	"fmt"
	"log"
	"net"
)

func tcpClient() {
	conn, error := net.Dial("tcp", "localhost:8080")

	if error != nil {
		log.Fatalln(error)
	}

	fmt.Fprintf(conn, "Hello!")

	buffer := make([]byte, 1024)
	n, error := conn.Read(buffer)

	if error != nil {
		log.Println(error)
	}

	fmt.Printf("Server: %s\n", string(buffer[:n]))

	conn.Close()
}
