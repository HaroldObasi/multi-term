package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type LineBuffer struct {
	gapSize  int
	gapStart int
	gapEnd   int
	buffer   []rune
	screen   *Screen
	cursor   *Cursor
}

func (lb *LineBuffer) String() string {
	return fmt.Sprintf("Gs: %v, Ge: %v, len: %v", lb.gapStart, lb.gapEnd, lb.GetText())
}

func NewLineBuffer(s string, gapSize int, screen *Screen, cursor *Cursor) *LineBuffer {

	//TODO make sure this accepts an array of runes instead of a string

	runes := []rune(s)
	buffer := make([]rune, gapSize+len(runes))

	copy(buffer, runes)

	gapStart := len(runes)
	gapEnd := (gapStart + gapSize) - 1

	lb := &LineBuffer{
		gapStart: gapStart,
		gapSize:  gapSize,
		gapEnd:   gapEnd,
		buffer:   buffer,
		screen:   screen,
		cursor:   cursor,
	}

	// lb.Write(s)
	return lb
}

// redraws the line on the screen at the specified position
func (lb *LineBuffer) ReDraw(x, y int) {
	//buffer : [1, 2, 0, 0, 0] str: "12"
	//specified position is 1

	// from 1 to the width of the screen

	str := lb.GetText()
	width, _ := lb.screen.tScreen.Size()

	for ; x < width; x++ {
		if x >= len(str) {
			lb.screen.tScreen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		} else {
			lb.screen.tScreen.SetContent(x, y, rune(str[x]), nil, tcell.StyleDefault)
		}
	}
	lb.screen.tScreen.Show()
}

func (lb *LineBuffer) GetGapSize() int {
	return (lb.gapEnd - lb.gapStart) + 1
}

func (lb *LineBuffer) GetText() string {
	var sb strings.Builder
	sb.Grow(len(lb.buffer) - lb.GetGapSize())
	sb.WriteString(string(lb.buffer[:lb.gapStart]))
	sb.WriteString(string(lb.buffer[lb.gapEnd+1:]))
	return sb.String()
}

func (lb *LineBuffer)GetBufferWithoutGap() []rune{
	first := lb.buffer[:lb.gapStart]
	second := lb.buffer[lb.gapEnd+1:]

	return append(first, second...)
}

func (lb *LineBuffer) GetRunes() string {
	first := lb.buffer[:lb.gapStart]
	second := lb.buffer[lb.gapEnd+1:]

	return string(append(first, second...))
}

func (lb *LineBuffer) Write(s string) {
	for _, r := range s {
		lb.Insert(r)
	}
	// lb.screen.WriteDebug("Done writing string")

}
// adds a rune to the buffer, without changing the cursor
func (lb *LineBuffer) Add(r rune) {
	if lb.GetGapSize() <= 1 {
		lb.Grow()
	}
	lb.buffer[lb.gapStart] = r

	lb.gapStart++
}

func (lb *LineBuffer) AddString(s string) {
	for _, r := range s {
		lb.Add(rune(r))
	}
}

func (lb *LineBuffer) Insert(r rune) {
	if lb.GetGapSize() <= 1 {
		lb.Grow()
	}

	lb.GoTo(lb.cursor.x)
	lb.buffer[lb.gapStart] = r

	// x := lb.gapStart
	lb.cursor.SetPos(lb.cursor.x+1, lb.cursor.y, lb.screen.tabBuffer)

	lb.gapStart++

	lb.screen.tScreen.Show()
}


func (lb *LineBuffer) Delete() {
	cursor := lb.cursor
	cursorX, cursorY := cursor.GetCursorPos()

	if cursorX <= 0 && cursorY <=0 {
		return
	}

	if cursorX <= 0 {
		// delete the newline character
		lb.screen.tabBuffer.DeleteLine(cursorY)
		return
	}

	lb.GoTo(cursorX)

	lb.buffer[lb.gapStart-1] = 0
	lb.gapStart--

	lb.cursor.SetPos(lb.cursor.x-1, lb.cursor.y, lb.screen.tabBuffer)
}

func (lb *LineBuffer) Grow() {
	newCapacity := lb.gapSize * 2
	newBuffer := make([]rune, newCapacity+len(lb.buffer))

	copy(newBuffer, lb.buffer[:lb.gapStart])

	newGapEnd := lb.gapStart + newCapacity
	copy(newBuffer[newGapEnd:], lb.buffer[lb.gapEnd:])

	lb.buffer = newBuffer
	lb.gapEnd = newGapEnd

}

func (lb *LineBuffer) GoLeft() {
	if lb.gapStart <= 0 {
		return
	}

	lb.buffer[lb.gapEnd] = lb.buffer[lb.gapStart-1]
	lb.gapEnd--

	lb.buffer[lb.gapStart-1] = 0 // Clear the original position
	lb.gapStart--
}

func (lb *LineBuffer) GoRight() {
	if lb.gapEnd >= len(lb.buffer)-1 {
		return
	}

	lb.buffer[lb.gapStart] = lb.buffer[lb.gapEnd+1]
	lb.gapStart++

	lb.buffer[lb.gapEnd+1] = 0 // Clear the original position
	lb.gapEnd++
}

// Moves the gap to start at the specified position, so it is the same as the cursor position
func (lb *LineBuffer) GoTo(pos int) {
	// lb.screen.WriteDebug(fmt.Sprintf("Old gap start: %d, pos: %d", lb.gapStart, pos), 3)

	if pos < 0 || pos >= len(lb.buffer) {
		return
	}

	if pos < lb.gapStart {
		// loop till gapstart
		diff := lb.gapStart - pos

		for i := 0; i < diff; i++ {
			lb.GoLeft()
		}
	} else if pos > lb.gapStart {
		// loop till gapend
		diff := pos - lb.gapStart

		for i := 0; i < diff; i++ {
			lb.GoRight()
		}

	}
	// lb.screen.WriteDebug(fmt.Sprintf("new gap start: %d", lb.gapStart), 4)

}

func (lb *LineBuffer) GoToEnd() {
	index := len(lb.buffer) - lb.GetGapSize()
	lb.GoTo(index)
}

// returns the length of the buffer excluding the gap
func (lb *LineBuffer) Len() int {
	return len(lb.buffer) - lb.GetGapSize()
}
