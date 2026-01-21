package main

import "math/rand"

var frame, ack *int

func unreliable() bool {
	return rand.Float32() < 0.3
}

func main() {
	gbnDriver()
}
