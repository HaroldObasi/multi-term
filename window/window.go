package window

import (
	"fmt"
	"os"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	screen tcell.Screen
	screenStyle tcell.Style
	tab *buffer.TabBuffer
}

func NewWindow() (*Window, error) {
	win := &Window{}
	screen, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}
	
	if err := screen.Init(); err != nil {
		return nil, err
	}

	win.screen = screen // Assign screen to window

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	win.screenStyle = defStyle // Assign default style to window
	win.screen.SetStyle(win.screenStyle)

	tb := buffer.NewTabBuffer(win.screen)
	win.tab = tb // Assign tab buffer to window


	return win, nil
}

func (win *Window) Start() error {
	for {
		// Update screen
		win.screen.Show()
	
		// Poll and Handle events
		win.HandleEvents()
	}

}

func (win *Window) HandleEvents() {
	ev := win.screen.PollEvent()

	// Process event
	switch ev := ev.(type) {
	case *tcell.EventResize:
		win.screen.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			win.Quit()
		}
		switch ev.Key(){
		case tcell.KeyRune:
			fmt.Print("rune")
		}
	}
}

func (win *Window) Quit() {
	win.screen.Fini()
	os.Exit(0)
}