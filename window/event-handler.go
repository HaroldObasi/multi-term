package window

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

var Quit = func(win *Window) {
	win.Screen.Fini()
	os.Exit(0)
}

func (win *Window) HandleDirection(key tcell.Key) {
	switch key {
	// case tcell.KeyUp:
	// 	win.Tab.CursorUp()
	// case tcell.KeyDown:
	// 	win.Tab.CursorDown()
	case tcell.KeyLeft:
		win.Tab.GoLeft()
	case tcell.KeyRight:
		win.Tab.GoRight()
	}
}
