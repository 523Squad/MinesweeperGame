package game

import (
	"fmt"

	gc "github.com/rthornton128/goncurses"
)

const (
	cellScaleX = 3
	cellScaleY = 2
	winWidth   = 80
	winHeight  = 40
)

type gameContract interface {
	choose(colCoord int, rowCoord int)
	performLeftClick(coolCord int, rowCoord int)
}

type coordinate struct {
	x, y, width, height int
}

type viewState struct {
	win      *gc.Window
	board    *Board
	position *coordinate
	elapsed  int
}

// Play starts main part of the game.
func (board *Board) Play(level int) {
	win, err := initScreen()
	if err != nil {
		panic(err)
	}
	board.initGame(level)
	state := &viewState{
		board: board,
		win:   win,
		position: &coordinate{
			width:  board.dimension,
			height: board.dimension,
		},
	}
	play(state)
	state.win.Delete()
}

// Reset resets board state.
func (board *Board) Reset() {
	board.resetGame()
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
	gc.InitPair(6, gc.C_CYAN, gc.C_BLACK)

	stdscr.Clear()
	my, mx := stdscr.MaxYX()
	y, x := (my-winHeight)/2, (mx-winWidth)/2
	win, err := gc.NewWindow(winHeight, winWidth, y, x)
	if err != nil {
		return nil, err
	}
	win.Keypad(true)
	// Update every second to approx. timer better.
	win.Timeout(1000)
	return win, nil
}

func play(state *viewState) {
	state.draw()
	for state.board.continuePlaying() {
		if state.handleKey() {
			break
		}
		state.draw()
		state.elapsed++
	}
	if state.board.gameWin {
		state.drawWin()
	}
	if state.board.gameOver {
		state.drawLoss()
	}
}

func (state *viewState) draw() {
	win := state.win
	win.Clear()
	win.Box(0, 0)
	my, mx := win.MaxYX()
	board := state.board
	height, width := state.board.dimension, state.board.dimension
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
	// Draw time elapsed
	x, y := (mx-width*cellScaleX)/2, (my-height*cellScaleY)/2-2
	win.Move(y, x)
	win.Printf("%03d", state.elapsed)
	// Draw flags left
	x, y = (mx+width*cellScaleX)/2-5, (my-height*cellScaleY)/2-2
	win.Move(y, x)
	win.Printf("%03d", state.board.flagsLeft)
	win.Refresh()
}

func (state *viewState) drawLoss() {
	state.draw()
	win := state.win
	my, mx := win.MaxYX()
	x, y := mx/2-6, (my-state.board.dimension*cellScaleY)/2-4
	win.Move(y, x)
	win.Printf("GAME OVER")
	state.waitForExit()
}

func (state *viewState) drawWin() {
	state.draw()
	win := state.win
	my, mx := win.MaxYX()
	x, y := mx/2-4, (my-state.board.dimension*cellScaleY)/2-4
	win.Move(y, x)
	win.Printf("YOU WON")
	state.waitForExit()
}

func (state *viewState) waitForExit() {
	for {
		switch state.win.GetChar() {
		case 'q', gc.KEY_RETURN, gc.KEY_ENTER, gc.Key('\r'):
			return
		}
	}
}

func (state *viewState) handleKey() bool {
	switch state.win.GetChar() {
	case gc.Key('q'):
		return true
	case 'w', gc.KEY_UP:
		state.position.up()
	case 'd', gc.KEY_RIGHT:
		state.position.right()
	case 's', gc.KEY_DOWN:
		state.position.down()
	case 'a', gc.KEY_LEFT:
		state.position.left()
	case gc.KEY_RETURN, gc.KEY_ENTER, gc.Key('\r'):
		state.board.choose(state.position.y, state.position.x)
	case 'f', gc.KEY_TAB:
		state.board.flag(state.position.y, state.position.x)
	}
	return false
}

func (state *viewState) whichColor(value *point, c *coordinate) int16 {
	if state.position.x == c.x && state.position.y == c.y {
		return 5
	}
	if value.hasFlag {
		return 6
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
	if value.hasFlag {
		return 'F'
	}
	if !value.touched {
		return '#'
	}
	if value.isBomb {
		return 'B'
	}
	if value.bombsNumber == 0 {
		return '0'
	}
	return int16(value.bombsNumber + '0')
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
