package main

import (
	"fmt"
	"time"
)

func main() {
	// Run in interval and sends the current time of the returned channel
	tick := time.Tick(100 * time.Millisecond)
	// Waits for the duration and sends the current time on the returned channel
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case t := <-tick:
			fmt.Printf("tick. %v\n", t)
		case b := <-boom:
			// b is the current time
			fmt.Printf("BOOM! %v\n", b)
			return
		default:
			fmt.Printf("    .\n")
			// Pause the current goroutine for at least the duraction
			time.Sleep(50 * time.Millisecond)
		}
	}
}
