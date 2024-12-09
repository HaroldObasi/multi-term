package buffer

import (
	"fmt"

	"github.com/HaroldObasi/multi-term/cursor"
)

type LineBuffer struct {
	Buffer []byte
	GapStart int
	GapEnd int
	GapSize int
	Cursor *cursor.Cursor
}

// String representation of LineBuffer
func (lb *LineBuffer) String() string {
	// var sb strings.Builder
	// for i := 0; i < len(lb.Buffer); i++ {
	// 	if i >= lb.GapStart && i < lb.GapEnd {
	// 		continue
	// 	}
	// 	sb.WriteByte(lb.Buffer[i])
	// }
	// return sb.String()

	return fmt.Sprintf(`
		GapStart: %v,
		GapEnd: %v,
		GapSize: %v,
		String: %v
	`, lb.GapStart, lb.GapEnd, lb.GapSize, lb.GetString())
}

// Method that returns the buffer as a string without the gap
func (lb *LineBuffer) GetString() string {
	return string(lb.Buffer[:lb.GapStart]) + string(lb.Buffer[lb.GapEnd:])
}

func NewLineBuffer(cursor *cursor.Cursor) *LineBuffer {
	GapSize := 10
	Buffer := make([]byte, GapSize)
	return &LineBuffer{
		Buffer: Buffer,
		GapSize: GapSize,
		GapStart: 0,
		GapEnd: GapSize,
		Cursor: cursor,
	}
}

func (lb *LineBuffer) GoTo(pos int){
	if pos < 0 || pos >= len(lb.Buffer) {
		return
	}

	if pos < lb.GapStart {
		// loop till GapStart
		diff := lb.GapStart - pos
		for i := 0; i < diff; i++ {
			lb.GoLeft()
		}
	} else if pos > lb.GapStart {
		// loop till GapEnd
		diff := pos - lb.GapStart
		for i := 0; i < diff; i++ {
			lb.GoRight()
		}
	}

}

func (lb *LineBuffer) GoLeft() {
	if lb.GapStart <= 0 {
		return
	}

	lb.Buffer[lb.GapEnd] = lb.Buffer[lb.GapStart-1]
	lb.GapEnd--

	lb.Buffer[lb.GapStart-1] = 0 // Clear the original position
	lb.GapStart--
}

func (lb *LineBuffer) GoRight() {
	if lb.GapEnd >= len(lb.Buffer)-1 {
		return
	}

	lb.Buffer[lb.GapStart] = lb.Buffer[lb.GapEnd+1]
	lb.GapStart++

	lb.Buffer[lb.GapEnd+1] = 0 // Clear the original position
	lb.GapEnd++
}

func (lb *LineBuffer) InsertRune(r rune) {
	pos := lb.Cursor.X
	if lb.GapStart >= len(lb.Buffer) {
		return
	}

	lb.GoTo(pos)

	lb.Buffer[lb.GapStart] = byte(r)
	lb.GapStart++
	lb.Cursor.GoTo(lb.Cursor.X+1, lb.Cursor.Y)
}