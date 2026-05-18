package shared

type GameState int

const (
	StateAiming GameState = iota
	StateFlying
	StateCrashed
	StateWon
	StateLost
)

type LevelMenuState int

const (
	LevelMenuClosed LevelMenuState = iota
	LevelMenuPause
	LevelMenuCrash
	LevelMenuWin
	LevelMenuLose
)
