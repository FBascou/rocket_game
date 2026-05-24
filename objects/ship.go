package objects

import (
	"github.com/FBascou/rocket_game/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	Position        math2d.Vector2
	Velocity        math2d.Vector2
	Acceleration    math2d.Vector2
	Size            float64
	Rotation        float64
	MagnetRadius    float64
	CollisionRadius float64
	SpriteKey       string
	Sprite          *ebiten.Image
}

// Getter function
func (s Ship) GetX() float64 {
	return s.Position.X
}

// Getter function
func (s Ship) GetY() float64 {
	return s.Position.Y
}
