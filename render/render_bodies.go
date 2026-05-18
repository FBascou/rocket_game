package render

import (
	"image/color"
	"math/rand/v2"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

var PlanetSprites = []string{
	"planet_green_01",
	"planet_green_02",
	"planet_blue_01",
	"planet_red_01",
}

func GetUniquePlanetSprite(used map[string]bool) string {

	available := []string{}

	for _, sprite := range PlanetSprites {
		if !used[sprite] {
			available = append(available, sprite)
		}
	}

	if len(available) == 0 {
		for k := range used {
			delete(used, k)
		}

		available = PlanetSprites
	}

	selected := available[rand.IntN(len(available))]
	used[selected] = true

	return selected
}

func drawPlanet(screen *ebiten.Image, planet entities.Planet) {
	drawFilledCircle(
		screen,
		planet.X,
		planet.Y,
		planet.Size,
		color.White,
	)
}

func drawBlackHole(screen *ebiten.Image, planet entities.Planet) {
	drawFilledCircle(
		screen,
		planet.X,
		planet.Y,
		planet.Size,
		color.White,
	)
}

func drawAsteroid(screen *ebiten.Image, planet entities.Planet) {
	drawFilledCircle(
		screen,
		planet.X,
		planet.Y,
		planet.Size,
		color.White,
	)
}

func drawComet(screen *ebiten.Image, planet entities.Planet) {
	drawFilledCircle(
		screen,
		planet.X,
		planet.Y,
		planet.Size,
		color.White,
	)
}

// drawWormHole

// TODO: Bodies should not be inside of Planets
func DrawBodies(screen *ebiten.Image, planets []entities.Planet) {
	for _, planet := range planets {

		switch planet.BodyType {
		case entities.BodyPlanet:
			drawPlanet(screen, planet) //TODO: Create these funcs drawPlanet

		case entities.BodyBlackHole:
			drawBlackHole(screen, planet) //TODO: Create these funcs drawBlackHole

		case entities.BodyAsteroid:
			drawAsteroid(screen, planet) //TODO: Create these funcs drawAsteroid

		case entities.BodyComet:
			drawComet(screen, planet) //TODO: Create these funcs drawComet
		case entities.BodyDestination:
			drawPlanet(screen, planet)
		}
	}
}
