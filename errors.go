package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v", float64(err))
}

func sqrt(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var results [10]float64
	z := float64(1)
	for i := 0; i < len(results); i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		roundToTwo := func(num float64) float64 {
			return (math.Round(z * 100)) / 100
		}
		results[i] = roundToTwo(z)
		// fmt.Println(i > 0 && results[i - 1] == results[i])
		if i > 0 && results[i-1] == results[i] {
			return z, 0
		}
	}
	return z, 0
}
