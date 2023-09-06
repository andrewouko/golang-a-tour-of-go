package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
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
			return z
		}
	}
	return z
}

func main() {
	fmt.Println("return", Sqrt(25))

	fmt.Println(results)
}
