package window

import (
	"os"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen tcell.Screen
	screenStyle tcell.Style
	Tab *buffer.TabBuffer
}

func NewWindow() (*Window, error) {
	win := &Window{}
	Screen, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}
	
	if err := Screen.Init(); err != nil {
		return nil, err
	}

	win.Screen = Screen // Assign Screen to window

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	win.screenStyle = defStyle // Assign default style to window
	win.Screen.SetStyle(win.screenStyle)

	tb := buffer.NewTabBuffer(win.Screen)
	win.Tab = tb // Assign tab buffer to window


	return win, nil
}

func NewTestWindow() (*Window, error) {
	win := &Window{}
	Screen := tcell.NewSimulationScreen("UTF-8")
	
	if err := Screen.Init(); err != nil {
		return nil, err
	}

	win.Screen = Screen // Assign Screen to window

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	win.screenStyle = defStyle // Assign default style to window
	win.Screen.SetStyle(win.screenStyle)

	tb := buffer.NewTabBuffer(win.Screen)
	win.Tab = tb // Assign tab buffer to window

	return win, nil
}

func (win *Window) Start() error {
	for {
		// Update Screen
		win.Screen.Show()
	
		// Poll and Handle events
		win.HandleEvents()
	}

}

func (win *Window) HandleEvents() {
	ev := win.Screen.PollEvent()

	// Process event
	switch ev := ev.(type) {
	case *tcell.EventResize:
		win.Screen.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			win.Quit()
		}
		switch ev.Key(){
		case tcell.KeyRune:
			ch := ev.Rune()
			win.Tab.InsertRune(ch)
		}
	}
}

func (win *Window) Quit() {
	win.Screen.Fini()
	os.Exit(0)
}