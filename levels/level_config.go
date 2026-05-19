package levels

import "github.com/FBascou/rocket_game/entities"

type MineralConfig struct {
	X float64
	Y float64
}

type BodyConfig struct {
	ID              int
	BodyType        entities.BodyType
	X               float64
	Y               float64
	Size            float64
	GravityRadius   float64
	GravityStrength float64
	IsDestination   bool
	Minerals        []MineralConfig

	// Visual
	SpriteKey string

	// Optional generation
	AutoGenerateMinerals bool
	MineralCount         int
}

type ShipConfig struct {
	X            float64
	Y            float64
	MagnetRadius float64
}

type LevelConfig struct {
	ID     int
	Name   string
	Lives  int
	Bodies []BodyConfig
	Ship   ShipConfig
}
