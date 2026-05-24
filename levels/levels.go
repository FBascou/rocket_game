package levels

import "github.com/FBascou/rocket_game/objects"

var Levels = []LevelConfig{
	{
		ID:    0,
		Name:  "First Launch",
		Lives: 5,

		Ship: ShipConfig{
			X:            400,
			Y:            540,
			MagnetRadius: 75,
		},

		Bodies: []BodyConfig{
			{
				ID:              0,
				BodyType:        objects.BodyPlanet,
				X:               400,
				Y:               400,
				Size:            50,
				GravityRadius:   140,
				GravityStrength: 1,
				SpriteKey:       "planet1",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},
			{
				ID:              2,
				BodyType:        objects.BodyDestination,
				X:               420,
				Y:               120,
				Size:            75,
				GravityRadius:   140,
				GravityStrength: 2.5,
				IsDestination:   true,
				SpriteKey:       "planet2",
			},
		},
	},

	{
		ID:    1,
		Name:  "Gravity Fork",
		Lives: 5,

		Ship: ShipConfig{
			X:            400,
			Y:            540,
			MagnetRadius: 75,
		},

		Bodies: []BodyConfig{
			{
				ID:              0,
				BodyType:        objects.BodyPlanet,
				X:               420,
				Y:               300,
				Size:            50,
				GravityRadius:   80,
				GravityStrength: 0.8,
				SpriteKey:       "planet1",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},
			{
				ID:              1,
				BodyType:        objects.BodyPlanet,
				X:               300,
				Y:               400,
				Size:            50,
				GravityRadius:   80,
				GravityStrength: 0.8,
				SpriteKey:       "planet1",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},
			{
				ID:              2,
				BodyType:        objects.BodyDestination,
				X:               420,
				Y:               120,
				Size:            75,
				GravityRadius:   140,
				GravityStrength: 2.5,
				IsDestination:   true,
				SpriteKey:       "planet2",
			},
		},
	},
}
