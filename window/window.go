package window

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	tScreen tcell.Screen
	screenStyle tcell.Style
}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Start() error {
	tScreen, err := tcell.NewScreen()

	if err != nil {
		return err
	}
	
	if err := tScreen.Init(); err != nil {
		return err
	}

	s.tScreen = tScreen

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	s.screenStyle = defStyle
	s.tScreen.SetStyle(s.screenStyle)

	quit := func() {
		s.tScreen.Fini()
		os.Exit(0)
	}

	for {
		// Update screen
		s.tScreen.Show()
	
		// Poll event
		ev := s.tScreen.PollEvent()
	
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.tScreen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}