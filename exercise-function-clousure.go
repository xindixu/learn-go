package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a_prev_prev := 0
	a_prev := 0

	return func() int {
		if a_prev_prev == 0 && a_prev == 0 {
			a_prev = 1
			return 0
		}

		cur := a_prev_prev + a_prev
		a_prev_prev = a_prev
		a_prev = cur
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
