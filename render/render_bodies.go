package render

import (
	"image/color"
	"math/rand/v2"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

var bodiesprites = []string{
	"planet_green_01",
	"planet_green_02",
	"planet_blue_01",
	"planet_red_01",
}

func GetUniqueBodySprite(used map[string]bool) string {

	available := []string{}

	for _, sprite := range bodiesprites {
		if !used[sprite] {
			available = append(available, sprite)
		}
	}

	if len(available) == 0 {
		for k := range used {
			delete(used, k)
		}

		available = bodiesprites
	}

	selected := available[rand.IntN(len(available))]
	used[selected] = true

	return selected
}

func drawBody(screen *ebiten.Image, body entities.Body) {
	drawFilledCircle(
		screen,
		body.X,
		body.Y,
		body.Size,
		color.White,
	)
}

func drawBlackHole(screen *ebiten.Image, body entities.Body) {
	drawFilledCircle(
		screen,
		body.X,
		body.Y,
		body.Size,
		color.White,
	)
}

func drawAsteroid(screen *ebiten.Image, body entities.Body) {
	drawFilledCircle(
		screen,
		body.X,
		body.Y,
		body.Size,
		color.White,
	)
}

func drawComet(screen *ebiten.Image, body entities.Body) {
	drawFilledCircle(
		screen,
		body.X,
		body.Y,
		body.Size,
		color.White,
	)
}

// drawWormHole

// TODO: Bodies should not be inside of bodies
func DrawBodies(screen *ebiten.Image, bodies []entities.Body) {
	for _, body := range bodies {

		switch body.BodyType {
		case entities.BodyPlanet:
			drawBody(screen, body) //TODO: Create these funcs drawBody

		case entities.BodyBlackHole:
			drawBlackHole(screen, body) //TODO: Create these funcs drawBlackHole

		case entities.BodyAsteroid:
			drawAsteroid(screen, body) //TODO: Create these funcs drawAsteroid

		case entities.BodyComet:
			drawComet(screen, body) //TODO: Create these funcs drawComet
		case entities.BodyDestination:
			drawBody(screen, body)
		}
	}
}
