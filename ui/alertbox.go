package widgets;

import (
	"github.com/NDRAEY/curses"
	. "github.com/NDRAEY/Bulucomon/types"
	"strings"
)

func AlertBox(screen Screen, title, text string, width, height int) {
	tw, th := (screen.Width-width)/2, (screen.Height-height)/2

	alrth := curses.NewWindow(height, width, th, tw)
	alrth.Box(0,0)

	alrtt := curses.NewWindow(3,width,th,tw)
	alrtt.Border(0,0,0,0,0,0,
				int(curses.ACS_VLINE()),
				int(curses.ACS_VLINE()))
	alrtt.Mvaddstr(1,int((width-len(title))/2),title)

	alrt := curses.NewWindow(height-5,width-2,th+3,tw+1)
	alrt.Addstr(text)

	alrth.Mvaddstr(height-1,0,"")

	alrth.Refresh()
	alrt.Refresh()
	alrtt.Refresh()
	
	char := byte(0)

	for char!=10 && char!=13 {
		char = byte(alrt.Getch())
		if char==10 || char==13 {
			alrth.Mvaddstr(0,0,strings.Repeat(" ",width*height))
			alrth.Refresh()
			curses.Stdscr.Refresh()
		}
	}
}
