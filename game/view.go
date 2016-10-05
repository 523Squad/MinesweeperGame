package game

import (
	"fmt"

	gc "github.com/rthornton128/goncurses"
)

const (
	cellScaleX = 3
	cellScaleY = 2
	winWidth   = 50
	winHeight  = 30
)

type gameContract interface {
	choose(x, y int)
}

type coordinate struct {
	x, y, width, height int
}

type viewState struct {
	win      *gc.Window
	board    *Board
	position *coordinate
}

// Play starts main part of the game.
func (board *Board) Play(level int) {
	win, err := initScreen()
	if err != nil {
		panic(err)
	}
	board.initGame()
	state := &viewState{
		board: board,
		win:   win,
		position: &coordinate{
			width:  board.dimension,
			height: board.dimension,
		},
	}
	play(state)
}

func initScreen() (*gc.Window, error) {
	stdscr := gc.StdScr()

	if err := gc.StartColor(); err != nil {
		return nil, fmt.Errorf("Init color fail: %v", err)
	}

	gc.InitPair(1, gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(2, gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(3, gc.C_MAGENTA, gc.C_BLACK)
	gc.InitPair(4, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(5, gc.C_BLACK, gc.C_YELLOW)

	stdscr.Clear()
	my, mx := stdscr.MaxYX()
	y, x := (my-winHeight)/2, (mx-winWidth)/2
	win, err := gc.NewWindow(winHeight, winWidth, y, x)
	if err != nil {
		return nil, err
	}
	return win, nil
}

func play(state *viewState) {
	for state.board.continuePlaying() {
		state.draw()
		if state.handleKey() {
			break
		}
	}
	state.board = &Board{}
}

func (state *viewState) draw() {
	win := state.win
	win.Clear()
	win.Box(0, 0)
	my, mx := win.MaxYX()
	board := state.board
	height, width := len(board.field), len(board.field[0])
	for i, row := range board.field {
		for j, value := range row {
			pos := &coordinate{x: j, y: i}
			cnum := state.whichColor(value, pos)
			win.ColorOn(cnum)
			x, y := (mx-width*cellScaleX)/2+j*cellScaleX, (my-height*cellScaleY)/2+i*cellScaleY
			win.Move(y, x)
			win.Printf("%c ", state.whichChar(value, pos))
			win.ColorOff(cnum)
		}
	}
	win.Refresh()
}

func (state *viewState) handleKey() bool {
	p := state.position
	switch state.win.GetChar() {
	case 'q':
		return true
	case 'w' | gc.KEY_UP:
		//		if p.y > 0 {
		p.y--
		//		}
	case 'd' | gc.KEY_RIGHT:
		//		if p.x < p.width-1 {
		p.x++
		//		}
	case 's' | gc.KEY_DOWN:
		//		if p.y < p.height-1 {
		p.y++
		//		}
	case gc.KEY_LEFT:
		//		if p.x > 0 {
		p.x--
		//		}
		// case gc.KEY_ENTER:
		// 	state.board.choose(state.position.x, state.position.y)
	}
	return false
}

func (state *viewState) whichColor(value *point, c *coordinate) int16 {
	if state.position.x == c.x && state.position.y == c.y {
		return 5
	}
	if !value.touched {
		return 2
	}
	if value.isBomb {
		return 3
	}
	return 4
}

func (state *viewState) whichChar(value *point, c *coordinate) int16 {
	if state.position.x == c.x && state.position.y == c.y {
		return '&'
	}
	if !value.touched {
		return '#' // '█'
	}
	if value.isBomb {
		return 'B' // '💣'
	}
	return int16(value.bombsNumber /* + '0'*/)
}

func (p *coordinate) up() {
	if p.y > 0 {
		p.y--
	}
}

func (p *coordinate) right() {
	if p.x < p.width-1 {
		p.x++
	}
}

func (p *coordinate) down() {
	if p.y < p.height-1 {
		p.y++
	}
}

func (p *coordinate) left() {
	if p.x > 0 {
		p.x--
	}
}
