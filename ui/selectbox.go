package widget;

import (
	"github.com/NDRAEY/curses"
	. "github.com/NDRAEY/Bulucomon/types"
)

func SelectBox(screen Screen, title string, values []string) string {
	maxlen := 0
	for _, elm := range values {
		if len(elm)>maxlen {
			maxlen = len(elm)
		}
	}
	maxlen += 4

	width := 4+maxlen
	height:= 4+len(values)

	y := (screen.Height-height)/2
	x := (screen.Width-width)/2

	box := curses.NewWindow(height, width, y, x)
	box.Box(0,0)
	box.Mvaddstr(1, 1, title)

	for idx, elm := range values {
		box.Mvaddstr(idx+3, 2, elm)
	}

	box.Attron(Inverse.Attribute())
	box.Mvaddstr(3, 1, ">"+values[0]+strings.Repeat(" ", maxlen-len(values[0])))
	box.Attroff(Inverse.Attribute())
	box.Refresh()

	position := 0
	for {
		key := int(box.Getch())
		if key==Keydown {
			if position<len(values)-1 {position++}
			prevpos := position-1
			if prevpos<0 { prevpos=0 }
			box.Mvaddstr(3+prevpos, 1, " "+values[prevpos]+strings.Repeat(" ",maxlen-len(values[prevpos])))
			box.Mvaddstr(3+position, 1, ">"+values[position]+strings.Repeat(" ", maxlen-len(values[position])))

		}else if key==Keyup {
			if position>0 {position--}
			nxpos := position+1
			if nxpos>len(values)-1 { nxpos=len(values)-1 }
			box.Mvaddstr(3+nxpos, 1, " "+values[nxpos]+strings.Repeat(" ",maxlen-len(values[nxpos])))
			box.Mvaddstr(3+position, 1, ">"+values[position]+strings.Repeat(" ", maxlen-len(values[position])))
		}else if key==Keyenter {
			box.Mvaddstr(0,0,strings.Repeat(" ", width*height))
			box.Refresh()
			return values[position]
		}
		box.Refresh()
	}
}
