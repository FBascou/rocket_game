package physics

import (
	"math"

	"github.com/FBascou/rocket_game/entities"
)

// Returns boolean if whether ship crashed into a non-destination planet by being pulled by gravity
func ApplyGravity(
	ship *entities.Ship,
	planets []entities.Planet,
	gravityFalloff float64,
	dragPreviewGravityStrength float64,
) bool {
	for _, planet := range planets {
		dx := planet.X - ship.X
		dy := planet.Y - ship.Y

		distance := math.Sqrt(dx*dx + dy*dy)

		// collision check + adding CollisionRadius to Ship or else the collision counts the center of the ship with the planet
		if distance < planet.Size+ship.CollisionRadius {
			if !planet.IsDestination {
				return true
			}
		}

		if distance < 1 {
			// continue skips only one interation and moves on to the next planet in the loop
			continue
		}

		// 50 is a placeholder to not make it accelerate insanely fast when the ship gets closer to planet
		// outside gravity field
		if distance > planet.GravityRadius {
			continue
		}

		// force := planet.GravityStrength / (distance*distance + game.GameConfig.GravityFalloff)

		falloff := distance + gravityFalloff
		// force := planet.GravityStrength / falloff
		force := (planet.GravityStrength * dragPreviewGravityStrength) / falloff

		nx := dx / distance
		ny := dy / distance

		ship.VX += nx * force
		ship.VY += ny * force
	}
	return false
}
