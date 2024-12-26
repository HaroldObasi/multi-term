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

func (tb *TabBuffer) GoTo(i int) {

}

func (tb *TabBuffer) GoLeft() {
	X, Y := tb.Cursor.GetPos()
	tb.Cursor.GoTo(
		X-1,
		Y,
	)

}

func (tb *TabBuffer) GoRight() {
	X, Y := tb.Cursor.GetPos()
	tb.Cursor.GoTo(
		X+1,
		Y,
	)
}

func (tb *TabBuffer) CursorUp() {

}
