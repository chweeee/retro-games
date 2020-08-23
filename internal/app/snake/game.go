package snake

import (
	"github.com/nsf/termbox-go"
	"time"
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type Game struct {
	arena  *arena
	score  int
	isOver bool
}

func initScore() int {
	return 0
}

func initArena() *arena {
	return newArena(initSnake(), pointsChan, 50, 100)
}

func initSnake() *snake {
	return newSnake(
		RIGHT,
		[]coord{
			{x: 1, y: 1},
			{x: 1, y: 2},
			{x: 1, y: 3},
			{x: 1, y: 4},
		},
	)
}

func (g *Game) end() {
	g.isOver = true
}

// reset the game object. why not just return a NewGame()?
func (g *Game) retry() {
	g.arena = initArena()
	g.score = initScore()
	g.isOver = false
}

func (g *Game) addPoints(p int) {
	g.score += p
}

// the more you eat the faster snek mooves
func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}

// Creates a new game object
func NewGame() *Game {
	return &Game{arena: initArena(), score: initScore()}
}

// Methods: functions with defined receivers
// similar to a function of an object instance in OOP
func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	g.render()

	// some go concurrency shit, need to read up
	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	// this loops 5evaaa<3<3~~
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				d := keyToDirection(e.key)       // gets the direction from user input
				g.arena.snake.changeDirection(d) // change current direction of the snake
			case RETRY:
				g.retry()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.arena.moveSnake(); err != nil {
					g.end()
				}
			}

			if err := g.render(); err != nil {
				panic(err)
			}

			time.Sleep(g.moveInterval())
		}
	}
}
