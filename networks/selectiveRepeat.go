package main

import "fmt"

func selectiveRepeat(total int) {
	acked := make([]bool, total)
	base := 0

	for base < total {
		// Send unacknowledged frames in window
		for i := base; i < base+WINDOW_SIZE && i < total; i++ {
			if !acked[i] {
				unreliableFrame(i)
			}
		}

		// Receive acknowledged independently
		for i := base; i < base+WINDOW_SIZE && i < total; i++ {
			if acked[i] {
				continue
			}

			ok, delayed := unreliableACK(i)
			if ok && !delayed {
				acked[i] = true
			}
		}

		// Slide window
		for base < total && acked[base] {
			base++
		}

		fmt.Println()
	}
}
