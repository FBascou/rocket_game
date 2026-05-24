package objects

import "github.com/FBascou/rocket_game/math2d"

type Mineral struct {
	Position  math2d.Vector2
	Velocity  math2d.Vector2
	Collected bool
	// Value     int
	// color R, G, B, A uint8
}
