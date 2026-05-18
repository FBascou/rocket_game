package game

import (
	"github.com/FBascou/rocket_game/entities"
	"github.com/FBascou/rocket_game/levels"
	"github.com/FBascou/rocket_game/render"
	"github.com/FBascou/rocket_game/shared"
)

func (game *Game) LoadLevel(index int) {
	if index < 0 || index >= len(levels.Levels) {
		return
	}

	level := levels.Levels[index]

	game.LevelNumber = index
	game.Lives = level.Lives

	game.Planets = []entities.Planet{}
	usedSprites := map[string]bool{}

	for _, p := range level.Planets {

		planet := entities.Planet{
			X:               p.X,
			Y:               p.Y,
			BodyType:        p.BodyType,
			Size:            p.Size,
			GravityRadius:   p.GravityRadius,
			GravityStrength: p.GravityStrength,
			IsDestination:   p.IsDestination,
			SpriteKey:       p.SpriteKey,
		}

		if p.AutoGenerateMinerals {
			planet.Minerals = game.generateMinerals(
				planet,
				p.MineralCount,
			)
		} else {
			for _, m := range p.Minerals {
				planet.Minerals = append(
					planet.Minerals,
					entities.Mineral{
						X: m.X,
						Y: m.Y,
					},
				)
			}
		}

		if p.SpriteKey == "" {
			p.SpriteKey = render.GetUniquePlanetSprite(usedSprites)
		}

		usedSprites[p.SpriteKey] = true

		game.Planets = append(game.Planets, planet)
	}

	game.Ship = entities.Ship{
		X:               level.Ship.X,
		Y:               level.Ship.Y,
		MagnetRadius:    level.Ship.MagnetRadius,
		CollisionRadius: 6,
	}

	game.InitialPlanets = deepClonePlanets(game.Planets)

	game.GameState = shared.StateAiming
}

func (game *Game) nextLevel() {
	if game.LevelNumber+1 >= len(levels.Levels) {
		return
	}

	game.LoadLevel(game.LevelNumber + 1)
}

func (game *Game) previousLevel() {
	if game.LevelNumber-1 < 0 {
		return
	}

	game.LoadLevel(game.LevelNumber - 1)
}
