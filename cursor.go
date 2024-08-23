package main

import "github.com/gdamore/tcell/v2"

type Cursor struct {
	x int
	y int
}

func (c *Cursor) SetPos(x, y int, s tcell.Screen) {
	sWidth, sHeight := s.Size()

	if x < 0 {
		x = 0
	} else if x >= sWidth {
		x = sWidth - 1
	}

	if y < 0 {
		y = 0
	} else if y >= sHeight {
		y = sHeight - 1
	}

	c.x = x
	c.y = y

	s.ShowCursor(c.x, c.y)
	s.Show()
}
