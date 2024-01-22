package test

import "math"

func Sum(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func Abs(value float64) float64 {
	return math.Abs(value)
}
