package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayP(s *string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(*s)
	}
}
func main() {
	/*
		NOTE:
		Goroutine is a function executing concurrently with other goroutines in the same address space.
		go f(x,y,z)
		The evaluation of f, x, y, and z happens in the current goroutine and
		the execution of f happens in the new goroutine.
	*/
	// order of the words printed out might vary
	word := "world"
	go say(word) // prints out "world"
	word = "hello"
	say(word) // prints out "hello"

	// pWord didn't change, but the word changes
	word = "world"
	pWord := &word
	go sayP(pWord) // prints out "hello"
	word = "hello"
	sayP(pWord) // prints out "hello"

}
