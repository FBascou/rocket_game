package levels

import "github.com/FBascou/rocket_game/entities"

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
				BodyType:        entities.BodyPlanet,
				X:               500,
				Y:               220,
				Size:            30,
				GravityRadius:   80,
				GravityStrength: 0.5,
				SpriteKey:       "planet1",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},
			{
				ID:              1,
				BodyType:        entities.BodyPlanet,
				X:               380,
				Y:               320,
				Size:            30,
				GravityRadius:   80,
				GravityStrength: 0.5,
				SpriteKey:       "planet1",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},
			{
				ID:              2,
				BodyType:        entities.BodyDestination,
				X:               420,
				Y:               120,
				Size:            50,
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
			// more planets
		},
	},
}
