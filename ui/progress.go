package widgets;

import (
	"github.com/NDRAEY/curses"
	. "github.com/NDRAEY/Bulucomon/types"
	"strconv"
	"strings"
)

type Progress struct {
	width, height int
	title         string
	prc, pbp, remaining, fill int
}

func ProgressBar(screen Screen, title string, min int, max int, curr int) Progress {
	width  := screen.Width-6
	height := 8
	tw, th := (screen.Width-width)/2, (screen.Height-height)/2

	titlewnd := curses.NewWindow(3, width, th, tw)
	titlewnd.Border(0,0,0,0,0,0,
				int(curses.ACS_VLINE()),
				int(curses.ACS_VLINE()))

	pgb := curses.NewWindow(height-3, width, th+2, tw)
	pgb.Box(0, 0)

	titlewnd.Mvaddstr(1, 1, title)
	titlewnd.Refresh()

	prc := int((float32(curr)/float32(max-min))*100)
	pgb.Mvaddstr(2, 2, strconv.Itoa(prc)+"%")

	pbp := 2+len(strconv.Itoa(prc)+"%")+1
	remaining := width-pbp-2
	fill := int(float32(remaining)*(float32(curr)/float32(100)))
	
	pgb.Mvaddstr(2, pbp, strings.Repeat("=", fill))

	pgb.Refresh()

	return Progress{
		width: width,
		height: height,
		title: title,
		prc: prc,
		pbp: pbp,
		remaining: remaining,
		fill: fill}
}
