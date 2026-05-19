package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Body struct {
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
func (body Body) GetX() float64 {
	return body.X
}

// Getter function
func (body Body) GetY() float64 {
	return body.Y
}
