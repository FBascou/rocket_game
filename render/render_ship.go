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

	if ship.Sprite == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	w := ship.Sprite.Bounds().Dx()
	h := ship.Sprite.Bounds().Dy()

	scale := 24.0 / float64(w)

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Rotate(ship.Rotation)
	op.GeoM.Translate(ship.X, ship.Y)

	screen.DrawImage(ship.Sprite, op)
}
