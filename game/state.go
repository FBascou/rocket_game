package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/FBascou/rocket_game/math2d"
	"github.com/FBascou/rocket_game/physics"
	"github.com/FBascou/rocket_game/shared"
)

func (game *Game) updateAiming() {
	// start drag while !game.Launched prevents ship or physics to work
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		game.Dragging = true
		game.DragStartX, game.DragStartY = ebiten.CursorPosition()
	}

	// end drag (release)
	if game.Dragging && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		game.Dragging = false

		endX, endY := ebiten.CursorPosition()

		// drag vector (direction + strength of launch)
		dx := float64(game.DragStartX - endX)
		dy := float64(game.DragStartY - endY)

		vx := dx * game.GameConfig.LaunchPower
		vy := dy * game.GameConfig.LaunchPower

		// limit ship's initial speed/launch energy (6 is a placeholder) if dragged too hard
		// game.Ship.Velocity.X, game.Ship.Velocity.Y
		game.Ship.Velocity = math2d.ClampMagnitude(math2d.Vector2{X: vx, Y: vy}, game.GameConfig.MaxLaunchSpeed)

		// now ship can move and physics are applied
		game.GameState = shared.StateFlying
	}
}

func (game *Game) updateFlying() {
	hasShipCrashed := physics.ApplyGravity(
		&game.Ship,
		game.Bodies,
		game.GameConfig.GravityFalloff,
		game.GameConfig.GravityMultiplier,
	)

	if hasShipCrashed {
		game.GameState = shared.StateCrashed
		return
	}

	physics.ApplyFriction(
		&game.Ship.Velocity,
		game.GameConfig.ShipFriction,
	)

	physics.ApplySpeedDamping(
		&game.Ship.Velocity,
		game.GameConfig.MaxShipSpeed,
	)

	// clamps all ship's speed, including if accelerated by gravity to 6
	// it's just here to test gameplay
	// game.Ship.Velocity.X, game.Ship.Velocity.Y = clampVelocity(game.Ship.Velocity.X, game.Ship.Velocity.Y, 6)

	game.collectMinerals()
	game.checkWin()

	// ship's movement
	game.Ship.Position.X += game.Ship.Velocity.X
	game.Ship.Position.Y += game.Ship.Velocity.Y

	// ship's speed (vector magnitude)
	speed := math.Sqrt(game.Ship.Velocity.X*game.Ship.Velocity.X + game.Ship.Velocity.Y*game.Ship.Velocity.Y)

	if speed > 0.1 {
		game.Ship.Rotation = math.Atan2(game.Ship.Velocity.Y, game.Ship.Velocity.X) + math.Pi/2
	}
}

func (game *Game) updateCrashed() {
	game.Lives--
	game.CrashCount++

	game.resetShipAfterCrash()

	if game.Lives <= 0 {
		game.GameState = shared.StateLost
		game.LevelMenuState = shared.LevelMenuLose
	} else {
		game.GameState = shared.StateAiming
		game.LevelMenuState = shared.LevelMenuCrash
	}
}

// per-frame behavior while in won state
func (game *Game) updateWon() {
	game.LevelMenuState = shared.LevelMenuWin
}

func (game *Game) updateLost() {
	game.LevelMenuState = shared.LevelMenuLose
}

// one-time transition
func (game *Game) getWinResult() {
	game.GameState = shared.StateWon
	game.collectStars()
	game.LevelMenuState = shared.LevelMenuWin
}
