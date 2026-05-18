package render

import (
	"image/color"
	"math"

	"github.com/FBascou/rocket_game/entities"
	"github.com/FBascou/rocket_game/physics"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawDragIndicator(
	screen *ebiten.Image,
	ship entities.Ship,
	planets []entities.Planet,
	gameConfig DragIndicatorConfig,
	dragStartX int,
	dragStartY int,
) {
	cursorX, cursorY := ebiten.CursorPosition()

	// launch vector
	dx := float64(dragStartX - cursorX)
	dy := float64(dragStartY - cursorY)

	vx := dx * gameConfig.LaunchPower
	vy := dy * gameConfig.LaunchPower

	// use same clamp as actual launch
	vx, vy = physics.ClampVelocity(
		vx,
		vy,
		gameConfig.MaxLaunchSpeed,
	)

	// predicted x/y position (simulates a fake future path)
	px := ship.X
	py := ship.Y

	steps := gameConfig.DragLineSteps

	// simple dotted line (step-based) when dragging the ship
	for i := 0; i < steps; i++ {
		// fake gravity preview
		for _, planet := range planets {

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
			// force := planet.GravityStrength / (distance * distance + gameConfig.GravityFalloff)
			falloff := distance + gameConfig.GravityFalloff
			// force := planet.GravityStrength / falloff
			force := (planet.GravityStrength * gameConfig.DragPreviewGravityStrength) / falloff

			nx := dx / distance
			ny := dy / distance

			vx += nx * force
			vy += ny * force
		}

		// add friction
		vx *= gameConfig.ShipFriction
		vy *= gameConfig.ShipFriction

		// preview step multiplier to make dots more spaced
		previewStep := gameConfig.DragPreviewStep

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
