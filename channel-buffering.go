package main

import "fmt"

func main() {
	/*
		NOTE: Channels can be buffered
		ch := make(chan int, 2)
		This channel will only accept 2 int values

		Sends to a buffered channel block only when the buffer is full
		Receives block when the buffer is empty

		Receivers always block until there is data to receive
		For a buffered channel:
		Sender blocks until the values has been copied to the buffer
		If the buffer is full, sender blocks until some receiver has retrieved a value

	*/

	ch := make(chan int, 2)
	// Send 2 values and receive 2 values
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Send 3 values without receive any -> error!
	ch <- 3
	ch <- 4
	ch <- 5
}
