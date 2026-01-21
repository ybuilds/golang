package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ackMap = make(map[int]bool)
var lock sync.Mutex

func srSender() {
	base := 0

	for base < n {
		for i := base; i < base+win && i < n; i++ {
			lock.Lock()
			acked := ackMap[i]
			lock.Unlock()

			if !acked {
				fmt.Printf("sending frame %d\n", i)
				f := i
				frame = &f
			}
		}

		start := time.Now()
		for time.Since(start) < time.Second {
			lock.Lock()
			for base < n && ackMap[base] {
				base++
			}
			lock.Unlock()
		}
	}
}

func srReceiver() {
	buf := make(map[int]bool)

	for len(buf) < n {
		if frame == nil {
			continue
		}

		if unreliable() {
			fmt.Printf("frame corrupted\n")
			frame = nil
			continue
		}

		if buf[*frame] {
			frame = nil
			continue
		}

		fmt.Printf("frame %d received\n", *frame)
		buf[*frame] = true

		if !unreliable() {
			lock.Lock()
			ackMap[*frame] = true
			lock.Unlock()
		} else {
			fmt.Printf("ack lost\n")
		}

		frame = nil
	}
}

func srDriver() {
	rand.Seed(time.Now().UnixNano())

	go srSender()
	go srReceiver()

	time.Sleep(time.Second * 30)
}
