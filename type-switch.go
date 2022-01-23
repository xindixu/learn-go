package main

import "fmt"

func describe(i interface{}) {
	// value, concrete type
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	describe(i)
	// NOTE: Type switch
	// Cases specify types (not values)
	// Compare to the type of the value held by the given interface value
	// `type` keyword: used to test
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
