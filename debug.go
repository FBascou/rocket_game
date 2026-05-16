package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (game *Game) getDebugFloatVariables() []DebugVariable {
	return []DebugVariable{
		{"LaunchPower", 0.01, &game.GameConfig.LaunchPower},
		{"MaxLaunchSpeed", 0.2, &game.GameConfig.MaxLaunchSpeed},
		{"ShipFriction", 0.01, &game.GameConfig.ShipFriction},
		{"MaxShipSpeed", 0.2, &game.GameConfig.MaxShipSpeed},
		{"MagnetRadius", 5, &game.GameConfig.MagnetRadius},
		{"MagnetPull", 0.05, &game.GameConfig.MagnetPull},
		{"MagnetFollowStrength", 0.01, &game.GameConfig.MagnetFollowStrength},
		{"MineralDamping", 0.01, &game.GameConfig.MineralDamping},
		{"GravityFalloff", 5, &game.GameConfig.GravityFalloff},
		{"DragPreviewStep", 0.1, &game.GameConfig.DragPreviewStep},
		{"DragPreviewGravityStrength", 0.5, &game.GameConfig.DragPreviewGravityStrength},
	}
}

func (game *Game) getDebugIntVariables() []DebugIntVariable {
	return []DebugIntVariable{
		{"DragLineSteps", 1, &game.GameConfig.DragLineSteps},
	}
}

func (game *Game) updateDebugTuner() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		game.DebugMode = !game.DebugMode
	}

	if !game.DebugMode {
		return
	}

	total := len(game.getDebugFloatVariables()) + len(game.getDebugIntVariables())

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		game.DebugSelectedIndex++
		if game.DebugSelectedIndex >= total {
			game.DebugSelectedIndex = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		game.DebugSelectedIndex--
		if game.DebugSelectedIndex < 0 {
			game.DebugSelectedIndex = total - 1
		}
	}

	index := 0

	for _, variable := range game.getDebugFloatVariables() {
		if index == game.DebugSelectedIndex {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				*variable.Value += variable.Step
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				*variable.Value -= variable.Step
			}
		}
		index++
	}

	for _, variable := range game.getDebugIntVariables() {
		if index == game.DebugSelectedIndex {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				*variable.Value += variable.Step
			}

			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				*variable.Value -= variable.Step
			}
		}
		index++
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		game.printCurrentConfig()
	}
}

func (game *Game) drawDebugTuner(screen *ebiten.Image) {
	if !game.DebugMode {
		return
	}

	x := 550
	y := 20

	ebitenutil.DebugPrintAt(screen, "DEBUG TUNER", x, y)

	y += 20

	index := 0

	for _, variable := range game.getDebugFloatVariables() {
		prefix := "  "

		if index == game.DebugSelectedIndex {
			prefix = "> "
		}

		text := fmt.Sprintf("%s%s: %.2f",
			prefix,
			variable.Name,
			*variable.Value,
		)

		ebitenutil.DebugPrintAt(screen, text, x, y)

		y += 16
		index++
	}

	for _, variable := range game.getDebugIntVariables() {
		prefix := "  "

		if index == game.DebugSelectedIndex {
			prefix = "> "
		}

		text := fmt.Sprintf("%s%s: %d",
			prefix,
			variable.Name,
			*variable.Value,
		)

		ebitenutil.DebugPrintAt(screen, text, x, y)

		y += 16
		index++
	}
}

func (game *Game) printCurrentConfig() {
	fmt.Println("GameConfig{")
	fmt.Printf("    DragLineSteps: %d,\n", game.GameConfig.DragLineSteps)
	fmt.Printf("    DragPreviewStep: %.2f,\n", game.GameConfig.DragPreviewStep)
	fmt.Printf("    DragPreviewGravityStrength: %.2f,\n", game.GameConfig.DragPreviewGravityStrength)
	fmt.Printf("    LaunchPower: %.2f,\n", game.GameConfig.LaunchPower)
	fmt.Printf("    MaxLaunchSpeed: %.2f,\n", game.GameConfig.MaxLaunchSpeed)
	fmt.Printf("    ShipFriction: %.2f,\n", game.GameConfig.ShipFriction)
	fmt.Printf("    MaxShipSpeed: %.2f,\n", game.GameConfig.MaxShipSpeed)
	fmt.Printf("    MagnetRadius: %.2f,\n", game.GameConfig.MagnetRadius)
	fmt.Printf("    MagnetPull: %.2f,\n", game.GameConfig.MagnetPull)
	fmt.Printf("    MagnetFollowStrength: %.2f,\n", game.GameConfig.MagnetFollowStrength)
	fmt.Printf("    MineralDamping: %.2f,\n", game.GameConfig.MineralDamping)
	fmt.Printf("    GravityFalloff: %.2f,\n", game.GameConfig.GravityFalloff)
	fmt.Println("}")
}
