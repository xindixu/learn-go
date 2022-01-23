package main

import "fmt"

/*
NOTE: Type assertion

t := i.(T)
Assert the interface value i holds a concrete type T
If true, return the underlying T value
If false, trigger a panic

t, ok := i.(T)
Assert the interface value i holds a concrete type T
If true, return the underlying T value and true
If false, return the zero value of type T and false. No panic

*/
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	describe(s) // (hello, String)

	s, ok := i.(string)
	fmt.Println(s, ok)
	describe(s) // (hello, String)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	describe(f) // (0, float)

	f = i.(float64) // panic
	fmt.Println(f)
	describe(f)

}

func describe(i interface{}) {
	// value, concrete type
	fmt.Printf("(%v, %T)\n", i, i)
}
