package game

import (
	"fmt"
	"image/color"

	"github.com/FBascou/rocket_game/render"
	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// const outlineVectorAlpha float32 = 0.04

// runs every frame (~60 times per second)
func (game *Game) Update() error {

	// update GameConfig debugging
	game.updateDebugTuner()

	// open menu after some level event (crash, win, lose, etc.)
	if game.LevelMenuState != shared.LevelMenuClosed {
		game.updateMenu()

		// close menu with "P"
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.LevelMenuState = shared.LevelMenuClosed
		}

		return nil
	}

	// open menu with "P"
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		game.LevelMenuState = shared.LevelMenuPause
	}

	// restart game only for testing
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		game.LoadLevel(game.LevelNumber)
	}

	switch game.GameState {
	case shared.StateAiming:
		game.updateAiming()
	case shared.StateFlying:
		game.updateFlying()
	case shared.StateCrashed:
		game.updateCrashed()
	case shared.StateWon:
		game.updateWon()
	case shared.StateLost:
		game.updateLost()
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	// fill screen background color with black
	screen.Fill(color.Black)

	render.DrawShip(screen, game.Ship)
	render.DrawBodies(screen, game.Bodies)
	render.DrawMinerals(screen, game.Bodies)

	// Draw GameConfig debug
	game.drawDebugTuner(screen)

	// draw the line when dragging
	if game.Dragging {
		render.DrawDragIndicator(
			screen,
			game.Ship,
			game.Bodies,
			game.GetDragIndicatorConfig(),
			game.DragStartX,
			game.DragStartY,
		)
	}

	// create a func for showing score below
	if game.GameState == shared.StateWon {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Minerals: %d, Lives: %d, Stars: %d", game.CollectedMinerals, game.Lives, game.Stars))
	} else {
		if game.GameState == shared.StateLost {
			ebitenutil.DebugPrint(screen, "YOU LOST")
		} else {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("Minerals: %d, Lives: %d, Crash count: %d", game.CollectedMinerals, game.Lives, game.CrashCount))
		}
	}

	if game.LevelMenuState != shared.LevelMenuClosed {
		render.DrawLevelMenu(screen, game.LevelMenuState, game.LevelNumber, game.Stars, game.ScreenWidth, game.ScreenHeight)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.ScreenWidth, game.ScreenHeight
}
