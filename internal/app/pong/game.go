package pong

import (
	"fmt"
)

type Game struct {
	// arena *arena
	score  int
	isOver bool
}

func initScore() int {
	return 0
}

/**
func initArena() *arena {
	return newArena(initSnake(), pointsChan, 50, 100)
}
**/
func NewGame() *Game {
	return &Game{score: initScore()}
}

func (g *Game) Start() {
	fmt.Println("Starting pong!")
}
