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
				X:               350,
				Y:               420,
				Size:            30,
				GravityRadius:   80,
				GravityStrength: 0.5,
				SpriteKey:       "planet_green_01",

				AutoGenerateMinerals: true,
				MineralCount:         3,
			},

			{
				ID:              1,
				BodyType:        entities.BodyDestination,
				X:               420,
				Y:               120,
				Size:            50,
				GravityRadius:   140,
				GravityStrength: 2.5,
				IsDestination:   true,
				SpriteKey:       "planet_blue_01",
			},
		},
	},

	{
		ID:    1,
		Name:  "Gravity Fork",
		Lives: 4,

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
