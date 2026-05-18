package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawFilledCircle(
	screen *ebiten.Image,
	x, y, radius float64,
	clr color.Color,
) {
	vector.FillCircle(
		screen,
		float32(x),
		float32(y),
		float32(radius),
		clr,
		false,
	)
}
