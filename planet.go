package main

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
}

// Getter function
func (p Planet) GetX() float64 {
	return p.X
}

// Getter function
func (p Planet) GetY() float64 {
	return p.Y
}
