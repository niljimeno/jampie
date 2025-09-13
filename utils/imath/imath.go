package imath

import (
	"math"
)

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func Clamp(value float64, min float64, max float64) float64 {
	return math.Max(math.Min(value, max), min)
}