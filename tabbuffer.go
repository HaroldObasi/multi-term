package main

type TabBuffer struct {
	lines    []*LineBuffer
	gapStart int
	gapEnd   int
	gapSize  int
	cursor   *Cursor
	screen   *Screen
	file     *File
}

func NewTabBuffer(s string, gapSize int, screen *Screen) *TabBuffer {
	lines := make([]*LineBuffer, gapSize)
	cursor := &Cursor{0, 0, screen}
	file := NewFile("test.txt")

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
	}

	dat := file.ReadFile()

	for _, c := range dat {
		if c == '\n' {
			tb.lines[tb.cursor.y].Insert(rune(c))
			tb.AddLine("")

		} else {
			tb.lines[tb.cursor.y].Insert(rune(c))
		}
	}

	screen.tScreen.Show()

	return &tb
}

// func (tb *TabBuffer) String() string {
// 	return "TabBuffer"
// }

func (tb *TabBuffer) GetGapSize() int {
	return (tb.gapEnd - tb.gapStart) + 1
}

func (tb *TabBuffer) _Grow() {
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
		tb.screen.WriteDebug("Growing", 1)
		// tb._Grow()
	}

	line := NewGapBuffer(s, 10, tb.screen, tb.cursor)
	tb.lines[tb.gapStart] = line

	tb.cursor.SetPos(0, tb.cursor.y+1)
	tb.gapStart++
}

func (tb *TabBuffer) Write(char rune) {
	selectedLine := tb.lines[tb.cursor.y]
	selectedLine.GoTo(tb.cursor.x)
	selectedLine.Insert(char)
}

// gets length of lines exluding the gap
func (tb *TabBuffer) Len() int {
	return len(tb.lines) - tb.GetGapSize()
}
