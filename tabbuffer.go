package main

import (
	"fmt"

	"github.com/HaroldObasi/multi-term/utils"
	"github.com/gdamore/tcell/v2"
)

type TabBuffer struct {
	lines    []*LineBuffer
	gapStart int
	gapEnd   int
	gapSize  int
	bounds   [4][2]int
	cursor   *Cursor
	screen   *Screen
	file     *File
}

func NewTabBuffer(s string, gapSize int, screen *Screen, filename string, bounds [4][2]int) *TabBuffer {
	lines := make([]*LineBuffer, gapSize)

	upperBound := bounds[0][1]

	cursor := NewCursor(0, upperBound, screen)

	if filename == "" {
		filename = "test.txt"
	}

	file := NewFile(filename)

	newLine := NewLineBuffer(s, 10, screen, cursor)
	lines[0] = newLine

	gapStart := 1
	gapEnd := (gapStart + gapSize) - 1

	tb := TabBuffer{
		gapStart: gapStart,
		gapSize:  gapSize,
		gapEnd:   gapEnd,
		lines:    lines,
		cursor:   cursor,
		screen:   screen,
		file:     file,
		bounds:   bounds,
	}
	return &tb
}

func NewTabBufferFromFile(filename string, screen *Screen, bounds [4][2]int) *TabBuffer {
	file := NewFile(filename)
	dat := file.ReadFile()
	arr := utils.SplitRuneArray([]rune(string(dat)), 10)

	var gapStart, gapEnd, gapSize int
	var lines []*LineBuffer

	upperBound := bounds[0][1]
	cursor := NewCursor(0, upperBound, screen)

	if len(arr) > 0 {
		gapStart = len(arr)
		gapSize = 10
		gapEnd = (gapStart + gapSize) - 1

		lines = make([]*LineBuffer, len(arr)+gapSize)
		for i := range arr {
			lines[i] = NewLineBuffer(string(arr[i]), 10, screen, cursor)
		}
	} else {
		gapStart = 1
		gapSize = 10
		gapEnd = gapSize - 1

		lines = make([]*LineBuffer, gapSize)
		lines[0] = NewLineBuffer("", 10, screen, cursor)
	}

	return &TabBuffer{
		lines:    lines,
		gapStart: gapStart,
		gapSize:  gapSize,
		gapEnd:   gapEnd,
		cursor:   cursor,
		screen:   screen,
		file:     file,
		bounds:   bounds,
	}
}

func (tb *TabBuffer) WriteFileToScreen() {
	dat := tb.file.ReadFile()

	upperBound := tb.bounds[0][1]
	x := 0

	for _, c := range dat {
		if c == '\n' {
			tb.screen.tScreen.SetContent(x, upperBound, rune(c), nil, tcell.StyleDefault)
			upperBound++
			x = 0

		} else if c == '\t' {
			x += 4
		} else {
			tb.screen.tScreen.SetContent(x, upperBound, rune(c), nil, tcell.StyleDefault)
			x++
		}
	}
	tb.screen.tScreen.Show()
}

// func (tb *TabBuffer) String() string {
// 	return "TabBuffer"
// }

func (tb *TabBuffer) GetGapSize() int {
	return (tb.gapEnd - tb.gapStart) + 1
}

func (tb *TabBuffer) GetUpperBound() int {
	return tb.bounds[0][1]
}

func (tb *TabBuffer) GetLowerBound() int {
	return tb.bounds[2][1]
}

func (tb *TabBuffer) GetLine(y int) *LineBuffer {
	// get the upper bound of the screen

	upperBound := tb.bounds[0][1]
	return tb.lines[y-upperBound]
}

func (tb *TabBuffer) Grow() {
	newCapacity := tb.gapSize * 2
	newLines := make([]*LineBuffer, newCapacity+len(tb.lines))

	copy(newLines, tb.lines[:tb.gapStart])

	newGapEnd := tb.gapStart + newCapacity

	copy(newLines[newGapEnd:], tb.lines[tb.gapEnd:])

	tb.lines = newLines
	tb.gapEnd = newGapEnd
}

func (tb *TabBuffer) GetValidLines() []*LineBuffer {
	first := tb.lines[:tb.gapStart]
	second := tb.lines[tb.gapEnd+1:]

	return append(first, second...)
}

func (tb *TabBuffer) GoTo(pos int) {
	if pos < 0 || pos >= len(tb.lines) {
		return
	}

	if pos < tb.gapStart {
		diff := tb.gapStart - pos

		for i := 0; i < diff; i++ {
			tb.GoLeft()
		}
	} else if pos > tb.gapStart {
		diff := pos - tb.gapStart

		for i := 0; i < diff; i++ {
			tb.GoRight()
		}
	}
}

func (tb *TabBuffer) GoLeft() {
	if tb.gapStart <= 0 {
		return
	}

	tb.lines[tb.gapEnd] = tb.lines[tb.gapStart-1]
	tb.gapEnd--

	tb.lines[tb.gapStart] = &LineBuffer{}
	tb.gapStart--
}

func (tb *TabBuffer) GoRight() {
	if tb.gapEnd >= len(tb.lines) {
		return
	}

	tb.lines[tb.gapStart] = tb.lines[tb.gapEnd+1]
	tb.gapStart++

	tb.lines[tb.gapEnd+1] = &LineBuffer{}
	tb.gapEnd++
}

func (tb *TabBuffer) AddLine(s string, y int, x int) {

	// tb.screen.WriteDebug("Adding line", 1)
	if tb.GetGapSize() <= 1 {
		tb.Grow()
	}

	tb.screen.WriteDebug(fmt.Sprintf("init buf: %v. gs: %v, ge: %v", tb.lines, tb.gapStart, tb.gapEnd), 1)

	//buffer handling
	// first go to position

}

func (tb *TabBuffer) ReDraw(y int) {
	posY := y + tb.GetUpperBound() //2

	validLines := tb.GetValidLines()
	linesToRedraw := validLines[y:] // account for the upper bound, change this whole uppoerbound logic later

	// tb.screen.WriteDebug(fmt.Sprintf("lines to redraw: %v", ), 2)

	for count := 0; count < len(linesToRedraw); count++ {
		linesToRedraw[count].ReDraw(0, posY)
		posY++
	}
}

func (tb *TabBuffer) Write(char rune) {
	selectedLine := tb.GetLine(tb.cursor.y)
	selectedLine.GoTo(tb.cursor.x)
	selectedLine.Insert(char)
}

// gets length of lines exluding the gap
func (tb *TabBuffer) Len() int {
	return len(tb.lines) - tb.GetGapSize()
}
