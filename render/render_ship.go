package render

import (
	"image/color"

	"github.com/FBascou/rocket_game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the ship image and its magnet outline
// GeoM order matters: Translate(center of body) > Scale > Rotate > Translate(world position)
func DrawShip(screen *ebiten.Image, ship objects.Ship) {
	drawFilledCircle(
		screen,
		ship.Position.X,
		ship.Position.Y,
		ship.MagnetRadius,
		color.RGBA{0, 30, 30, 30},
	)

	if ship.Sprite == nil {
		return
	}

	options := &ebiten.DrawImageOptions{}

	w := ship.Sprite.Bounds().Dx()
	h := ship.Sprite.Bounds().Dy()

	scale := (ship.Size * 2) / float64(w)

	options.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	options.GeoM.Scale(scale, scale)
	options.GeoM.Rotate(ship.Rotation)
	options.GeoM.Translate(ship.Position.X, ship.Position.Y)

	screen.DrawImage(ship.Sprite, options)
}
