package main

import (
	"fmt"
	"math"
)

var delta = 1e-4

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func NewtonSqrt(x float64) (int, float64) {
	z_old := 0.0
	z := x
	i := 0

	for Abs(z-z_old) > delta {
		z_old, z = z, z - (z * z - x) / (2 * z)
		i++
	}
	return i, z
}

func main() {
	for i := 0; i < 100; i++ {
		math := math.Sqrt(float64(i))
		iter, newton := NewtonSqrt(float64(i))
		diff := math - newton
		fmt.Printf("math.Sqrt(%v)  = %v\n", i, math)
		fmt.Printf("NewtonSqrt(%v) = %v\n", i, newton)
		fmt.Printf("Iterations     = %v\n", iter)
		fmt.Printf("Diff           = %v\n", diff)
		fmt.Printf("\n")
	}
}
