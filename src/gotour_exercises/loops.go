package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	iterations := 0
	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		iterations++
		if math.Abs(prev-z) < 0.000001 {
			fmt.Println("Iterations:", iterations)
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
