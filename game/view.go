package game

import gc "github.com/rthornton128/goncurses"

const (
	cellScaleX = 3
	cellScaleY = 2
)

// Play starts main part of the game.
func (board *Board) Play() {
	initScreen()
	board.initGame()
	play(board)
}

func initScreen() {
	stdscr := gc.StdScr()

	gc.InitPair(1, gc.C_BLACK, gc.C_WHITE)
	gc.InitPair(2, gc.C_BLACK, gc.C_CYAN)
	gc.InitPair(3, gc.C_BLACK, gc.C_MAGENTA)
	gc.InitPair(4, gc.C_BLACK, gc.C_GREEN)

	stdscr.Clear()
}

type gameContract interface {
	choose(p *point)
}

func play(board *Board) {
	stdscr := gc.StdScr()
	for board.continuePlaying() {
		board.handleKey(stdscr)
		board.draw(stdscr)
	}
}

func (board *Board) draw(win *gc.Window) {
	win.Clear()
	my, mx := win.MaxYX()
	height, width := len(board.field), len(board.field[0])
	for i, row := range board.field {
		for j, value := range row {
			cnum := whichColor(value)
			win.ColorOn(cnum)
			x, y := (mx-width*cellScaleX)/2+j, (my-height*cellScaleY)/2+i
			win.MovePrint(y, x, whichChar(value))
			win.ColorOff(cnum)
		}
	}
	win.Refresh()
}

func (board *Board) handleKey(win *gc.Window) {
	// Stub. TODO: Implement
}

func whichColor(value *point) int16 {
	if !value.touched && !value.isBomb {
		return 1
	}
	if value.isBomb {
		return 3
	}
	return 4
}

func whichChar(value *point) rune {
	if !value.touched && !value.isBomb {
		return '█'
	}
	if value.isBomb {
		return '💣'
	}
	return rune(value.bombsNumber)
}
