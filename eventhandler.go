package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func HandleEvents(screen *Screen) {
	tScreen := screen.tScreen
	for {
		switch ev := tScreen.PollEvent().(type) {
		case *tcell.EventResize:
			tScreen.Sync()
		case *tcell.EventKey:
			switch ev.Key() {

			case tcell.KeyEscape:
				tScreen.Fini()
				os.Exit(0)

			case tcell.KeyRune:
				ch := ev.Rune()
				HandleInsertRune(screen, ch)

			case tcell.KeyBackspace, tcell.KeyBackspace2:
				HandleBackspace(screen)
			}
		}
	}
}

func HandleInsertRune(screen *Screen, r rune) {
	// figure out where the cursor is
	// insert the rune at the cursor
	tb := screen.tabBuffer

	// lb := tb.lines[0]

	cursor := tb.cursor
	line := tb.lines[cursor.y]
	line.Insert(r)
}

func HandleBackspace(screen *Screen) {
	tb := screen.tabBuffer
	cursor := tb.cursor
	line := tb.lines[cursor.y]
	line.Delete()
}
