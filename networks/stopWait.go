package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swSender() {
	for s := 0; s < 10; s++ {
		for {
			fmt.Printf("sending frame %d\n", s)
			frame = &s

			start := time.Now()
			for time.Since(start) < time.Second {
				if ack != nil && *ack == s {
					ack = nil
					goto next
				}
			}
			fmt.Printf("no ack. retransmit\n")
		}
	next:
	}
}

func swReceiver() {
	exp := 0
	for exp < 10 {
		if frame == nil {
			continue
		}

		if unreliable() {
			fmt.Printf("frame corrupted\n")
			frame = nil
			continue
		}

		if *frame == exp {
			fmt.Printf("received frame %d\n", *frame)
			if !unreliable() {
				ack = frame
				exp++
			} else {
				fmt.Printf("ack lost\n")
			}
		}

		frame = nil
	}
}

func swDriver() {
	rand.Seed(time.Now().UnixNano())

	go swSender()
	go swReceiver()

	time.Sleep(time.Second * 20)
}
