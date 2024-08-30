package main

type TabBuffer struct {
	lines []LineBuffer
	gapStart int
	gapEnd int
	gapSize int
}


func NewTabBuffer(s string, gapSize int) *TabBuffer {
	lines := make([]LineBuffer, 0)
	lines = append(lines, *NewGapBuffer(s, 10))

	gapStart := 1
	gapEnd := gapStart + gapSize

	return &TabBuffer{
		gapStart: gapStart,
		gapSize: gapSize,
		gapEnd: gapEnd,
		lines: lines,
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

func (tb *TabBuffer) GoLeft(){
	if tb.gapStart <= 0 {
		return
	}

	tb.lines[tb.gapEnd] = tb.lines[tb.gapStart -1]
	tb.gapEnd--

	tb.lines[tb.gapStart] = LineBuffer{}
	tb.gapStart--
}

func (tb *TabBuffer) GoRight(){
	if tb.gapEnd >= len(tb.lines) {
		return
	}

	tb.lines[tb.gapStart] = tb.lines[tb.gapEnd + 1]
	tb.gapStart++

	tb.lines[tb.gapEnd+1] = LineBuffer{}
	tb.gapEnd++
}

func (tb *TabBuffer) AddLine(s string) {
	if tb.GetGapSize() <= 1 {
		tb._Grow()
	}

	line := NewGapBuffer(s, 10)
	tb.lines[tb.gapStart] = *line
	tb.gapStart++
}
