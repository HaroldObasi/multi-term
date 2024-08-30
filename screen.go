package main

import "github.com/gdamore/tcell/v2"

type Screen struct {
	tabBuffer *TabBuffer
	tScreen   tcell.Screen
}

func NewScreen() (*Screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.ShowCursor(0, 0)

	screen := &Screen{tScreen: s, tabBuffer: NewTabBuffer("", 10)}

	screen.CreateDebugArea()
	return screen, nil

}

func (s *Screen) CreateDebugArea() {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, sHeight := s.tScreen.Size()

	for y := sHeight - 5; y < sHeight; y++ {
		for x := 0; x < sWidth; x++ {
			s.tScreen.SetContent(x, y, ' ', nil, style)
		}
	}
	s.tScreen.Show()
}

func (s *Screen) WriteDebug(str string) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	_, sHeight := s.tScreen.Size()

	startingPoint := 5

	for i, char := range str {
		s.tScreen.SetContent(i, sHeight-startingPoint, char, nil, style)
	}

	s.tScreen.Show()
}
