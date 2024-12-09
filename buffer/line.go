package buffer

type LineBuffer struct {
	buffer []byte
	gapStart int
	gapEnd int
	gapSize int
}

func NewLineBuffer() *LineBuffer {
	gapSize := 10
	buffer := make([]byte, gapSize)
	return &LineBuffer{
		buffer: buffer,
		gapSize: gapSize,
		gapStart: 0,
		gapEnd: gapSize,
	}
}

func (lb *LineBuffer) GoTo(pos int){
	if pos < 0 || pos >= len(lb.buffer) {
		return
	}

	if pos < lb.gapStart {
		// loop till gapstart
		diff := lb.gapStart - pos
		for i := 0; i < diff; i++ {
			lb.GoLeft()
		}
	} else if pos > lb.gapStart {
		// loop till gapend
		diff := pos - lb.gapStart
		for i := 0; i < diff; i++ {
			lb.GoRight()
		}
	}

}

func (lb *LineBuffer) GoLeft() {
	if lb.gapStart <= 0 {
		return
	}

	lb.buffer[lb.gapEnd] = lb.buffer[lb.gapStart-1]
	lb.gapEnd--

	lb.buffer[lb.gapStart-1] = 0 // Clear the original position
	lb.gapStart--
}

func (lb *LineBuffer) GoRight() {
	if lb.gapEnd >= len(lb.buffer)-1 {
		return
	}

	lb.buffer[lb.gapStart] = lb.buffer[lb.gapEnd+1]
	lb.gapStart++

	lb.buffer[lb.gapEnd+1] = 0 // Clear the original position
	lb.gapEnd++
}