package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Planet struct {
	X, Y float64
	// Physical size
	Size float64
	// Gravity field size
	GravityRadius float64
	// Gravity pull size
	GravityStrength float64
	Minerals        []Mineral
	IsDestination   bool

	BodyType  BodyType
	SpriteKey string
	Sprite    *ebiten.Image
}

// Getter function
func (planet Planet) GetX() float64 {
	return planet.X
}

// Getter function
func (planet Planet) GetY() float64 {
	return planet.Y
}
