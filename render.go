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
	alpha float32,
) {
	r, g, b, _ := clr.RGBA()

	fillColor := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(alpha * 255),
	}

	vector.FillCircle(
		screen,
		float32(x),
		float32(y),
		float32(radius),
		fillColor,
		false,
	)
}

func drawShipVector(screen *ebiten.Image, ship Ship) {
	size := float32(12)

	x := float32(ship.X)
	y := float32(ship.Y)

	frontX := x + float32(math.Cos(ship.Rotation - math.Pi / 2)) * size
	frontY := y + float32(math.Sin(ship.Rotation - math.Pi / 2)) * size

	leftX := x + float32(math.Cos(ship.Rotation + math.Pi * 0.75)) * size
	leftY := y + float32(math.Sin(ship.Rotation + math.Pi * 0.75)) * size

	rightX := x + float32(math.Cos(ship.Rotation - math.Pi * 0.75)) * size
	rightY := y + float32(math.Sin(ship.Rotation - math.Pi * 0.75)) * size

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
		outlineVectorAlpha,
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
			planet.Gravity,
			color.RGBA{0, 0, 30, 30},
			outlineVectorAlpha,
		)

		vector.FillCircle(
			screen,
			float32(planet.X),
			float32(planet.Y),
			float32(planet.Radius),
			color.RGBA{30, 30, 255, 255},
			false,
		)
	} else {
		drawFilledCircle(
			screen,
			planet.X,
			planet.Y,
			planet.Gravity,
			color.RGBA{30, 0, 0, 30},
			outlineVectorAlpha,
		)

		vector.FillCircle(
			screen,
			float32(planet.X),
			float32(planet.Y),
			float32(planet.Radius),
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
	// missing title for crashed
	title := "PAUSED"

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
	// TODO: if LevelMenuLost should be a restart button and hide replay button
	if game.LevelMenuState == LevelMenuWin {
		game.drawButton(screen, 100, 440, 200, 40, "NEXT LEVEL")
	} else {
		game.drawButton(screen, 100, 440, 200, 40, "KEEP PLAYING")
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
					speed := math.Sqrt(m.VX * m.VX + m.VY * m.VY)
					brightness := uint8(math.Min(255, speed * 50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(m.X) + dx, int(m.Y) + dy, col)
				}
			}
		}
	}
}
