package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method
// pointer indirection
// When calling the method, receiver can be a value or a pointer, regardless the defined receiver type
func (v *Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function
// When calling the function, arguments must be the same type as defined
func ScaleFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// no need to get pointer from value or get value from pointer
	v.ScaleMethod(2)
	fmt.Println(v.AbsMethod())

	// must get pointer from value or get value from pointer to match the defined type
	ScaleFunction(&v, 10)
	fmt.Println(AbsFunction(v))

	p := &Vertex{4, 3}
	// no need to get pointer from value or get value from pointer
	p.ScaleMethod(3)
	fmt.Println(p.AbsMethod())

	// must get pointer from value or get value from pointer to match the defined type
	ScaleFunction(p, 8)
	fmt.Println(AbsFunction(*p))

	fmt.Println(v, p)

}
