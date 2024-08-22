package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

type Window struct {
	x int
	y int
}

func (w *Window) Write(char rune, s tcell.Screen) {
	// style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue.TrueColor()).Background(tcell.ColorWhite)
	s.SetContent(w.x, w.y, char, nil, tcell.StyleDefault)

	w.x++
	s.Show()
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func displayHelloWorld(s tcell.Screen) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue.TrueColor()).Background(tcell.ColorWhite)
	emitStr(s, w/2-7, h/2, style, "Hello, World!")
	emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
	s.Show()
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {

	w := Window{
		x: 0,
		y: 0,
	}

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	s.SetStyle(defStyle)
	s.ShowCursor(0, 0)

	// displayHelloWorld(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			// displayHelloWorld(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
			if ev.Key() == tcell.KeyCtrlL {
				w.Write('a', s)

				// displayHelloWorld(s)
			}
			if ch := ev.Rune(); ch != 0 {
				w.Write(ch, s)
			}

		}
	}
}