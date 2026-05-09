package main

import "github.com/hajimehoshi/ebiten/v2"

type Mineral struct {
	X, Y      float64
	VX, VY    float64
	Collected bool
	Image     *ebiten.Image
	// Value     int
	// color R, G, B, A uint8
}