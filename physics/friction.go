package physics

import (
	"github.com/FBascou/rocket_game/entities"
)

/*
Slows the ship down every frame, no matter what
0.99 will make it feel floaty
0.95 will make it feel more arcade
*/
func ApplyFriction(ship *entities.Ship, shipFriction float64) {
	ship.VX *= shipFriction
	ship.VY *= shipFriction
}
