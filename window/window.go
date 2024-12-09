package window

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	screen tcell.Screen
	screenStyle tcell.Style
}

func NewWindow() *Window {
	return &Window{}
}

func (win *Window) Start() error {
	screen, err := tcell.NewScreen()

	if err != nil {
		return err
	}
	
	if err := screen.Init(); err != nil {
		return err
	}

	win.screen = screen

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	win.screenStyle = defStyle
	win.screen.SetStyle(win.screenStyle)

	quit := func() {
		win.screen.Fini()
		os.Exit(0)
	}

	for {
		// Update screen
		win.screen.Show()
	
		// Poll event
		ev := win.screen.PollEvent()
	
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			win.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}