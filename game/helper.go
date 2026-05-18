package game

import (
	"image"
	"log"
	"os"

	"github.com/FBascou/rocket_game/entities"
	"github.com/FBascou/rocket_game/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// Interface for getter functions
type Positionable interface {
	GetX() float64
	GetY() float64
}

// Returns the destination planet (last in the slice)
func (game *Game) getDestinationPlanet() *entities.Planet {
	return &game.Planets[len(game.Planets)-1]
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

// Get the count of all minerals from all the planets in a level
func getTotalMineralsInLevel(planets []entities.Planet) int {
	minerals := 0
	for index := 0; index < len(planets); index++ {
		minerals += len(planets[index].Minerals)
	}
	return minerals
}

// Deep cloning planets and minerals for level restart
// Should be used for replaying whole level and restarting entire game
func deepClonePlanets(planets []entities.Planet) []entities.Planet {
	cloned := make([]entities.Planet, len(planets))

	for index, planet := range planets {
		cloned[index] = planet

		clonedMinerals := make([]entities.Mineral, len(planet.Minerals))
		copy(clonedMinerals, planet.Minerals)

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
		LaunchPower:                g.GameConfig.LaunchPower,
		MaxLaunchSpeed:             g.GameConfig.MaxLaunchSpeed,
		DragLineSteps:              g.GameConfig.DragLineSteps,
		DragPreviewStep:            g.GameConfig.DragPreviewStep,
		GravityFalloff:             g.GameConfig.GravityFalloff,
		DragPreviewGravityStrength: g.GameConfig.DragPreviewGravityStrength,
		ShipFriction:               g.GameConfig.ShipFriction,
	}
}
