package physics

import "math"

// Limit velocity on drag and release
func ClampVelocity(vx, vy, max float64) (float64, float64) {
	// speed = length of velocity vector:
	// pythagorean theorem: speed² = vx² + vy² => speed = √(vx² + vy²)
	speed := math.Sqrt(vx*vx + vy*vy)

	if speed > max {
		scale := max / speed
		vx *= scale
		vy *= scale
	}

	return vx, vy
}
