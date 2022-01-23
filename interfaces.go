package main

import (
	"fmt"
	"math"
)

// NOTE: Interfaces are implemented implicitely. = there's no "implement" keyword
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// These works!
	// a MyFloat implements Abser
	a = f
	fmt.Println(a.Abs())

	// a *Vertex implements Abser
	a = &v
	fmt.Println(a.Abs())

	// Error!
	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	a = v
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
