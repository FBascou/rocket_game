package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input → changes state → Draw reads state

const dimensionWidth int = 400
const dimensionHeight int = 600
const outlineImageAlpha float32 = 0.1
var dragging bool
var startX, startY int

// runs every frame (~60 times per second)
func (game *Game) Update() error {
	if game.LevelMenuState != LevelMenuClosed {
		game.updateMenu()

		// close menu with P
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			game.LevelMenuState = LevelMenuClosed
		}

		return nil
	}

	if game.Won {
		// game.initializeNewLevel()
		game.LevelMenuState = LevelMenuWin
	} else if game.Lives == 0 {
		game.LevelMenuState = LevelMenuLose
		return nil
	} 

	// only for testing
	if ebiten.IsKeyPressed(ebiten.Key(ebiten.KeyR)) {
		game.initializeNewLevel()
	}

if inpututil.IsKeyJustPressed(ebiten.KeyP) {
	game.LevelMenuState = LevelMenuPause
}

	// start drag while !game.Launched prevents ship or physics to work
	if  !game.Launched && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		dragging = true
		startX, startY = ebiten.CursorPosition()
	}

	// end drag (release)
	if dragging && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		dragging = false

		endX, endY := ebiten.CursorPosition()

		// drag vector (direction + strength of launch)
		dx := float64(startX - endX)
		dy := float64(startY - endY)

		vx := dx * 0.1
		vy := dy * 0.1

		// limit ship's speed/launch energy (5 is a placeholder) if dragged too hard
		game.Ship.VX, game.Ship.VY = clampVelocity(vx, vy, 5)
		
		// now ship can move and physics are applied
		game.Launched = true
	}

	if game.Launched {
		game.applyGravity()
		game.applyFriction()

		// clamps all ship's speed, including if accelerated by gravity to 5
		// it's just here to test gameplay 
		// game.Ship.VX, game.Ship.VY = clampVelocity(game.Ship.VX, game.Ship.VY, 5)

		game.collectMinerals()
		game.checkWin()

		// ship's movement
		game.Ship.X += game.Ship.VX
		game.Ship.Y += game.Ship.VY

		// ship's speed (vector magnitude)
		speed := math.Sqrt(game.Ship.VX * game.Ship.VX + game.Ship.VY * game.Ship.VY)

		// ship's rotation
		if speed > 0.1 {
			game.Ship.Rotation = math.Atan2(game.Ship.VY, game.Ship.VX) + math.Pi / 2
		}
	}

	if game.Crashed {
		game.Lives--
		game.CrashCount++

		game.resetShipAfterCrash()

		if game.Lives == 0 {
			game.LevelMenuState = LevelMenuLose
		} else {
			game.LevelMenuState = LevelMenuPause
		}

		game.Crashed = false
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
	if dragging {
    cx, cy := ebiten.CursorPosition()

    // simple line (step-based) when dragging the ship
    steps := 20
    for i := 0; i < steps; i++ {
			t := float64(i) / float64(steps)

			// linear interpolation (lerp):
			// L=A(1−t)+Bt: A = start, B = end, t = how far between them (0 → 1)
			x := float64(startX) * (1 - t) + float64(cx) * t
			y := float64(startY) * (1 - t) + float64(cy) * t

			// sets white dots on the line
			screen.Set(int(x), int(y), color.White)
    }
	}

	// create a func for showing score
	if game.Won {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Minerals: %d, Lives: %d, Stars: %d", game.CollectedMinerals, game.Lives, game.Stars))
		// return
	} else {
		if game.Lives == 0 {
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

	game.initializeAssets()
	game.Ship = game.generateShip()
	game.Planets = game.generatePlanets()
	game.InitialPlanets = deepClonePlanets(game.Planets)
	game.Lives = 5
	
	game.initializeAssetOutlines() 
	
	// game.Menu = game.generateLevelMenu()
	ebiten.SetWindowSize(dimensionWidth, dimensionHeight)
  ebiten.SetWindowTitle("Ship Express (ShipEx/ShipX)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}