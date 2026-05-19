package render

import (
	"image/color"
	"math/rand/v2"

	"github.com/FBascou/rocket_game/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

var bodiesprites = []string{
	"planet1",
	"planet2",
	"planet3",
	"planet4",
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
		body.GravityRadius,
		color.RGBA{0, 30, 30, 30},
	)

	if body.Sprite == nil {
		return
	}

	options := &ebiten.DrawImageOptions{}

	w, h := body.Sprite.Bounds().Dx(), body.Sprite.Bounds().Dy()

	scaleX := (body.Size * 2) / float64(w)
	scaleY := (body.Size * 2) / float64(h)

	options.GeoM.Scale(scaleX, scaleY)

	options.GeoM.Translate(
		body.X-body.Size,
		body.Y-body.Size,
	)

	screen.DrawImage(body.Sprite, options)
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
