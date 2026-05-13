package main

type Ship struct {
	X, Y         float64
	VX, VY       float64
	Rotation     float64
	MagnetRadius float64
}

// Getter function
func (s Ship) GetX() float64 {
	return s.X
}

// Getter function
func (s Ship) GetY() float64 {
	return s.Y
}