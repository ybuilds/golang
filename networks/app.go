package main

import "math/rand"

var frame, ack *int

const n = 10
const win = 3

func unreliable() bool {
	return rand.Float32() < 0.3
}

func main() {
	srDriver()
}
