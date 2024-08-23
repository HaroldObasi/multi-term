package main

import "github.com/gdamore/tcell/v2"

type Window struct { 
	x int
	y int
}

func (w *Window) Write(char rune, s tcell.Screen) {
	s.SetContent(w.x, w.y, char, nil, tcell.StyleDefault)
	w.x++
	s.Show()
}

func (w *Window) CreateDebugArea (s tcell.Screen){
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, sHeight := s.Size()

	for y := sHeight - 5; y < sHeight; y++ {
		for x := 0; x < sWidth; x++ {
			s.SetContent(x, y, ' ', nil, style)
		}
	}
	s.Show()
}

func (w *Window) WriteDebug(s tcell.Screen, str string){
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	_, sHeight := s.Size()

	startingPoint := 5

	for i, char := range str {
		s.SetContent(i, sHeight - startingPoint, char, nil, style)
	}

	s.Show()
}