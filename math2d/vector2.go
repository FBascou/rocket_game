package math2d

import "math"

type Vector2 struct {
	X float64
	Y float64
}

func Add(a, b Vector2) Vector2 {
	return Vector2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func Sub(a, b Vector2) Vector2 {
	return Vector2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func MultiplyScalar(v Vector2, scalar float64) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func DivideScalar(v Vector2, scalar float64) Vector2 {
	if scalar == 0 {
		return Vector2{}
	}

	return Vector2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func Magnitude(v Vector2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MagnitudeSquared(v Vector2) float64 {
	return v.X*v.X + v.Y*v.Y
}

func Dot(a, b Vector2) float64 {
	return a.X*b.X + a.Y*b.Y
}
