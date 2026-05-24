package math2d

import (
	"math"
)

func Distance(a, b Vector2) float64 {
	v := Sub(b, a)

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func DistanceSquared(a, b Vector2) float64 {
	v := Sub(b, a)

	return v.X*v.X + v.Y*v.Y
}
