package buffer

import (
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
	lines[0] = NewLineBuffer(cursor)

	return &TabBuffer{
		GapSize:  gapSize,
		Lines:    lines,
		GapStart: 0,
		GapEnd:   gapSize,
		Cursor:   cursor,
	}
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
