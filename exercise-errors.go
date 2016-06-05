package main

import (
	"fmt"
)

var delta = 1e-4

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z_old := 0.0
	z := x
	i := 0
	
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	for Abs(z-z_old) > delta {
		z_old, z = z, z - (z * z - x) / (2 * z)
		i++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
