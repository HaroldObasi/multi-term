package main

import (
	"fmt"
	"strings"
)

type GapBuffer struct {
	gapSize  int
	gapStart int
	gapEnd   int
	buffer   []rune
}

func (g *GapBuffer) String() string {
	return fmt.Sprintf("Gap Start: %v, Gap End: %v, Buffer: %v", g.gapStart, g.gapEnd, g.GetText())
}

func NewGapBuffer(s string, gapSize int) *GapBuffer {

	runes := []rune(s)
	buffer := make([]rune, gapSize+len(runes))

	copy(buffer, runes)

	gapStart := len(runes)
	gapEnd := (gapStart + gapSize) - 1

	return &GapBuffer{
		gapStart: gapStart,
		gapSize:  gapSize,
		gapEnd:   gapEnd,
		buffer:   buffer,
	}
}

func (g *GapBuffer) GetGapSize() int {
	return (g.gapEnd - g.gapStart) + 1
}

func (g *GapBuffer) GetText() string {
	var sb strings.Builder
	sb.Grow(len(g.buffer) - g.GetGapSize())
	sb.WriteString(string(g.buffer[:g.gapStart]))
	sb.WriteString(string(g.buffer[g.gapEnd+1:]))
	return sb.String()
}

func (g *GapBuffer) Insert(r rune) {
	if g.GetGapSize() <= 1 {
		g.Grow()
	}
	g.buffer[g.gapStart] = r
	g.gapStart++
}

func (g *GapBuffer) Grow() {
	newCapacity := g.gapSize * 2
	newBuffer := make([]rune, newCapacity+len(g.buffer))


	copy(newBuffer, g.buffer[:g.gapStart])

	newGapEnd := g.gapStart + newCapacity
	copy(newBuffer[newGapEnd:], g.buffer[g.gapEnd:])

	g.buffer = newBuffer
	g.gapEnd = newGapEnd

}

func (g *GapBuffer) GoLeft() {
	if g.gapStart <= 0 {
		return
	}

	g.buffer[g.gapEnd] = g.buffer[g.gapStart-1]
	g.gapEnd--

	g.buffer[g.gapStart-1] = 0 // Clear the original position
	g.gapStart--
}

func (g *GapBuffer) GoRight() {
	if g.gapEnd >= len(g.buffer)-1 {
		return
	}

	g.buffer[g.gapStart] = g.buffer[g.gapEnd+1]
	g.gapStart++

	g.buffer[g.gapEnd+1] = 0 // Clear the original position
	g.gapEnd++
}

func (g *GapBuffer) ChangeCursorPos(index int) {

	// gap starts on cursor index

	if index < 0 || index >= len(g.buffer) {
		return // Invalid index
	}

	gapSize := g.GetGapSize()

	if index < g.gapStart {
		// Move gap left
		tmp := make([]rune, g.gapStart-index)
		copy(tmp, g.buffer[index:g.gapStart])
		copy(g.buffer[index:index+gapSize], g.buffer[g.gapStart:g.gapEnd+1]) //good

		g.gapStart = index               //good
		g.gapEnd = (index + gapSize) - 1 //good

		copy(g.buffer[g.gapEnd+1:(g.gapEnd+1)+len(tmp)], tmp) // correct index on lside,

	} else if index > g.gapStart {

	}

}