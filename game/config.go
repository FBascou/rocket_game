package game

import (
	"github.com/FBascou/rocket_game/entities"
	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	Ships        map[string]*ebiten.Image
	Bodies       map[string]*ebiten.Image
	MineralImage *ebiten.Image
}

type DebugVariable struct {
	Name  string
	Step  float64
	Value *float64
}

type DebugIntVariable struct {
	Name  string
	Step  int
	Value *int
}

type GameConfig struct {
	// Drag line dot amount
	DragLineSteps int
	// Drag line dot space multiplier
	DragPreviewStep float64
	// Gravity Strength affects the direction drag preview
	DragPreviewGravityStrength float64

	// Ship's initial launch speed
	LaunchPower    float64
	MaxLaunchSpeed float64

	// Ship's speed
	ShipFriction float64
	MaxShipSpeed float64

	// Ship's magnet strength
	MagnetRadius         float64
	MagnetPull           float64
	MagnetFollowStrength float64
	MineralDamping       float64

	// Planet gravity
	GravityFalloff float64
}

type Level struct {
	screenWidth  int
	screenHeight int
}

type Game struct {
	Assets             Assets
	Level              Level
	LevelMenuState     shared.LevelMenuState
	GameState          shared.GameState
	GameConfig         GameConfig
	Ship               entities.Ship
	InitialBodies      []entities.Body
	Bodies             []entities.Body
	LevelNumber        int
	CollectedMinerals  int
	CrashCount         int
	Lives              int
	Stars              int
	Dragging           bool
	DragStartX         int
	DragStartY         int
	DebugMode          bool
	DebugSelectedIndex int
	ScreenWidth        int
	ScreenHeight       int
}
