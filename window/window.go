package window

import (
	"fmt"
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
	width, height := win.Screen.Size()

	for y := range height {
		// fmt.Printf("y: %v, height: %v \n", y, height)

		if y >= win.Tab.GetContentLength() {
			break
		}

		linesWithoutGap := win.Tab.GetLinesWithoutGap()
		currentLine := linesWithoutGap[y]

		for x := range width {

			if x >= currentLine.GetContentLength() {
				win.Screen.SetContent(x, y, 0, nil, win.screenStyle)
				continue
			}

			bufferWithoutGap := currentLine.GetBufferWithoutGap()
			b := bufferWithoutGap[x]

			win.Screen.SetContent(x, y, rune(b), nil, win.screenStyle)
		}
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

			win.events <- "insert"

		case tcell.KeyBackspace, tcell.KeyBackspace2:
			win.Tab.BackDelete()
			win.events <- "backspace"

		case tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight:
			win.HandleDirection(ev.Key())

		case tcell.KeyTab:
			win.Tab.InsertString("te")
			win.events <- "insert_string"
		}
		fmt.Print("\033[20;0H\033[K")
		fmt.Printf("Start,End: %d,%d\t", win.Tab.Lines[0].GapStart, win.Tab.Lines[0].GapEnd)
		fmt.Printf("%v\t\t\t", win.Tab.Lines[0].Buffer)
		fmt.Println(win.Tab.Lines[0].GetString())
	}
}

func (win *Window) Quit() {
	win.Screen.Fini()
	os.Exit(0)
}
