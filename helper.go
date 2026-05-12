package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RadomXY struct {
	X, Y float64	
}

// Interface for getter functions
type Positionable interface {
	GetX() float64
	GetY() float64
}

// Limit velocity on drag and release
func clampVelocity(vx, vy, max float64) (float64, float64) {
	// speed = length of velocity vector:
	// pythagorean theorem: speed² = vx² + vy² => speed = √(vx² + vy²)
	speed := math.Sqrt(vx * vx + vy * vy)

	if speed > max {
		scale := max / speed
		vx *= scale
		vy *= scale
	}

	return vx, vy
}

// Returns the destination planet (last in the slice)
func (game *Game) getDestinationPlanet() *Planet {
	return &game.Planets[len(game.Planets)-1]
}

// Creates a reusable circle image
func createCircleImage(radius int, clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(radius * 2, radius * 2)

	for i := -radius; i < radius; i++ {
		for j := -radius; j < radius; j++ {
			if i * i + j * j <= radius * radius {
				img.Set(i + radius, j + radius, clr)
			}
		}
	}

	return img
}

// Creates the background circle outline image for objects
func createOutlineImage(screen *ebiten.Image, object Positionable, metric float64, outline *ebiten.Image, alpha float32) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		object.GetX() - metric, 
		object.GetY() - metric,
	)
	// set 10% opacity
	options.ColorScale.ScaleAlpha(alpha)
	screen.DrawImage(outline, options)
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
func getTotalMineralsInLevel(planets []Planet) int {
	minerals := 0
	for index := 0; index < len(planets); index++ {
		minerals += len(planets[index].Minerals)
	}
	return minerals
}

// Deep cloning planets and minerals for level restart
// Should be used for replaying whole level and restarting entire game 
func deepClonePlanets(planets []Planet) []Planet {
    cloned := make([]Planet, len(planets))

    for index, planet := range planets {
        cloned[index] = planet

        clonedMinerals := make([]Mineral, len(planet.Minerals))
        copy(clonedMinerals, planet.Minerals)

        cloned[index].Minerals = clonedMinerals
    }

    return cloned
}

func (game *Game) drawButton(
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

func isPointInsideRect(
	px, py int,
	x, y, w, h float64,
) bool {
	return float64(px) >= x &&
		float64(px) <= x + w &&
		float64(py) >= y &&
		float64(py) <= y + h
}
