package main

type LevelMenuState int

const (
	LevelMenuClosed LevelMenuState = iota
	LevelMenuPause
	LevelMenuCrash
	LevelMenuWin
	LevelMenuLose
)