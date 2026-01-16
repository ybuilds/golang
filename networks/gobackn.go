package main

import "fmt"

func goBackN(total int) {
	base := 0
	nextSeq := 0

	for base < total {
		// Send window
		for nextSeq < base+WINDOW_SIZE && nextSeq < total {
			unreliableFrame(nextSeq)
			nextSeq++
		}

		ackProgress := false

		// Receive acknowledgement
		for seq := base; seq < nextSeq; seq++ {
			ok, delayed := unreliableACK(seq)
			if !ok {
				break
			}
			if delayed {
				fmt.Println("Timeout waiting for delayed ACK → retransmit")
				break
			}
			base++
			ackProgress = true
		}

		if !ackProgress {
			fmt.Println("Timeout → Go back to", base)
			nextSeq = base
		}

		fmt.Println()
	}
}
