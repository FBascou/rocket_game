package render

import (
	"image/color"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the ship image and its magnet outline
// GeoM order matters: Translate(center of body) > Scale > Rotate > Translate(world position)
func DrawShip(screen *ebiten.Image, ship entities.Ship) {
	drawFilledCircle(
		screen,
		ship.X,
		ship.Y,
		ship.MagnetRadius,
		color.RGBA{0, 30, 30, 30},
	)

	drawFilledCircle(
		screen,
		ship.X,
		ship.Y,
		6,
		color.White,
	)
}
