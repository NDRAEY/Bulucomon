package widgets;

import (
	"github.com/NDRAEY/curses"
	. "github.com/NDRAEY/Bulucomon/types"
	"strings"
)

func InputBox(screen Screen, text string) string {
	boxw, boxh := 27, 6
	tw, th := (screen.Width-boxw)/2, (screen.Height-boxh)/2

	inp := curses.NewWindow(boxh, boxw, th, tw)
	inp.Box(0,0)
	inp.Mvaddstr(1,2,text)
	inp.Mvaddch(3,1,int(' '))
	inp.Keypad(true)

	char := byte(0)
	buffer := []byte{}
	pos := 1

	for char!=10 && char!=13 {
		char = byte(inp.Getch())
		if char==127 {
			if pos>1 {
				inp.Mvaddch(3,pos,' ')
				inp.Wmove(3,pos)
				pos--
				buffer = buffer[:pos-1]
				inp.Refresh()
			}
			continue
		}
		inp.Addch(int(char))
		curses.Stdscr.Refresh()
		pos++
		if char==10 || char==13 {continue}
		buffer = append(buffer, char)
	}
	inp.Mvaddstr(0,0,strings.Repeat(" ",boxh*boxw))
	inp.Refresh()
	return string(buffer[:])
}
