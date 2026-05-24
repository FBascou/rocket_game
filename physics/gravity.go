package physics

import (
	"math"

	"github.com/FBascou/rocket_game/math2d"
	"github.com/FBascou/rocket_game/objects"
)

func CalculateGravityForce(
	from math2d.Vector2,
	body objects.Body,
	gravityFalloff float64,
	gravityMultiplier float64,
) math2d.Vector2 {

	dx := body.Position.X - from.X
	dy := body.Position.Y - from.Y

	distanceSq := dx*dx + dy*dy
	// continue skips only one interation and moves on to the next body in the loop

	if distanceSq < 1 {
		// continue
		return math2d.Vector2{}
	}

	distance := math.Sqrt(distanceSq)

	// 50 is a placeholder to not make it accelerate insanely fast when the ship gets closer to body
	// outside gravity field
	if distance > body.GravityRadius {
		// continue
		return math2d.Vector2{}
	}

	falloff := distance + gravityFalloff
	// force := body.GravityStrength / falloff
	// force := (body.GravityStrength * GravityMultiplier) / falloff
	// force := body.GravityStrength / (distanceance*distanceance + gravityFalloff)
	force := (body.GravityStrength * gravityMultiplier) / falloff

	nx := dx / distance
	ny := dy / distance

	return math2d.Vector2{
		X: nx * force,
		Y: ny * force,
	}
}

// Returns boolean if whether ship crashed into a non-destination body by being pulled by gravity
func ApplyGravity(
	ship *objects.Ship,
	bodies []objects.Body,
	gravityFalloff float64,
	GravityMultiplier float64,
) bool {
	for _, body := range bodies {
		distance := math2d.Distance(ship.Position, body.Position)

		// collision check + adding CollisionRadius to Ship or else the collision counts the center of the ship with the body
		collisionDistance := body.Size + ship.CollisionRadius - 4
		if distance < collisionDistance {
			if !body.IsDestination {
				return true
			}
		}

		force := CalculateGravityForce(ship.Position, body, gravityFalloff, GravityMultiplier)

		ship.Velocity.X += force.X
		ship.Velocity.Y += force.Y
	}
	return false
}
