package main

type LevelMenuState int

const (
	LevelMenuClosed LevelMenuState = iota
	LevelMenuPause
	LevelMenuWin
	LevelMenuLose
)