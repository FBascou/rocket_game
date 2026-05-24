package game

import (
	"math"
	"math/rand/v2"

	"github.com/FBascou/rocket_game/math2d"
	"github.com/FBascou/rocket_game/objects"
	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Initializes ebiten.Image images for ship, bodies, minerals, etc.
func (game *Game) InitializeAssets() {
	game.Assets.Ships = map[string]*ebiten.Image{
		"ship": loadImage("assets/ships/ship.png"),
	}

	game.Assets.Bodies = map[string]*ebiten.Image{
		"planet1": loadImage("assets/planets/planet1.png"),
		"planet2": loadImage("assets/planets/planet2.png"),
	}
}

// This should be called when:
// resetting ship only,
// restore minerals to original state,
// not regenerating bodies,
// restart the same level if player crashed less than 5 times
func (game *Game) resetLevel() {
	game.GameState = shared.StateAiming
	game.Ship = game.generateShip()
	game.Bodies = deepCloneBodies(game.InitialBodies)
	game.CollectedMinerals = 0
}

func (game *Game) resetShipAfterCrash() {
	game.GameState = shared.StateAiming
	game.Ship = game.generateShip()
}

func (game *Game) generateShip() objects.Ship {
	return objects.Ship{
		Position: math2d.Vector2{
			X: float64(game.ScreenWidth) / 2,
			Y: float64(game.ScreenHeight) - 50,
		},
		Rotation:        0,
		MagnetRadius:    game.GameConfig.MagnetRadius,
		CollisionRadius: game.GameConfig.ShipCollisionRadius,
	}
}

func (game *Game) generateMinerals(body objects.Body, count int) []objects.Mineral {
	minerals := []objects.Mineral{}
	minDistance := 15.0

	for len(minerals) < count {
		angle := rand.Float64() * 2 * math.Pi
		// uniform distribution of minerals across the circle
		// minerals distribute evenly over the whole body area
		radius := math.Sqrt(rand.Float64()) * body.Size

		x := body.Position.X + radius*math.Cos(angle)
		y := body.Position.Y + radius*math.Sin(angle)

		valid := true

		for _, mineral := range minerals {
			dx := mineral.Position.X - x
			dy := mineral.Position.Y - y
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < minDistance {
				valid = false
				break
			}
		}

		if valid {
			minerals = append(minerals, objects.Mineral{Position: math2d.Vector2{X: x, Y: y}})
		}
	}

	return minerals
}

func (game *Game) collectMinerals() {
	for pi := range game.Bodies {
		body := &game.Bodies[pi]

		for mi := range body.Minerals {
			mineral := &body.Minerals[mi]

			if mineral.Collected {
				continue
			}

			// apply velocity and move mineral
			mineral.Position.X += mineral.Velocity.X
			mineral.Position.Y += mineral.Velocity.Y

			// damping (prevents infinite speed)
			mineral.Velocity.X *= game.GameConfig.MineralDamping
			mineral.Velocity.Y *= game.GameConfig.MineralDamping

			// recalculate distance (NEW position)
			dx := game.Ship.Position.X - mineral.Position.X
			dy := game.Ship.Position.Y - mineral.Position.Y

			// distance := math.Sqrt(dx*dx + dy*dy)
			distance := math2d.Distance(mineral.Position, game.Ship.Position)

			// ship's magnet pull strength and mineral's speed
			if distance > 0 && distance < game.Ship.MagnetRadius {
				nx := dx / distance
				ny := dy / distance

				shipSpeed := math2d.Speed(game.Ship.Velocity)

				// base magnet pull with smooth acceleration, becomes stronger when closer (far=weak pull, near=strong pull)
				strength := 1 - (distance / game.Ship.MagnetRadius) // 0 to 1

				// ship's magnetic pull (numbers are placeholders)
				pull := (game.GameConfig.MagnetPull + shipSpeed*0.1) * strength

				mineral.Velocity.X += nx * pull
				mineral.Velocity.Y += ny * pull

				// to prevent mineral lag when getting pulled by a fast travelling ship, faster ship = stronger magnet pul
				// 0.2 is the follow through strength to the ship, 0 = minerals will lag behind a faster ship, X = minerals will reach ship fast
				follow := game.GameConfig.MagnetFollowStrength
				mineral.Position.X += dx * follow
				mineral.Position.Y += dy * follow

				// dynamic max speed (3 is a placeholder for now)
				mineralMaxSpeed := shipSpeed + 3
				mineralSpeed := math.Sqrt(mineral.Velocity.X*mineral.Velocity.X + mineral.Velocity.Y*mineral.Velocity.Y)

				if mineralSpeed > mineralMaxSpeed {
					scale := mineralMaxSpeed / mineralSpeed
					mineral.Velocity.X *= scale
					mineral.Velocity.Y *= scale
				}
			}

			// collect mineral
			if distance < 10 {
				mineral.Collected = true
				game.CollectedMinerals++
			}
		}
	}
}

func (game *Game) checkWin() {
	destinationPlanet := game.getDestinationPlanet()

	distance := math2d.Distance(game.Ship.Position, destinationPlanet.Position)

	if distance < destinationPlanet.Size {
		game.getWinResult()
	}
}

func (game *Game) collectStars() {
	totalMinerals := getTotalMineralsInLevel(game.Bodies)
	allMineralsCollected := game.CollectedMinerals == totalMinerals
	noCrashes := game.CrashCount == 0

	// if is3stars {
	// 	game.Stars = 3 // All minerals + all lives
	// } else if is2stars {
	// 	game.Stars = 2 // Not all minerals + all lives
	// } else if is1star {
	// 	game.Stars = 1 // No minerals + not all lives
	// } else {
	// 	game.Stars = 0 // No minerals + 0 lives
	// }

	switch {
	case allMineralsCollected && noCrashes:
		game.Stars = 3
	case noCrashes:
		game.Stars = 2
	default:
		game.Stars = 1
	}
}

func (game *Game) updateMenu() {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	mx, my := ebiten.CursorPosition()

	// replay button
	if isPointInsideRect(mx, my, 100, 380, 200, 40) {
		game.resetLevel()
		game.LevelMenuState = shared.LevelMenuClosed
	}

	// next level button or continue button (keep playing)
	if isPointInsideRect(mx, my, 100, 440, 200, 40) {
		// go next level
		if game.LevelMenuState == shared.LevelMenuWin {
			game.nextLevel()
			game.LevelMenuState = shared.LevelMenuClosed
			return
		}

		// keep playing after crash/pause
		if game.LevelMenuState == shared.LevelMenuCrash ||
			game.LevelMenuState == shared.LevelMenuPause {
			game.LevelMenuState = shared.LevelMenuClosed
			return
		}
	}
}
