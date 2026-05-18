package render

import (
	"fmt"
	"image/color"

	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawLevelMenu(
	screen *ebiten.Image,
	levelMenuState shared.LevelMenuState,
	levelNumber int,
	stars int,
	screenWidth int,
	screenHeight int,
	// drawButton func(
	// 	screen *ebiten.Image,
	// 	x, y, w, h float64,
	// 	label string,
	// ),
) {
	// dark transparent overlay
	vector.FillRect(
		screen,
		0,
		0,
		float32(screenWidth),
		float32(screenHeight),
		color.RGBA{0, 0, 0, 180},
		false,
	)

	// menu panel
	panelX := 50.0
	panelY := 150.0
	panelW := 300.0
	panelH := 300.0

	vector.FillRect(
		screen,
		float32(panelX),
		float32(panelY),
		float32(panelW),
		float32(panelH),
		color.RGBA{80, 80, 80, 255},
		false, // anti-aliasing
	)

	// title of menu
	title := "PAUSED"

	switch levelMenuState {
	case shared.LevelMenuCrash:
		title = "SHIP DESTROYED"
	case shared.LevelMenuWin:
		title = "LEVEL COMPLETE"
	case shared.LevelMenuLose:
		title = "YOU LOST"
	}

	ebitenutil.DebugPrintAt(screen, title, 120, 180)

	// stats
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Stars: %d", stars),
		120,
		220,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Level: %d", levelNumber),
		120,
		250,
	)

	// fake buttons for now
	DrawButton(screen, 100, 320, 200, 40, "SETTINGS")
	DrawButton(screen, 100, 380, 200, 40, "REPLAY")

	switch levelMenuState {
	case shared.LevelMenuWin:
		// if win
		DrawButton(screen, 100, 440, 200, 40, "NEXT LEVEL")
	case shared.LevelMenuCrash:
		// if crash
		DrawButton(screen, 100, 440, 200, 40, "KEEP PLAYING")
	default:
		// if pause
		DrawButton(screen, 100, 440, 200, 40, "RESUME")
	}
	// shared.DrawButton(screen, 100, 440, 200, 40, "X")
}
