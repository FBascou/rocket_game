package render

import (
	"image/color"
	"math"

	"github.com/FBascou/rocket_game/math2d"
	"github.com/FBascou/rocket_game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw the minerals image
func DrawMinerals(screen *ebiten.Image, body []objects.Body) {

	for _, body := range body {
		for _, mineral := range body.Minerals {
			if mineral.Collected {
				continue
			}

			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					// speed := math.Sqrt(mineral.Velocity.X*mineral.Velocity.X + mineral.Velocity.Y*mineral.Velocity.Y)
					speed := math2d.Speed(mineral.Velocity)
					brightness := uint8(math.Min(255, speed*50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(mineral.Position.X)+dx, int(mineral.Position.Y)+dy, col)
				}
			}
		}
	}
}
