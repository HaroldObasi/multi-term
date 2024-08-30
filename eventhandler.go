package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func HandleEvents(screen *Screen){
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

				// case tcell.KeyRune:
				// 	ch := ev.Rune()
				// 	w.Write(ch, s)

				case tcell.KeyBackspace, tcell.KeyBackspace2:
					// screen.Delete(s)
					screen.WriteDebug( "Backspace")
			}
		}
	}
}