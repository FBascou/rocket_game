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
		force := planet.Gravity / (distance * distance + 50)

		nx := dx / distance
		ny := dy / distance

		game.Ship.VX += nx * force
		game.Ship.VY += ny * force
	}
}

/*
0.99 will make it feel floaty
0.95 will make it feel more arcade
*/
func (game *Game) applyFriction() {
	game.Ship.VX *= 0.99
	game.Ship.VY *= 0.99
}
