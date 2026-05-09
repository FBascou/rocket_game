package main

import "github.com/hajimehoshi/ebiten/v2"

type Planet struct {
	X, Y           float64
	Radius         float64
	Gravity        float64
	Size           int
	Minerals       []Mineral
	Image *ebiten.Image
	GravityOutline *ebiten.Image
	IsDestination  bool
}

// Getter function
func (p Planet) GetX() float64 {
	return p.X
}

// Getter function
func (p Planet) GetY() float64 {
	return p.Y
}