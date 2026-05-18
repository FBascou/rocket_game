package render

import (
	"image/color"
	"math"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the minerals image
func DrawMinerals(screen *ebiten.Image, planets []entities.Planet) {

	for _, p := range planets {
		for _, m := range p.Minerals {
			if m.Collected {
				continue
			}

			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					speed := math.Sqrt(m.VX*m.VX + m.VY*m.VY)
					brightness := uint8(math.Min(255, speed*50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(m.X)+dx, int(m.Y)+dy, col)
				}
			}
		}
	}
}
