package types;

import (
	"github.com/NDRAEY/curses"
)

type Screen struct {
	Width, Height int
}

const (
	Keydown  int = 66
	Keyup    int = 65
	Keyenter int = 10
)

var Inverse curses.ColorPair
