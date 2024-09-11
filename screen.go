package main

import (
	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	tabBuffer *TabBuffer
	tScreen   tcell.Screen
}

func NewScreen(argv []string) (*Screen, error) {
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

	screen := &Screen{tScreen: s}
	var filename string

	if len(argv) > 1 {
		filename = argv[1]
	} else {
		filename = ""
	}

	width, height := s.Size()

	debugAreaHeght := 5
	fileInfoAreaHeight := 1

	screen.CreateDebugArea(debugAreaHeght)
	screen.CreateFileInfoArea(fileInfoAreaHeight)

	bounds := [4][2]int{
		{0, fileInfoAreaHeight}, {width, fileInfoAreaHeight}, {0, height - debugAreaHeght}, {width, height - debugAreaHeght},
	}

	screen.tabBuffer = NewTabBuffer("", 10, screen, filename, bounds)


	return screen, nil

}

//creates the debug area at the bottom of the screen
func (s *Screen) CreateDebugArea(height int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, sHeight := s.tScreen.Size()

	for y := sHeight - height; y < sHeight; y++ {
		for x := 0; x < sWidth; x++ {
			s.tScreen.SetContent(x, y, ' ', nil, style)
		}
	}
	s.tScreen.Show()
}

func (s *Screen) WriteDebug(str string, y int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	_, sHeight := s.tScreen.Size()

	startingPoint := 5

	for i, char := range str {
		s.tScreen.SetContent(i, (sHeight-startingPoint)+y, char, nil, style)
	}

	s.tScreen.Show()
}

func (s *Screen) CreateFileInfoArea(height int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, _ := s.tScreen.Size()

	for y := 0; y < height; y++ {
		for x := 0; x < sWidth; x++ {
			s.tScreen.SetContent(x, y, ' ', nil, style)
		}
	}
	s.tScreen.Show()
}
