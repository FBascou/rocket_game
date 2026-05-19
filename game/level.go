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

	game.Bodies = []entities.Body{}
	usedSprites := map[string]bool{}

	for _, b := range level.Bodies {

		body := entities.Body{
			X:               b.X,
			Y:               b.Y,
			BodyType:        b.BodyType,
			Size:            b.Size,
			GravityRadius:   b.GravityRadius,
			GravityStrength: b.GravityStrength,
			IsDestination:   b.IsDestination,
			SpriteKey:       b.SpriteKey,
			Sprite:          game.Assets.Bodies[b.SpriteKey],
		}

		if b.AutoGenerateMinerals {
			body.Minerals = game.generateMinerals(
				body,
				b.MineralCount,
			)
		} else {
			for _, mineral := range b.Minerals {
				body.Minerals = append(
					body.Minerals,
					entities.Mineral{
						X: mineral.X,
						Y: mineral.Y,
					},
				)
			}
		}

		if b.SpriteKey == "" {
			b.SpriteKey = render.GetUniqueBodySprite(usedSprites)
		}

		usedSprites[b.SpriteKey] = true

		game.Bodies = append(game.Bodies, body)
	}

	game.Ship = entities.Ship{
		X:               level.Ship.X,
		Y:               level.Ship.Y,
		MagnetRadius:    level.Ship.MagnetRadius,
		CollisionRadius: 6,
		SpriteKey:       "ship",
		Sprite:          game.Assets.Ships["ship"],
	}

	game.InitialBodies = deepCloneBodies(game.Bodies)

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
