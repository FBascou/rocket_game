package physics

import (
	"github.com/FBascou/rocket_game/math2d"
)

// Speed clamping allows normal movement to remain uncapped but extreme velocity gradually stabilizes
// It only activates when speed becomes too high
// Limits gravity acceleration once the ship is launched
func ApplySpeedDamping(velocity *math2d.Vector2, maxShipSpeed float64) {
	speed := math2d.Speed(*velocity)

	// if ship becomes too fast, gradually slow it down
	// 0.98 means keep 98% of current speed every frame so the ship loses 2% speed per frame
	if speed > maxShipSpeed {
		// velocity.X *= 0.98
		// velocity.Y *= 0.98

		scale := maxShipSpeed / speed

		velocity.X *= scale
		velocity.Y *= scale
	}
}
