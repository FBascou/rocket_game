package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input → changes state → Draw reads state

const dimensionWidth int = 400
const dimensionHeight int = 600
// const outlineVectorAlpha float32 = 0.04

// runs every frame (~60 times per second)
func (game *Game) Update() error {
	// open menu after some level event (crash, win, lose, etc.)
	if game.LevelMenuState != LevelMenuClosed {
		game.updateMenu()

		// close menu with "P"
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.LevelMenuState = LevelMenuClosed
		}

		return nil
	}

	// open menu with "P"
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		game.LevelMenuState = LevelMenuPause
	}

	// restart game only for testing
	if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyR)) {
		game.initializeNewLevel()
	}

	switch game.GameState {
		case StateAiming:
			game.updateAiming()
		case StateFlying:
			game.updateFlying()
		case StateCrashed:
			game.updateCrashed()
		case StateWon:
			game.updateWon()
		case StateLost:
			game.updateLost()
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	// fill screen background color with black
	screen.Fill(color.Black)

	game.drawShip(screen)
	game.drawPlanets(screen)
	game.drawMinerals(screen)

	// draw the line when dragging
	if game.Dragging {
		game.drawDragIndicator(screen)
	}

	// create a func for showing score below
	if game.GameState == StateWon {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Minerals: %d, Lives: %d, Stars: %d", game.CollectedMinerals, game.Lives, game.Stars))
	} else {
		if game.GameState == StateLost {
			ebitenutil.DebugPrint(screen, "YOU LOST")
		} else {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("Minerals: %d, Lives: %d, Crash count: %d", game.CollectedMinerals, game.Lives, game.CrashCount))
		}
	}

	if game.LevelMenuState != LevelMenuClosed {
		game.drawLevelMenu(screen)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return dimensionWidth, dimensionHeight
}

// Function that initializes the game and its assets
func main() {
	game := &Game{}
	game.GameConfig = GameConfig{
    DragLineSteps:        15,
		DragPreviewStep: 			2.0,
    LaunchPower:          0.1,
    MaxLaunchSpeed:       6,
    ShipFriction:         0.99,
    MaxShipSpeed:         7,
    MagnetRadius:         75,
    MagnetPull:           0.3,
    MagnetFollowStrength: 0.2,
    MineralDamping:       0.95,
    GravityFalloff:       50,
	}
	game.GameState = StateAiming
	game.initializeAssets()
	game.Ship = game.generateShip()
	game.Planets = game.generatePlanets()
	game.InitialPlanets = deepClonePlanets(game.Planets)
	game.Lives = 5
	
	// game.Menu = game.generateLevelMenu()
	ebiten.SetWindowSize(dimensionWidth, dimensionHeight)
  ebiten.SetWindowTitle("Ship Express (ShipEx/ShipX)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}