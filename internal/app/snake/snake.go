package snake

import (
	"errors"
)

type snake struct {
	body      []coord
	direction directionType
	length    int
}

func newSnake(d directionType, b []coord) *snake {
	return &snake{
		body:      b,
		direction: d,
		length:    len(b),
	}
}

// returns the coordinate of the head of the snake
func (s *snake) head() coord {
	return s.body[len(s.body)-1]
}

// generic function to check if snake is in a particular position
func (s *snake) isOnPosition(c coord) bool {
	// for-each range loop: useful for looping over arrays
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}

func (s *snake) die() error {
	return errors.New("Died")
}

/**
Main logic of moving the snek:
Move snek based on its current direction.
**/
func (s *snake) move() error {
	head := s.head()
	c := coord{x: head.x, y: head.y}

	switch s.direction {
	case RIGHT:
		c.x++
	case UP:
		c.y++
	case LEFT:
		c.x--
	case DOWN:
		c.y--
	}

	// snek ded if it touches itself
	if s.isOnPosition(c) {
		return s.die()
	}

	if s.length > len(s.body) {
		// snek has consumed food at coord c
		s.body = append(s.body, c)
	} else {
		// main magic: remove the tail of the snake, coord c becomes new head
		s.body = append(s.body[1:], c)
	}

	return nil
}

/**
func (s *snake) changeDirection(d directionType) {
	// Golang map type ==> hashtable: map[KeyType]ValueType
	opposites := map[directionType]directionType{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	// disallow snek to go into itself
	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}
**/

func (s *snake) changeDirection(d directionType) {
	s.direction = d
}
