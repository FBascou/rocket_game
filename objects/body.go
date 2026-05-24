package objects

import (
	"github.com/FBascou/rocket_game/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Body struct {
	Position math2d.Vector2
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
	return body.Position.X
}

// Getter function
func (body Body) GetY() float64 {
	return body.Position.Y
}
