package physics

import (
	"math"

	"github.com/FBascou/rocket_game/entities"
)

// Returns boolean if whether ship crashed into a non-destination body by being pulled by gravity
func ApplyGravity(
	ship *entities.Ship,
	bodies []entities.Body,
	gravityFalloff float64,
	dragPreviewGravityStrength float64,
) bool {
	for _, body := range bodies {
		dx := body.X - ship.X
		dy := body.Y - ship.Y

		distance := math.Sqrt(dx*dx + dy*dy)

		// collision check + adding CollisionRadius to Ship or else the collision counts the center of the ship with the body
		if distance < body.Size+ship.CollisionRadius {
			if !body.IsDestination {
				return true
			}
		}

		if distance < 1 {
			// continue skips only one interation and moves on to the next body in the loop
			continue
		}

		// 50 is a placeholder to not make it accelerate insanely fast when the ship gets closer to body
		// outside gravity field
		if distance > body.GravityRadius {
			continue
		}

		// force := body.GravityStrength / (distance*distance + game.GameConfig.GravityFalloff)

		falloff := distance + gravityFalloff
		// force := body.GravityStrength / falloff
		force := (body.GravityStrength * dragPreviewGravityStrength) / falloff

		nx := dx / distance
		ny := dy / distance

		ship.VX += nx * force
		ship.VY += ny * force
	}
	return false
}
