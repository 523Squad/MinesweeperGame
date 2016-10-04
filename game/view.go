package game

import (
	"log"

	gc "github.com/rthornton128/goncurses"
)

const (
	NOT_VISITED = -1
	MINE        = -2
	PLAYER      = -3
)

type State struct {
	Field [][]int
}

type MineSweeper interface {
	// Tells whether we should continue the game.
	Continue() bool
}

func (g *State) Play() {
	if !initScreen() {
		return
	}
	defer gc.End()

	play(g)
}

func initScreen() bool {
	stdscr, err := gc.Init()
	stdscr.Timeout(100)
	if err != nil {
		return false
	}
	if err := gc.StartColor(); err != nil {
		log.Fatal(err)
	}
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)

	gc.InitPair(1, gc.C_BLACK, gc.C_CYAN)
	gc.InitPair(2, gc.C_BLACK, gc.C_MAGENTA)
	gc.InitPair(3, gc.C_BLACK, gc.C_GREEN)

	stdscr.Clear()
	stdscr.Keypad(true)
	return true
}

func play(s *State) {
	if _, ok := s.(menu.Game); !ok {
		return
	}
	s = s.(menu.Game)
	stdscr := gc.StdScr()
	my, mx := stdscr.MaxYX()
	for s.Continue() {
		stdscr.Clear()
		height, width := len(s.Field), len(s.Field[0])
		for i, row := range s.Field {
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

func whichColor(value int) int16 {
	switch value {
	case NOT_VISITED:
		return 0
	case PLAYER:
		return 1
	case MINE:
		return 2
	}
	return 3
}

func whichChar(value int) rune {
	switch value {
	case NOT_VISITED:
		return '█'
	case PLAYER:
		return 'Ϙ'
	case MINE:
		return '💣'
	}
	return rune(value)
}
