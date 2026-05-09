package main

import "github.com/hajimehoshi/ebiten/v2"

type Ship struct {
	X, Y          float64
	VX, VY        float64
	Rotation      float64
	Minerals      int
	MagnetRadius  float64
	Image *ebiten.Image
	MagnetRadiusOutline *ebiten.Image
}

// Getter function
func (s Ship) GetX() float64 {
	return s.X
}

// Getter function
func (s Ship) GetY() float64 {
	return s.Y
}