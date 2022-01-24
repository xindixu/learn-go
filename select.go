package main

import "fmt"

/*
NOTE:
`select` lets a goroutine wait on multiple communication operations
A `select` blocks until one of its cases can run, then it execute this case.
It chooses one at random if multiple can run.
*/
func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// send x to ch, can always run
		case ch <- x:
			// fmt.Println("inside fib, just send %v to channel", x)
			x, y = y, x+y
		// receive from quit, can only run when exits data in quit
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)
	// goroutine runs concurrently with other goroutines
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: %v\n", i, <-ch)
		}
		quit <- 0
		fmt.Printf("after done %v", <-ch)
	}()
	fibonacci(ch, quit)
}
