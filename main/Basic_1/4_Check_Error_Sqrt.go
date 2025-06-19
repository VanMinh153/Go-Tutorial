package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	const epsilon = 1e-10
	for i := 0; i < 1000; i++ {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prevZ) < epsilon {
			break
		}
	}
	return z, nil
}

func main() {
	testValues := []float64{1, 2, 3, 4, 9, -2}

	for _, v := range testValues {
		result, err := Sqrt(v)
		if err != nil {
			fmt.Printf("Sqrt(%v) error: %v\n", v, err)
		} else {
			fmt.Printf("Sqrt(%v) ≈ %v (math.Sqrt = %v)\n", v, result, math.Sqrt(v))
		}
	}
}
