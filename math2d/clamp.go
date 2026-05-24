package math2d

import (
	"math"
)

func Clamp(value, min, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

// Velocity clamp
// Limit velocity on drag and release
func ClampMagnitude(v Vector2, max float64) Vector2 {
	speed := Speed(v)
	if speed > max {
		scale := max / speed
		v.X *= scale
		v.Y *= scale
	}

	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}

// Restricts value to 0 to 1
// Used for interpolation, alpha, and percentages
func Clamp01(value float64) float64 {
	return math.Max(0, math.Min(value, 1))
}
