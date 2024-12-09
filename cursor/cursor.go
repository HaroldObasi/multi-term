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

func (cursor Cursor) GetPos() (int, int) {
	return cursor.X, cursor.Y
}

func (cursor *Cursor) GoTo(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	if x > cursor.X {
		for i := 0; i < x-cursor.X; i++ {
			cursor.goRight()
		}
	} else if x < cursor.X {
		for i := 0; i < cursor.X-x; i++ {
			cursor.goLeft()
		}
	} else if y > cursor.Y {
		for i := 0; i < y-cursor.Y; i++ {
			cursor.goDown()
		}
	} else if y < cursor.Y {
		for i := 0; i < cursor.Y-y; i++ {
			cursor.goUp()
		}
	}

	cursor.screen.ShowCursor(cursor.X, cursor.Y)
}

func (cursor *Cursor) goLeft() {
	cursor.X--
}

func (cursor *Cursor) goRight() {
	cursor.X++
}

func (cursor *Cursor) goUp() {
	cursor.Y--
}

func (cursor *Cursor) goDown() {
	cursor.Y++
}
