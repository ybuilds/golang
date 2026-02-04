package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func udpClient() {
	conn, error := net.Dial("udp4", "localhost:8080")

	if error != nil {
		log.Fatalln(error)
	}

	defer conn.Close()

	message := "Hello to UDP!"

	_, error = conn.Write([]byte(message))

	if error != nil {
		log.Println(error)
	}

	fmt.Printf("Sent: %s\n", message)

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	buffer := make([]byte, 1024)

	n, error := conn.Read(buffer)

	if error != nil {
		log.Println(error)
		return
	}

	fmt.Printf("Server: %s\n", string(buffer[:n]))
}
