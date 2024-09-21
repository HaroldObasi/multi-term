package main

import "fmt"

type Cursor struct {
	x int
	y int

	screen *Screen
}

func (c *Cursor) String() string {
	return fmt.Sprintf("Cursor: (%v, %v)", c.x, c.y)
}

func NewCursor(x, y int, screen *Screen) *Cursor {
	cursor := &Cursor{x, y, screen}
	screen.tScreen.ShowCursor(cursor.x, cursor.y)
	return cursor
}

// use this method to set the position of the cursor
// dont directly change the x and y values
func (c *Cursor) SetPos(x, y int, tb *TabBuffer) {
	sWidth, _ := c.screen.tScreen.Size()

	upperBound := tb.bounds[0][1]
	lowerBound := tb.bounds[2][1]

	if x < 0 {
		x = 0
	} else if x > sWidth {
		x = sWidth
	}

	if y < upperBound {
		y = upperBound
	} else if y >= lowerBound {
		y = lowerBound - 1
	}

	c.x = x
	c.y = y

	c.screen.tScreen.ShowCursor(c.x, c.y)

	// primary, _, _, width := c.screen.tScreen.GetContent(c.x, c.y)
	// c.screen.WriteDebug(fmt.Sprintf("Item: %v, Width: %v", primary, width))

	c.screen.WriteDebug(fmt.Sprintf("Col: %d, Row: %d", c.x, c.y), 0)
	// c.screen.WriteDebug(fmt.Sprintf("buffer: %v ", c.screen.tabBuffer.lines[c.y].buffer), 1)
}
