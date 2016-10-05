package game

import gc "github.com/rthornton128/goncurses"

// Play starts main part of the game.
func (g *Board) Play() {
	initScreen()
	g.initGame()
	play(g)
}

func initScreen() {
	stdscr := gc.StdScr()

	gc.InitPair(1, gc.C_BLACK, gc.C_WHITE)
	gc.InitPair(2, gc.C_BLACK, gc.C_CYAN)
	gc.InitPair(3, gc.C_BLACK, gc.C_MAGENTA)
	gc.InitPair(4, gc.C_BLACK, gc.C_GREEN)

	stdscr.Clear()
}

func play(board *Board) {
	stdscr := gc.StdScr()
	my, mx := stdscr.MaxYX()
	for board.continuePlaying() {
		stdscr.Clear()
		height, width := len(board.field), len(board.field[0])
		for i, row := range board.field {
			for j, value := range row {
				cnum := whichColor(value)
				stdscr.ColorOn(cnum)
				x, y := (mx-width)/2+j, (my-height)/2+i
				stdscr.MovePrint(y, x, whichChar(value))
				stdscr.ColorOff(cnum)
			}
		}
		stdscr.Refresh()
	}
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
