package main

import (
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

	newLine := NewGapBuffer(s, 10, screen, cursor)
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

// adds the bytes from the read file to the buffers, without prn
func (tb *TabBuffer) InsertBytes(bytes []byte) {
	// split the file into different arrays. use the \n as a seperator
	// arr := utils.SplitRuneArray(bytes, 10)
}

func (tb *TabBuffer) WriteFileToScreen() {
	dat := tb.file.ReadFile()

	y := 1
	x := 0

	for _, c := range dat {
		if c == '\n' {
			// line := tb.GetLine(tb.cursor.y)
			// line.Insert(rune(c))
			tb.screen.tScreen.SetContent(x, y, rune(c), nil, tcell.StyleDefault)
			// tb.AddLine("")
			y++
			x = 0

		} else {
			tb.screen.tScreen.SetContent(x, y, rune(c), nil, tcell.StyleDefault)
			x++

			// line := tb.GetLine(tb.cursor.y)
			// line.Insert(rune(c))
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
	second := tb.lines[tb.gapEnd:]

	return append(first, second...)
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

func (tb *TabBuffer) AddLine(s string) {
	// tb.screen.WriteDebug("Adding line", 1)
	if tb.GetGapSize() <= 1 {
		tb.Grow()
	}

	line := NewGapBuffer(s, 10, tb.screen, tb.cursor)
	tb.lines[tb.gapStart] = line

	tb.cursor.SetPos(0, tb.cursor.y+1, tb)
	tb.gapStart++
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
