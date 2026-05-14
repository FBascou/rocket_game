package main

import (
	"math"
)

func (game *Game) applyGravity() {
	for _, planet := range game.Planets {
		dx := planet.X - game.Ship.X
		dy := planet.Y - game.Ship.Y

		distance := math.Sqrt(dx*dx + dy*dy)

		// collision check
		if distance < planet.Radius {
			if !planet.IsDestination {
				game.GameState = StateCrashed
			}
		}

		if distance < 1 {
			// continue skips only one interation and moves on to the next planet in the loop
				continue
		}

		// 50 is a placeholder to not make it accelerate insanely fast when the ship gets closer to planet
		force := planet.Gravity / (distance * distance + game.GameConfig.GravityFalloff)

		nx := dx / distance
		ny := dy / distance

		game.Ship.VX += nx * force
		game.Ship.VY += ny * force
	}
}

/*
Slows the ship down every frame, no matter what
0.99 will make it feel floaty
0.95 will make it feel more arcade
*/
func (game *Game) applyFriction() {
	game.Ship.VX *= game.GameConfig.ShipFriction
	game.Ship.VY *= game.GameConfig.ShipFriction
}

// Speed clamping allows normal movement to remain uncapped but extreme velocity gradually stabilizes
// It only activates when speed becomes too high
// Limits gravity acceleration once the ship is launched
func (game *Game) applySpeedDamping() {
	speed := math.Sqrt(game.Ship.VX * game.Ship.VX + game.Ship.VY * game.Ship.VY)

	// if ship becomes too fast, gradually slow it down
	// 0.98 means keep 98% of current speed every frame so the ship loses 2% speed per frame
	if speed > game.GameConfig.MaxShipSpeed {
		game.Ship.VX *= 0.98
		game.Ship.VY *= 0.98
	}
}

// Limit velocity on drag and release
func clampVelocity(vx, vy, max float64) (float64, float64) {
	// speed = length of velocity vector:
	// pythagorean theorem: speed² = vx² + vy² => speed = √(vx² + vy²)
	speed := math.Sqrt(vx * vx + vy * vy)

	if speed > max {
		scale := max / speed
		vx *= scale
		vy *= scale
	}

	return vx, vy
}