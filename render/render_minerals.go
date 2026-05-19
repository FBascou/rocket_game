package render

import (
	"image/color"
	"math"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the minerals image
func DrawMinerals(screen *ebiten.Image, body []entities.Body) {

	for _, body := range body {
		for _, mineral := range body.Minerals {
			if mineral.Collected {
				continue
			}

			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					speed := math.Sqrt(mineral.VX*mineral.VX + mineral.VY*mineral.VY)
					brightness := uint8(math.Min(255, speed*50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(mineral.X)+dx, int(mineral.Y)+dy, col)
				}
			}
		}
	}
}
