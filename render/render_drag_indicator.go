package render

import (
	"image/color"
	"math"

	"github.com/FBascou/rocket_game/math2d"
	"github.com/FBascou/rocket_game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawDragIndicator(
	screen *ebiten.Image,
	ship objects.Ship,
	bodies []objects.Body,
	gameConfig DragIndicatorConfig,
	dragStartX int,
	dragStartY int,
) {
	cursorX, cursorY := ebiten.CursorPosition()

	// launch vector
	dx := float64(dragStartX - cursorX)
	dy := float64(dragStartY - cursorY)

	// use same clamp as actual launch
	velocity := math2d.ClampMagnitude(
		math2d.Vector2{
			X: dx * gameConfig.LaunchPower,
			Y: dy * gameConfig.LaunchPower},
		gameConfig.MaxLaunchSpeed,
	)

	ship.Velocity = velocity

	// predicted x/y position (simulates a fake future path)
	px := ship.Position.X
	py := ship.Position.Y

	steps := gameConfig.DragLineSteps

	// simple dotted line (step-based) when dragging the ship
	for i := 0; i < steps; i++ {
		// fake gravity preview
		for _, body := range bodies {

			dx := body.Position.X - px
			dy := body.Position.Y - py

			distance := math.Sqrt(dx*dx + dy*dy)

			if distance < 1 {
				continue
			}

			// outside gravity field
			if distance > body.GravityRadius {
				continue
			}

			// 0.12 tuning constant means how strongly should the preview line bend visually to the body gravity (visual prediction)
			// force := body.GravityStrength / (distance * distance + gameConfig.GravityFalloff)
			falloff := distance + gameConfig.GravityFalloff
			// force := body.GravityStrength / falloff
			force := (body.GravityStrength * gameConfig.GravityMultiplier) / falloff

			nx := dx / distance
			ny := dy / distance

			ship.Velocity.X += nx * force
			ship.Velocity.Y += ny * force
		}

		// add friction
		ship.Velocity.X *= gameConfig.ShipFriction
		ship.Velocity.Y *= gameConfig.ShipFriction

		// preview step multiplier to make dots more spaced
		previewStep := gameConfig.DragPreviewStep

		// move prediction point (how/where each dot moves)
		// slow velocity = tightly packed dots
		// fast velocity = spaced dots
		// dots are still fairly dense because you're only advancing by 1 frame of movement per loop
		px += ship.Velocity.X * previewStep
		py += ship.Velocity.Y * previewStep

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
