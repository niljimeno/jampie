package imath

import (
	"math"
)

func Clamp(value float64, min float64, max float64) float64 {
	return math.Max(math.Min(value, max), min)
}
