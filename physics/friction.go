package physics

import (
	"github.com/FBascou/rocket_game/math2d"
)

/*
Slows the ship down every frame, no matter what
0.99 will make it feel floaty
0.95 will make it feel more arcade
*/
func ApplyFriction(velocity *math2d.Vector2, shipFriction float64) {
	velocity.X *= shipFriction
	velocity.Y *= shipFriction
}
