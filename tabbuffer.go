package main

type TabBuffer struct {
	lines    []LineBuffer
	gapStart int
	gapEnd   int
	gapSize  int
	cursor   *Cursor
	screen   *Screen
}

func NewTabBuffer(s string, gapSize int, screen *Screen) *TabBuffer {
	lines := make([]LineBuffer, 0)
	cursor := &Cursor{0, 0, screen}

	lines = append(lines, *NewGapBuffer(s, 10, screen, cursor))

	gapStart := 1
	gapEnd := gapStart + gapSize

	return &TabBuffer{
		gapStart: gapStart,
		gapSize:  gapSize,
		gapEnd:   gapEnd,
		lines:    lines,
	}
}

func (tb *TabBuffer) GetGapSize() int {
	return tb.gapEnd - tb.gapStart
}

func (tb *TabBuffer) _Grow() {
	newCapacity := tb.gapSize * 2
	newLines := make([]LineBuffer, newCapacity+len(tb.lines))

	copy(newLines, tb.lines[:tb.gapStart])

	newGapEnd := tb.gapStart + newCapacity

	copy(newLines[newGapEnd:], tb.lines[tb.gapEnd:])

	tb.lines = newLines
	tb.gapEnd = newGapEnd
}

func (tb *TabBuffer) GoLeft() {
	if tb.gapStart <= 0 {
		return
	}

	tb.lines[tb.gapEnd] = tb.lines[tb.gapStart-1]
	tb.gapEnd--

	tb.lines[tb.gapStart] = LineBuffer{}
	tb.gapStart--
}

func (tb *TabBuffer) GoRight() {
	if tb.gapEnd >= len(tb.lines) {
		return
	}

	tb.lines[tb.gapStart] = tb.lines[tb.gapEnd+1]
	tb.gapStart++

	tb.lines[tb.gapEnd+1] = LineBuffer{}
	tb.gapEnd++
}

func (tb *TabBuffer) AddLine(s string) {
	if tb.GetGapSize() <= 1 {
		tb._Grow()
	}

	line := NewGapBuffer(s, 10, tb.screen, tb.cursor)
	tb.lines[tb.gapStart] = *line
	tb.gapStart++
}

func (tb *TabBuffer) Write(char rune) {
	selectedLine := tb.lines[tb.cursor.y]
	selectedLine.GoTo(tb.cursor.x)
	selectedLine.Insert(char)
}
