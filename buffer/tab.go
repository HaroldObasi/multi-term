package buffer

import (
	"fmt"

	"github.com/HaroldObasi/multi-term/cursor"
	"github.com/gdamore/tcell/v2"
)

type TabBuffer struct {
	GapStart int
	GapEnd   int
	GapSize  int
	Lines    []*LineBuffer
	Cursor   *cursor.Cursor
}

func NewTabBuffer(screen tcell.Screen) *TabBuffer {
	gapSize := 10
	lines := make([]*LineBuffer, gapSize)
	cursor := cursor.NewCursor(screen)
	lines[0] = NewLineBuffer(cursor, []byte{})

	return &TabBuffer{
		GapSize:  gapSize,
		Lines:    lines,
		GapStart: 1,
		GapEnd:   gapSize - 1,
		Cursor:   cursor,
	}
}

func (tb *TabBuffer) GetGapSize() int {
	return (tb.GapEnd - tb.GapStart) + 1
}

func (tb *TabBuffer) GetContentLength() int {
	gapSize := tb.GetGapSize()
	linesLength := len(tb.Lines)

	fmt.Print("\033[27;0H\033[K")
	fmt.Printf("gap start: %v, gap end: %v", tb.GapStart, tb.GapEnd)
	return linesLength - gapSize
}

func (tb *TabBuffer) GetLinesWithoutGap() []*LineBuffer {
	if tb.GapEnd+1 >= len(tb.Lines) {
		return tb.Lines[:tb.GapStart]
	}

	var tmp []*LineBuffer

	tmp = append(tmp, tb.Lines[:tb.GapStart]...)
	tmp = append(tmp, tb.Lines[tb.GapEnd+1:]...)

	return tmp
}

func (tb *TabBuffer) InsertRune(r rune) {
	_, posY := tb.Cursor.GetPos()

	line := tb.Lines[posY]

	line.InsertRune(r)
}

func (tb *TabBuffer) BackDelete() {
	_, posY := tb.Cursor.GetPos()

	line := tb.Lines[posY]

	line.BackDelete()
}

func (tb *TabBuffer) InsertString(s string) {
	// iterate through the string and insert each rune
	for _, r := range s {
		tb.InsertRune(r)
	}
}

func (tb *TabBuffer) EnterLine() {
	// first get characters after cursor
	posX, posY := tb.Cursor.GetPos()
	linesWithoutGap := tb.GetLinesWithoutGap()

	currentLine := linesWithoutGap[posY]
	bufferWithoutGap := currentLine.GetBufferWithoutGap()

	bytesAfterCursor := bufferWithoutGap[posX:]
	// create new line buffer with characters
	newLine := NewLineBuffer(tb.Cursor, bytesAfterCursor)
	tb.InsertLine(newLine)

	// if gap is empty grow buffer
}

func (tb *TabBuffer) InsertLine(lb *LineBuffer) {
	if tb.GetGapSize() <= 1 {
		// grow tb
	}

	tb.Lines[tb.GapStart] = lb
	tb.GapStart++

	tb.Cursor.GoTo(0, tb.Cursor.Y+1)
	lb.GoTo(0)
}

func (tb *TabBuffer) GoTo(i int) {

}

func (tb *TabBuffer) GoLeft() {
	X, Y := tb.Cursor.GetPos()
	if X < 0 || Y < 0 {
		return
	}
	tb.Cursor.GoTo(
		X-1,
		Y,
	)
	currentLine := tb.Lines[Y]
	currentLine.GoLeft()
}

func (tb *TabBuffer) GoRight() {
	X, Y := tb.Cursor.GetPos()
	currentLine := tb.Lines[Y]

	if X >= currentLine.GetContentLength() {
		return
	}

	tb.Cursor.GoTo(
		X+1,
		Y,
	)
	currentLine.GoRight()
}

func (tb *TabBuffer) CursorUp() {

}
