package main

import "fmt"

type Cursor struct {
	x int
	y int

	screen *Screen

}

//use this method to set the position of the cursor
//dont directly change the x and y values
func (c *Cursor) SetPos(x, y int) {
	sWidth, sHeight := c.screen.tScreen.Size()

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

	c.screen.tScreen.ShowCursor(c.x, c.y)

	primary, _, _, width := c.screen.tScreen.GetContent(c.x, c.y)
	c.screen.WriteDebug(fmt.Sprintf("Item: %v, Width: %v", primary, width))

	// i want to use writeDebug method from window here

}
