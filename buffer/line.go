package buffer

import (
	"fmt"

	"github.com/HaroldObasi/multi-term/cursor"
)

type LineBuffer struct {
	Buffer   []byte
	GapStart int
	GapEnd   int
	GapSize  int // The Gapsize here represents how much we should grow the buffer by, not the current gapsize of the buffer
	Cursor   *cursor.Cursor
}

// String representation of LineBuffer
func (lb *LineBuffer) String() string {
	return fmt.Sprintf(`
		GapStart: %v,
		GapEnd: %v,
		GapSize: %v,
		String: %v,
		Buffer: %v
	`, lb.GapStart, lb.GapEnd, lb.GapSize, lb.GetString(), lb.Buffer)
}

func (lb *LineBuffer) GetString() string {
	// If GapEnd is the inclusive end of the gap, text after gap starts at GapEnd+1
	if lb.GapEnd+1 >= len(lb.Buffer) { // Check if there's any text after the gap
		return string(lb.Buffer[:lb.GapStart])
	}
	return string(lb.Buffer[:lb.GapStart]) + string(lb.Buffer[lb.GapEnd+1:])
}

// ...existing code...

func (lb *LineBuffer) GetBufferWithoutGap() []byte {
	// If GapEnd is the inclusive end of the gap, text after gap starts at GapEnd+1
	if lb.GapEnd+1 >= len(lb.Buffer) { // Check if there's any text after the gap
		return lb.Buffer[:lb.GapStart]
	}
	return append(lb.Buffer[:lb.GapStart], lb.Buffer[lb.GapEnd+1:]...)
}

// Method that returns the buffer as a string without the gap
// func (lb *LineBuffer) GetString() string {
// 	return string(lb.Buffer[:lb.GapStart]) + string(lb.Buffer[lb.GapEnd:])
// }

// func (lb *LineBuffer) GetBufferWithoutGap() []byte {
// 	return append(lb.Buffer[:lb.GapStart], lb.Buffer[lb.GapEnd:]...)
// }

func NewLineBuffer(cursor *cursor.Cursor) *LineBuffer {
	GapSize := 10
	Buffer := make([]byte, GapSize)
	return &LineBuffer{
		Buffer:   Buffer,
		GapSize:  GapSize,
		GapStart: 0,
		GapEnd:   GapSize - 1,
		Cursor:   cursor,
	}
}

func (lb *LineBuffer) GoTo(pos int) {
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

	// [ 1 2 0 0]
	// [ 1 0 0 2]

	lb.Buffer[lb.GapEnd] = lb.Buffer[lb.GapStart-1]
	lb.GapEnd--
	lb.Buffer[lb.GapStart-1] = 0
	lb.GapStart--
}

func (lb *LineBuffer) GoRight() {
	if lb.GapEnd >= len(lb.Buffer)-1 {
		return
	}

	// [ 1 0 0 2 ]
	// [ 1 2 0 0 ]

	lb.Buffer[lb.GapStart] = lb.Buffer[lb.GapEnd+1]
	lb.GapStart++
	lb.Buffer[lb.GapEnd+1] = 0
	lb.GapEnd++
}

func (lb *LineBuffer) InsertRune(r rune) {

	if lb.GetGapSize() <= 0 {
		lb.GrowBuffer()
	}

	pos := lb.Cursor.X
	if lb.GapStart >= len(lb.Buffer) {
		return
	}

	lb.GoTo(pos)

	lb.Buffer[lb.GapStart] = byte(r)
	lb.GapStart++
	lb.Cursor.GoTo(lb.Cursor.X+1, lb.Cursor.Y)

	// fmt.Println()
	// fmt.Println(lb.Buffer)
}

func (lb *LineBuffer) GrowBuffer() {
	newBuffer := make([]byte, lb.GapSize+len(lb.Buffer))

	copy(newBuffer, lb.Buffer[:lb.GapStart])

	newGapEnd := lb.GapStart + lb.GapSize
	copy(newBuffer[newGapEnd:], lb.Buffer[lb.GapEnd:])

	lb.Buffer = newBuffer
	lb.GapEnd = newGapEnd
}

func (lb *LineBuffer) GetGapSize() int {
	return lb.GapEnd - lb.GapStart
}
