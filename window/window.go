package window

import (
	"os"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen      tcell.Screen
	screenStyle tcell.Style
	Tab         *buffer.TabBuffer
	events      chan string
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

	win.events = make(chan string, 100)

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

func (win *Window) Render() {
	// width, height := win.Screen.Size()

	// text := win.Tab.Lines[0].GetString()
	buffer := win.Tab.Lines[0].GetBufferWithoutGap()

	x := 0

	// for i, ch := range text {
	// 	win.Screen.SetContent(i, 0, ch, nil, win.screenStyle)
	// }

	for _, b := range buffer {
		win.Screen.SetContent(x, 0, rune(b), nil, win.screenStyle)
		x++
	}

	win.Screen.Show()
}

func (win *Window) Start() error {
	go func() {
		for range win.events {
			win.Render()
		}
	}()

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
		switch ev.Key() {
		case tcell.KeyRune:
			ch := ev.Rune()
			win.Tab.InsertRune(ch)

			win.events <- "render"
		case tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight:
			win.HandleDirection(ev.Key())
		}
	}
}

func (win *Window) Quit() {
	win.Screen.Fini()
	os.Exit(0)
}
