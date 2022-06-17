package widgets

import (
	"github.com/NDRAEY/curses"
	. "github.com/NDRAEY/Bulucomon/types"
)

func TitledWindow(screen Screen, title string, width, height, x, y int) (*curses.Window, *curses.Window) {
	titlewnd := curses.NewWindow(3, width, y, x)
	titlewnd.Border(0,0,0,0,0,0,
				int(curses.ACS_VLINE()),
				int(curses.ACS_VLINE()))

	titlewnd.Mvaddstr(1, 1, title)

	win := curses.NewWindow(height-3, width, y+2, x)
	win.Box(0,0)

	titlewnd.Refresh()
	win.Refresh()

	return titlewnd, win
}
