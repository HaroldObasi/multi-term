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
	return fmt.Sprintf("Gap Start: %v, Gap End: %v, Buffer: %v", lb.gapStart, lb.gapEnd, lb.GetText())
}

func NewGapBuffer(s string, gapSize int, screen *Screen, cursor *Cursor) *LineBuffer {

	// screen.WriteDebug(fmt.Sprintf("Recieved cursr: %v", cursor.x))
	// time.Sleep(3 * time.Second)

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

	lb.Write(s)
	return lb
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

func (lb *LineBuffer) Write(s string) {
	for _, r := range s {
		lb.Insert(r)
	}
	// lb.screen.WriteDebug("Done writing string")

}

func (lb *LineBuffer) Insert(r rune) {
	if lb.GetGapSize() <= 1 {
		lb.Grow()
	}
	lb.buffer[lb.gapStart] = r

	lb.screen.tScreen.SetContent(lb.cursor.x, lb.cursor.y, r, nil, tcell.StyleDefault)
	lb.cursor.SetPos(lb.cursor.x+1, lb.cursor.y)

	lb.screen.tScreen.Show()

	lb.gapStart++

}

func (lb *LineBuffer) Delete() {
	if lb.gapStart <= 0 {
		return
	}

	lb.screen.tScreen.SetContent(lb.cursor.x-1, lb.cursor.y, 0, nil, tcell.StyleDefault)

	lb.gapStart--
	lb.buffer[lb.gapStart] = 0

	lb.cursor.SetPos(lb.gapStart, lb.cursor.y)

	lb.screen.tScreen.Show()
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

func (lb *LineBuffer) GoTo(pos int) {
	if pos < 0 || pos >= len(lb.buffer) {
		return
	}
	// [1, 2, 3, 4, 0, 0, 0]
	if pos < lb.gapStart {
		// loop till gapstart
		for i := pos; i < lb.gapStart; i++ {
			lb.GoLeft()
		}
	} else if pos > lb.gapStart {
		// loop till gapend
		for i := pos; i > lb.gapEnd; i-- {
			lb.GoRight()
		}
	}
}

// deprecate this
func (lb *LineBuffer) ChangeCursorPos(index int) {

	// gap starts on cursor index

	if index < 0 || index >= len(lb.buffer) {
		return // Invalid index
	}

	gapSize := lb.GetGapSize()

	if index < lb.gapStart {
		// Move gap left
		tmp := make([]rune, lb.gapStart-index)
		copy(tmp, lb.buffer[index:lb.gapStart])
		copy(lb.buffer[index:index+gapSize], lb.buffer[lb.gapStart:lb.gapEnd+1]) //good

		lb.gapStart = index               //good
		lb.gapEnd = (index + gapSize) - 1 //good

		copy(lb.buffer[lb.gapEnd+1:(lb.gapEnd+1)+len(tmp)], tmp) // correct index on lside,

	} else if index > lb.gapStart {

	}

}

// returns the length of the buffer excluding the gap
func (lb *LineBuffer) Len() int {
	return len(lb.buffer) - lb.GetGapSize()
}
