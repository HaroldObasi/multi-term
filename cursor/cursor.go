package cursor

import "github.com/gdamore/tcell/v2"

type Cursor struct {
	X, Y int
	screen tcell.Screen
}

func NewCursor(screen tcell.Screen) *Cursor {
	X, Y := 0, 0
	screen.ShowCursor(X, Y)
	return &Cursor{
		screen: screen,
		X: X,
		Y: Y,
	}
}