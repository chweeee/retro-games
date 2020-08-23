package snake

import (
	"math/rand"
	"time"
)

type arena struct {
	food       *food
	snake      *snake
	hasFood    func(*arena, coord) bool
	height     int
	width      int
	pointsChan chan (int)
}

func newArena(s *snake, pc chan (int), h, w int) *arena {
	rand.Seed(time.Now().UnixNano())
	newArena := &arena{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: pc,
		hasFood:    hasFood,
	}

	newArena.placeFood()
	return newArena
}

func (a *arena) placeFood() {
	var x, y int

	for {
		x = rand.Intn(a.width)
		y = rand.Intn(a.height)

		// before we initialise the food, check that snake is not on the cell
		if !a.isOccupied(coord{x: x, y: y}) {
			break
		}
	}

	a.food = newFood(x, y)
}

// checks if the current position has any food
func hasFood(a *arena, c coord) bool {
	return c.x == a.food.x && c.y == a.food.y
}

func (a *arena) isOccupied(c coord) bool {
	return a.snake.isOnPosition(c)
}

// move the freaking snake in the arena
func (a *arena) moveSnake() error {
	if err := a.snake.move(); err != nil {
		return err
	}

	// snake dies if it goes outta bounds
	if a.snakeLeftArena() {
		return a.snake.die()
	}

	// if there's food where snake's head is at
	// when we consume food, we increase length of the snek
	if a.hasFood(a, a.snake.head()) {
		go a.addPoints(a.food.points)
		a.snake.length++
		a.placeFood()
	}

	return nil
}

// TODO: change this definition
func (a *arena) snakeLeftArena() bool {
	h := a.snake.head()
	return h.x > a.width || h.y > a.height || h.x < 0 || h.y < 0
}

func (a *arena) addPoints(p int) {
	a.pointsChan <- p
}
