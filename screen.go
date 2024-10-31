package main

import (
	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	tabBuffer *TabBuffer
	tScreen   tcell.Screen
	debugAreaHeight int
	fileInfoAreaHeight int
}

func NewScreen(argv []string) (*Screen, error) {
		
	const DEBUG_AREA_HEIGHT = 5
	const FILE_INFO_AREA_HEIGHT = 1


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

	screen := &Screen{tScreen: s}
	var filename string

	if len(argv) > 1 {
		filename = argv[1]
	} else {
		filename = ""
	}

	screen.debugAreaHeight = DEBUG_AREA_HEIGHT
	screen.fileInfoAreaHeight = FILE_INFO_AREA_HEIGHT
	
	screen.CreateDebugArea()
	screen.CreateFileInfoArea()
	screen.WriteFileName(filename, 0)
	
	width, height := s.Size()

	bounds := [4][2]int{
		{0, screen.fileInfoAreaHeight}, 
		{width, screen.fileInfoAreaHeight}, 
		{0, height - screen.debugAreaHeight}, 
		{width, height - screen.debugAreaHeight},
	}


	if filename == "" {
		screen.tabBuffer = NewTabBuffer("", 10, screen, filename, bounds)
	} else {
		screen.tabBuffer = NewTabBufferFromFile(filename, screen, bounds)
	}

	// screen.tabBuffer.cursor.SetPos(0, bounds[0][1], screen.tabBuffer)

	return screen, nil

}

// creates the debug area at the bottom of the screen
func (s *Screen) CreateDebugArea() {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, sHeight := s.tScreen.Size()

	for y := sHeight - s.debugAreaHeight ; y < sHeight; y++ {
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

func (s *Screen) CreateFileInfoArea() {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, _ := s.tScreen.Size()

	for y := 0; y < s.fileInfoAreaHeight; y++ {
		for x := 0; x < sWidth; x++ {
			s.tScreen.SetContent(x, y, ' ', nil, style)
		}
	}
	s.tScreen.Show()
}

func (s *Screen) WriteFileName(str string, row int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	sWidth, _ := s.tScreen.Size()

	lenStr := len(str)

	for i, char := range str {
		s.tScreen.SetContent(((sWidth/2)-lenStr/2)+i, row, char, nil, style)
	}

	s.tScreen.Show()
}
