package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		game.Ship.VX, game.Ship.VY = clampVelocity(vx, vy, game.GameConfig.MaxLaunchSpeed)

		// now ship can move and physics are applied
		game.GameState = StateFlying
	}
}

func (game *Game) updateFlying() {
	game.applyGravity()
	game.applyFriction()
	game.applySpeedDamping()

	// clamps all ship's speed, including if accelerated by gravity to 6
	// it's just here to test gameplay 
	// game.Ship.VX, game.Ship.VY = clampVelocity(game.Ship.VX, game.Ship.VY, 6)

	game.collectMinerals()
	game.checkWin()

	// ship's movement
	game.Ship.X += game.Ship.VX
	game.Ship.Y += game.Ship.VY

	// ship's speed (vector magnitude)
		speed := math.Sqrt(game.Ship.VX * game.Ship.VX + game.Ship.VY * game.Ship.VY)

	if speed > 0.1 {
		game.Ship.Rotation = math.Atan2(game.Ship.VY, game.Ship.VX) + math.Pi / 2
	}
}

func (game *Game) updateCrashed() {
	game.Lives--
	game.CrashCount++

	game.resetShipAfterCrash()

	if game.Lives <= 0 {
		game.GameState = StateLost
		game.LevelMenuState = LevelMenuLose
	} else {
		game.GameState = StateAiming
		game.LevelMenuState = LevelMenuCrash
	}
}

// per-frame behavior while in won state
func (game *Game) updateWon() {
	game.LevelMenuState = LevelMenuWin
}

func (game *Game) updateLost() {
	game.LevelMenuState = LevelMenuLose
}

// one-time transition
func (game *Game) getWinResult() {
	game.GameState = StateWon
	game.collectStars()
	game.LevelMenuState = LevelMenuWin
}