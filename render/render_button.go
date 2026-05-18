package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawButton(
	screen *ebiten.Image,
	x, y, w, h float64,
	label string,
) {
	vector.FillRect(
		screen,
		float32(x),
		float32(y),
		float32(w),
		float32(h),
		color.RGBA{80, 80, 80, 255},
		false, // anti-aliasing
	)

	ebitenutil.DebugPrintAt(
		screen,
		label,
		int(x+70),
		int(y+12),
	)
}
