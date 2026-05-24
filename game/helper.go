package game

import (
	"image"
	"log"
	"os"

	"github.com/FBascou/rocket_game/objects"
	"github.com/FBascou/rocket_game/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// Interface for getter functions
type Positionable interface {
	GetX() float64
	GetY() float64
}

// Returns the destination planet (last in the slice)
func (game *Game) getDestinationPlanet() *objects.Body {
	// TODO: add if statment if body is planet and isDestination
	return &game.Bodies[len(game.Bodies)-1]
}

func loadImage(path string) *ebiten.Image {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

// Get the count of all minerals from all the bodies in a level
func getTotalMineralsInLevel(bodies []objects.Body) int {
	minerals := 0
	for index := 0; index < len(bodies); index++ {
		minerals += len(bodies[index].Minerals)
	}
	return minerals
}

// Deep cloning bodies and minerals for level restart
// Should be used for replaying whole level and restarting entire game
func deepCloneBodies(bodies []objects.Body) []objects.Body {
	cloned := make([]objects.Body, len(bodies))

	for index, body := range bodies {
		cloned[index] = body

		clonedMinerals := make([]objects.Mineral, len(body.Minerals))
		copy(clonedMinerals, body.Minerals)

		cloned[index].Minerals = clonedMinerals
	}

	return cloned
}

func isPointInsideRect(
	px, py int,
	x, y, w, h float64,
) bool {
	return float64(px) >= x &&
		float64(px) <= x+w &&
		float64(py) >= y &&
		float64(py) <= y+h
}

func (g *Game) GetDragIndicatorConfig() render.DragIndicatorConfig {
	return render.DragIndicatorConfig{
		LaunchPower:       g.GameConfig.LaunchPower,
		MaxLaunchSpeed:    g.GameConfig.MaxLaunchSpeed,
		DragLineSteps:     g.GameConfig.DragLineSteps,
		DragPreviewStep:   g.GameConfig.DragPreviewStep,
		GravityFalloff:    g.GameConfig.GravityFalloff,
		GravityMultiplier: g.GameConfig.GravityMultiplier,
		ShipFriction:      g.GameConfig.ShipFriction,
	}
}
