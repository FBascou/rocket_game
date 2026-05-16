package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Renderer interface {
	DrawPlanet(screen *ebiten.Image, p Planet)
	DrawShip(screen *ebiten.Image, s Ship)
}

type RenderStyle struct {
	ShipColor        color.Color
	PlanetColor      color.Color
	DestinationColor color.Color
	OutlineColor     color.Color
	BackgroundColor  color.Color
}

func drawFilledCircle(
	screen *ebiten.Image,
	x, y, radius float64,
	clr color.Color,
) {
	vector.FillCircle(
		screen,
		float32(x),
		float32(y),
		float32(radius),
		clr,
		false,
	)
}

func drawShipVector(screen *ebiten.Image, ship Ship) {
	size := float32(12)

	x := float32(ship.X)
	y := float32(ship.Y)

	frontX := x + float32(math.Cos(ship.Rotation-math.Pi/2))*size
	frontY := y + float32(math.Sin(ship.Rotation-math.Pi/2))*size

	leftX := x + float32(math.Cos(ship.Rotation+math.Pi*0.75))*size
	leftY := y + float32(math.Sin(ship.Rotation+math.Pi*0.75))*size

	rightX := x + float32(math.Cos(ship.Rotation-math.Pi*0.75))*size
	rightY := y + float32(math.Sin(ship.Rotation-math.Pi*0.75))*size

	vector.StrokeLine(screen, frontX, frontY, leftX, leftY, 2, color.White, false)
	vector.StrokeLine(screen, leftX, leftY, rightX, rightY, 2, color.White, false)
	vector.StrokeLine(screen, rightX, rightY, frontX, frontY, 2, color.White, false)
}

// Draw the ship image and its magnet outline
// GeoM order matters: Translate(center of body) > Scale > Rotate > Translate(world position)
func (game *Game) drawShip(screen *ebiten.Image) {
	drawFilledCircle(
		screen,
		game.Ship.X,
		game.Ship.Y,
		game.Ship.MagnetRadius,
		color.RGBA{0, 30, 30, 30},
	)

	drawShipVector(screen, game.Ship)
}

// Draw the planet image and its gravity outline
func drawPlanetVector(screen *ebiten.Image, planet Planet) {
	if planet.IsDestination {
		drawFilledCircle(
			screen,
			planet.X,
			planet.Y,
			planet.GravityRadius,
			color.RGBA{0, 0, 30, 30},
		)

		vector.FillCircle(
			screen,
			float32(planet.X),
			float32(planet.Y),
			float32(planet.Size),
			color.RGBA{30, 30, 255, 255},
			false,
		)
	} else {
		drawFilledCircle(
			screen,
			planet.X,
			planet.Y,
			planet.GravityRadius,
			color.RGBA{30, 0, 0, 30},
		)

		vector.FillCircle(
			screen,
			float32(planet.X),
			float32(planet.Y),
			float32(planet.Size),
			color.RGBA{255, 30, 30, 255},
			false,
		)
	}
}

func (game *Game) drawPlanets(screen *ebiten.Image) {
	for _, planet := range game.Planets {
		drawPlanetVector(screen, planet)
	}
}

func (game *Game) drawLevelMenu(screen *ebiten.Image) {
	// dark transparent overlay
	vector.FillRect(
		screen,
		0,
		0,
		float32(dimensionWidth),
		float32(dimensionHeight),
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

	if game.LevelMenuState == LevelMenuCrash {
		title = "SHIP DESTROYED"
	}

	if game.LevelMenuState == LevelMenuWin {
		title = "LEVEL COMPLETE"
	}

	if game.LevelMenuState == LevelMenuLose {
		title = "YOU LOST"
	}

	ebitenutil.DebugPrintAt(screen, title, 120, 180)

	// stats
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Stars: %d", game.Stars),
		120,
		220,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Level: %d", game.LevelNumber),
		120,
		250,
	)

	// fake buttons for now
	game.drawButton(screen, 100, 320, 200, 40, "SETTINGS")
	game.drawButton(screen, 100, 380, 200, 40, "REPLAY")

	switch game.LevelMenuState {
	case LevelMenuWin:
		// if win
		game.drawButton(screen, 100, 440, 200, 40, "NEXT LEVEL")
	case LevelMenuCrash:
		// if crash
		game.drawButton(screen, 100, 440, 200, 40, "KEEP PLAYING")
	default:
		// if pause
		game.drawButton(screen, 100, 440, 200, 40, "RESUME")
	}
	// game.drawButton(screen, 100, 440, 200, 40, "X")
}

// Draw the minerals image
func (game *Game) drawMinerals(screen *ebiten.Image) {

	for _, p := range game.Planets {
		for _, m := range p.Minerals {
			if m.Collected {
				continue
			}

			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					speed := math.Sqrt(m.VX*m.VX + m.VY*m.VY)
					brightness := uint8(math.Min(255, speed*50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(m.X)+dx, int(m.Y)+dy, col)
				}
			}
		}
	}
}

func (game *Game) drawDragIndicator(screen *ebiten.Image) {
	cursorX, cursorY := ebiten.CursorPosition()

	// launch vector
	dx := float64(game.DragStartX - cursorX)
	dy := float64(game.DragStartY - cursorY)

	vx := dx * game.GameConfig.LaunchPower
	vy := dy * game.GameConfig.LaunchPower

	// use same clamp as actual launch
	vx, vy = clampVelocity(
		vx,
		vy,
		game.GameConfig.MaxLaunchSpeed,
	)

	// predicted x/y position (simulates a fake future path)
	px := game.Ship.X
	py := game.Ship.Y

	steps := game.GameConfig.DragLineSteps

	// simple dotted line (step-based) when dragging the ship
	for i := 0; i < steps; i++ {
		// fake gravity preview
		for _, planet := range game.Planets {

			dx := planet.X - px
			dy := planet.Y - py

			distance := math.Sqrt(dx*dx + dy*dy)

			if distance < 1 {
				continue
			}

			// outside gravity field
			if distance > planet.GravityRadius {
				continue
			}

			// 0.12 tuning constant means how strongly should the preview line bend visually to the planet gravity (visual prediction)
			// force := planet.GravityStrength / (distance * distance + game.GameConfig.GravityFalloff)
			falloff := distance + game.GameConfig.GravityFalloff
			// force := planet.GravityStrength / falloff
			force := (planet.GravityStrength * game.GameConfig.DragPreviewGravityStrength) / falloff

			nx := dx / distance
			ny := dy / distance

			vx += nx * force
			vy += ny * force
		}

		// add friction
		vx *= game.GameConfig.ShipFriction
		vy *= game.GameConfig.ShipFriction

		// preview step multiplier to make dots more spaced
		previewStep := game.GameConfig.DragPreviewStep

		// move prediction point (how/where each dot moves)
		// slow velocity = tightly packed dots
		// fast velocity = spaced dots
		// dots are still fairly dense because you're only advancing by 1 frame of movement per loop
		px += vx * previewStep
		py += vy * previewStep

		// t = 0 means the first dot, whilst t = 1 is the last dot
		// thickness grows toward drag end
		t := float64(i) / float64(steps)

		// radius of drag line dots
		// 2 means radius near ship (minimum dot radius), 1 means extra growth amount
		radius := 2 - t*1

		// transparency - first dots have less opacity, last dots have more opacity
		alpha := uint8(120 + t*135)

		drawFilledCircle(
			screen,
			px,
			py,
			radius,
			color.RGBA{255, 255, 255, alpha},
		)
	}
}
