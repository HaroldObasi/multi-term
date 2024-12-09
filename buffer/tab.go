package buffer

import (
	"github.com/HaroldObasi/multi-term/cursor"
	"github.com/gdamore/tcell/v2"
)

type TabBuffer struct {
	gapStart int
	gapEnd int
	gapSize int
	lines []*LineBuffer
	cursor *cursor.Cursor
}

func NewTabBuffer(screen tcell.Screen) *TabBuffer {
	gapSize := 10
	lines := make([]*LineBuffer, gapSize)
	cursor := cursor.NewCursor(screen)

	return &TabBuffer{
		gapSize: gapSize,
		lines: lines,
		gapStart: 0,
		gapEnd: gapSize,
		cursor: cursor,
	}
}


func (tb *TabBuffer) GoTo(i int) {

}

func (tb *TabBuffer) GoLeft() {

}

func (tb *TabBuffer) GoRight() {

}

