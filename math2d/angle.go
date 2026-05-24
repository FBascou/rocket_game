package math2d

import "math"

// Returns the rotation angle of a point
func Angle(v Vector2) float64 {
	return math.Atan2(v.Y, v.X)
}

func AngleBetween(a, b Vector2) float64 {
	return math.Atan2(b.Y-a.Y, b.X-a.X)
}

func Rotate(v Vector2, radians float64) Vector2 {
	cosA := math.Cos(radians)
	sinA := math.Sin(radians)

	X := v.X*cosA - v.Y*sinA
	Y := v.X*sinA + v.Y*cosA

	return Vector2{X, Y}
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
