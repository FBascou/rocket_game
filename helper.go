package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
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
	speed := math.Sqrt(vx*vx + vy*vy)

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
