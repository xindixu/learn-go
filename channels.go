package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	time.Sleep(1000 * time.Millisecond)

	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	length := len(s)

	/*
		NOTE: Channels must be created before use
		ch := make(chan int) // create channel ch with type int
		ch <- v    // Send v to channel ch.
		v := <-ch  // Receive from ch, and assign value to v.

		Channel sends and receives until the other side is ready
		Receivers always block until there is data to receive
		Sender blocks until the receiver has received the value
	*/
	c := make(chan int)

	fmt.Println("Summing...")

	go sum(s[:length/2], c)
	go sum(s[length/2:], c)
	// these will wait for the result to comeback
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
