package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lost() bool {
	probabilty := 35
	return rand.Int()%100 < probabilty
}

func stopWait() {
	var frame int

	fmt.Print("Enter frame size: ")
	fmt.Scan(&frame)

	i := 0
	for i < frame {
		fmt.Printf("sending frame[%d]...\n", i)

		if lost() {
			fmt.Printf("frame[%d] lost\n", i)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Printf("received frame[%d]\n", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}
