package physics

import (
	"math"

	"github.com/FBascou/rocket_game/entities"
)

// Speed clamping allows normal movement to remain uncapped but extreme velocity gradually stabilizes
// It only activates when speed becomes too high
// Limits gravity acceleration once the ship is launched
func ApplySpeedDamping(ship *entities.Ship, maxShipSpeed float64) {
	speed := math.Sqrt(ship.VX*ship.VX + ship.VY*ship.VY)

	// if ship becomes too fast, gradually slow it down
	// 0.98 means keep 98% of current speed every frame so the ship loses 2% speed per frame
	if speed > maxShipSpeed {
		ship.VX *= 0.98
		ship.VY *= 0.98
	}
}
