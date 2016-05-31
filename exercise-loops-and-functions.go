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

func Sqrt(x float64) float64 {
	z_old := 0.0
	z := x

	for Abs(z-z_old) > delta {
		z_old, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
