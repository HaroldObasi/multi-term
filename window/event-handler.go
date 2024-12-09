package window

import "os"

var Quit = func(win *Window) {
	win.Screen.Fini()
	os.Exit(0)
}
