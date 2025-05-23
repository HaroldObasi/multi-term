package buffer

import (
	"fmt"
	"sync"

	"github.com/HaroldObasi/multi-term/cursor"
)

type LineBuffer struct {
	mu       sync.Mutex
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

func (lb *LineBuffer) GetContentLength() int {
	gapSize := lb.GetGapSize()
	bufferLength := len(lb.Buffer)

	return bufferLength - gapSize
}

func (lb *LineBuffer) GetBufferWithoutGap() []byte {
	// If GapEnd is the inclusive end of the gap, text after gap starts at GapEnd+1
	if lb.GapEnd+1 >= len(lb.Buffer) { // Check if there's any text after the gap
		return lb.Buffer[:lb.GapStart]
	}
	var tmp []byte
	tmp = append(tmp, lb.Buffer[:lb.GapStart]...)
	tmp = append(tmp, lb.Buffer[lb.GapEnd+1:]...)
	return tmp
}

func NewLineBuffer(cursor *cursor.Cursor, buffer []byte) *LineBuffer {
	GapSize := 10
	gapStart := len(buffer)

	emptyBuffer := make([]byte, GapSize)
	buffer = append(buffer, emptyBuffer...)

	gapEnd := len(buffer) - 1

	return &LineBuffer{
		Buffer:   buffer,
		GapSize:  GapSize,
		GapStart: gapStart,
		GapEnd:   gapEnd,
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

	lb.Buffer[lb.GapEnd] = lb.Buffer[lb.GapStart-1]
	lb.GapEnd--
	lb.Buffer[lb.GapStart-1] = 0
	lb.GapStart--
}

func (lb *LineBuffer) GoRight() {
	if lb.GapEnd >= len(lb.Buffer)-1 {
		return
	}

	lb.Buffer[lb.GapStart] = lb.Buffer[lb.GapEnd+1]
	lb.GapStart++
	lb.Buffer[lb.GapEnd+1] = 0
	lb.GapEnd++
}

func (lb *LineBuffer) InsertRune(r rune) {
	if lb.GetGapSize() <= 1 { // grow the buffer when there's only one space left
		lb.GrowBuffer()
	}

	if lb.GapStart >= len(lb.Buffer) {
		fmt.Print("\033[8;0H")
		fmt.Println("Buffer is full")
		return
	}

	// lb.GoTo(pos)

	lb.Buffer[lb.GapStart] = byte(r)
	lb.GapStart++

	lb.Cursor.GoTo(lb.Cursor.X+1, lb.Cursor.Y)
}

func (lb *LineBuffer) BackDelete() {
	// [97 0 0 0 0 0 0 0 98 99] == [0 0 0 0 0 0 0 0 98 99]
	// [0 0 0 0 0 0 0 97 98 99] == [0 0 0 0 0 0 0 97 98 99]

	if lb.GapStart <= 0 {
		return
	}

	lb.Buffer[lb.GapStart-1] = 0
	lb.GapStart -= 1
	lb.Cursor.GoTo(lb.Cursor.X-1, lb.Cursor.Y)
}

func (lb *LineBuffer) DeleteFromX() {

	arrSize := len(lb.Buffer) - lb.GapStart
	emptyBytes := make([]byte, arrSize)

	copy(lb.Buffer[lb.GapStart:], emptyBytes)

	lb.GapEnd = len(lb.Buffer) - 1
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
	return (lb.GapEnd - lb.GapStart) + 1
}
