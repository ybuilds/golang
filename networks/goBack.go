package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10
const win = 3

func gbnSender() {
	base := 0
	for base < n {
		for i := base; i < base+win && i < n; i++ {
			fmt.Printf("sending %d\n", i)
			f := i
			frame = &f
		}

		start := time.Now()
		for time.Since(start) < time.Second {
			if ack != nil && *ack >= base {
				base = *ack + 1
				break
			}
		}

		if ack == nil || *ack < base {
			fmt.Printf("timeout. resend from %d\n", base)
		}
	}
}

func gbnReceiver() {
	exp := 0
	for exp < n {
		if frame == nil {
			continue
		}

		if unreliable() {
			fmt.Printf("frame lost\n")
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
		} else {
			fmt.Printf("discarded out-of-order frames\n")
		}

		frame = nil
	}
}

func gbnDriver() {
	rand.Seed(time.Now().UnixNano())

	go gbnSender()
	go gbnReceiver()

	time.Sleep(time.Second * 20)
}
