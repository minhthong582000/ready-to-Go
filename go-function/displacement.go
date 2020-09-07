package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn  takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
// return a function which computes displacement as a function of time
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (1/2)*(a*math.Pow(t, 2)) + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Print("acceleration: ")
	fmt.Scan(&a)
	fmt.Print("initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("time: ")
	fmt.Scan(&t)

	f := GenDisplaceFn(a, v0, s0)

	fmt.Printf("----------\nDisplacement: %.2f\n", f(t))
}
