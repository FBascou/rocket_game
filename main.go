package main

import (
	_ "image/png"
	"log"

	"github.com/FBascou/rocket_game/game"
	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

// Function that initializes the game and its assets
func main() {
	g := &game.Game{}
	g.ScreenWidth = game.ScreenWidth
	g.ScreenHeight = game.ScreenHeight
	g.GameConfig = game.GameConfig{
		DragLineSteps:        18,
		DragPreviewStep:      1.0,
		GravityMultiplier:    10.0,
		LaunchPower:          0.08,
		MaxLaunchSpeed:       4,
		ShipFriction:         0.985,
		MaxShipSpeed:         4.5,
		MagnetRadius:         75,
		MagnetPull:           0.3,
		MagnetFollowStrength: 0.2,
		MineralDamping:       0.95,
		GravityFalloff:       50,
	}
	g.GameState = shared.StateAiming
	g.Lives = 5
	g.InitializeAssets()
	g.LoadLevel(0)

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Ship Express (ShipEx/ShipX)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
