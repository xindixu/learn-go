package main

import "fmt"

func yay() string {
	return "Hi"
}
func main() {
	// Declare i as an empty interface
	// NOTE: Empty interface: specifies zero methods
	// Empty interface can hold any type. Used by code that handles values of unknown type
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)

	describe(yay) // (0x109ae90, func() string)

}

// Any type can invoke this function
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
