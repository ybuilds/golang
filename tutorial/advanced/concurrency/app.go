package main

import (
	"fmt"
	"time"
)

func greet(message string, channel chan bool) {
	fmt.Println(message)
	channel <- true
	close(channel)
}

func slowGreet(message string, channel chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println(message)
	channel <- true
	close(channel)
}

func main() {
	done := make(chan bool)
	dones := make([]chan bool, 2)

	dones[0] = make(chan bool)
	dones[1] = make(chan bool)

	go slowGreet("Hello slow world!", done)
	go greet("Hello world!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
	}
}
