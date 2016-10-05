package menu

import gc "github.com/rthornton128/goncurses"

const (
	// Height of menu window
	HEIGHT = 10
	// Width of menu window
	WIDTH = 30
	// Margin of menu window
	MARGIN = 2
)

// Manager is a struct which manage game menu
type Manager struct {
	window *gc.Window
	y, x   int      // current windows position on terminal
	titles []string // titles of menu
	active int      // curent active menu item
	modes  []string // game modes names
	mode   int      // current mode
}

// Game interface which describe game that should run
// if user choose appropriate option
type Game interface {
	Play(int)
}

// Init standart ncurses screen
func (m *Manager) init() error {
	stdscr, err := gc.Init()
	if err != nil {
		return err
	}

	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)

	// Adjust the default mouse-click sensitivity to make it more responsive
	gc.MouseInterval(50)
	gc.MouseMask(gc.M_ALL, nil) // detect all mouse clicks

	my, mx := stdscr.MaxYX()

	m.titles = []string{"Play", "Mode: ", "Exit"}

	m.modes = []string{"Easy", "Midd", "Hard"}

	m.y, m.x = (my/2)-(HEIGHT/2), (mx/2)-(WIDTH/2)

	m.window, err = gc.NewWindow(HEIGHT, WIDTH, m.y, m.x)

	if err != nil {
		return err
	}

	return nil
}

// printMenu() function prints all menu items to window
func (m *Manager) printMenu() {
	m.window.Box(0, 0)
	for i, s := range m.titles {
		if i == 1 {
			s = s + m.modes[m.mode]
		}
		if i == m.active {
			m.window.AttrOn(gc.A_REVERSE)
			m.window.MovePrint(MARGIN+i, MARGIN, s)
			m.window.AttrOff(gc.A_REVERSE)
		} else {
			m.window.MovePrint(MARGIN+i, MARGIN, s)
		}
	}
}

// Identify menu item which user choose by means of mouse
func (m *Manager) getActive(my, mx int) int {
	row := my - m.y - MARGIN
	col := mx - m.x - MARGIN

	if row < 0 || row > len(m.titles)-1 {
		return -1
	}

	title := m.titles[row]

	if col >= 0 && col < len(title) {
		return row
	}
	return -1
}

// Handle pressed key
func (m *Manager) handleInput(key gc.Key) bool {
	switch key {
	case gc.KEY_LEFT:
		if m.active == 1 {
			if m.mode == 0 {
				m.mode = len(m.modes) - 1
			} else {
				m.mode--
			}
		}
	case gc.KEY_RIGHT:
		if m.active == 1 {
			if m.mode == len(m.modes)-1 {
				m.mode = 0
			} else {
				m.mode++
			}
		}
	case gc.KEY_UP:
		if m.active == 0 {
			m.active = len(m.titles) - 1
		} else {
			m.active--
		}
	case gc.KEY_DOWN:
		if m.active == len(m.titles)-1 {
			m.active = 0
		} else {
			m.active++
		}
	case gc.KEY_MOUSE:
		/* pull the mouse event off the queue */
		if md := gc.GetMouse(); md != nil {
			new := m.getActive(md.Y, md.X)
			if new != -1 {
				m.active = new
			}
		}
		fallthrough
	case gc.KEY_RETURN, gc.KEY_ENTER, gc.Key('\r'):
		return true
	}
	return false
}

// Refresh screen
func (m *Manager) refresh() {
	gc.StdScr().Clear()
	gc.StdScr().Refresh()
	m.window.Refresh()
}

// Run menu and give some choices for user before he begin game
func (m *Manager) Run(game Game) error {
	err := m.init()

	if err != nil {
		return err
	}
	defer gc.End()

	m.printMenu()
	m.refresh()

	var key gc.Key
	for key != 'q' {
		key = gc.StdScr().GetChar()
		if m.handleInput(key) {
			switch m.active {
			case 0:
				game.Play(m.mode)
			case 2: // exit
				return nil
			}
		}

		m.printMenu()
		m.refresh()
	}

	return nil
}
