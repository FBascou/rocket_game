package math2d

// Linear interpolation
// Calculates a value that sits at the specific percentage between two known points to create smooth transitions.
// start (x): starting point when t = 0
// end (y): ending point when t = 1
// t: interpolation factor, decimal between 0.0 (0%) and 1.0 (100%)
func Lerp(start, end, t float64) float64 {
	return start + (end-start)*t
}

func Vector2Lerp(start, end Vector2, t float64) Vector2 {
	return Vector2{
		X: start.X + (end.X-start.X)*t,
		Y: start.Y + (end.Y-start.Y)*t,
	}
}
