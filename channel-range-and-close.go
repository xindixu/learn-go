package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
		// send x to channel
		c <- x
	}

	// close the channel to indicate there's no more value will be sent
	// only the sender should close the channel
	// We don't usually need to close channels. It is only necessary when the receiver ust be told there're no more values coming
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// i receives value from channel repeatedly until channel is closed (receiver)
	for i := range c {
		fmt.Println(i)
	}
}
