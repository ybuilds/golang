package main

import (
	"fmt"
	"math/rand"
)

const (
	FRAME_LOSS_PROB = 0.25
	ACK_LOSS_PROB   = 0.25
	ACK_DELAY_PROB  = 0.25
	WINDOW_SIZE     = 4
)

func unreliableFrame(seq int) bool {
	if rand.Float64() < FRAME_LOSS_PROB {
		fmt.Println("Frame lost:", seq)
		return false
	}
	fmt.Println("Frame delivered:", seq)
	return true
}

func unreliableACK(seq int) (delivered bool, delayed bool) {
	if rand.Float64() < ACK_LOSS_PROB {
		fmt.Println("ACK lost:", seq)
		return false, false
	}
	if rand.Float64() < ACK_DELAY_PROB {
		fmt.Println("ACK delayed:", seq)
		return true, true
	}
	fmt.Println("ACK received:", seq)
	return true, false
}
