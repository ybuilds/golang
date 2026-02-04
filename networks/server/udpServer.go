package main

import (
	"fmt"
	"log"
	"net"
)

func udpServer() {
	address, error := net.ResolveUDPAddr("udp4", "localhost:8080")

	if error != nil {
		log.Fatalln(error)
	}

	conn, error := net.ListenUDP("udp4", address)

	if error != nil {
		log.Println(error)
	}

	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		n, clientAddress, error := conn.ReadFromUDP(buffer)

		if error != nil {
			log.Println(error)
			continue
		}

		fmt.Printf("Received from %s: %s\n", clientAddress, string(buffer[:n]))

		_, error = conn.WriteToUDP([]byte("Received\n"), clientAddress)

		if error != nil {
			log.Println(error)
		}
	}
}
