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

			case tcell.KeyCtrlS:
				// save the file
				HandleSave(screen)

			case tcell.KeyUp, tcell.KeyDown, tcell.KeyLeft, tcell.KeyRight:
				HandleDirection(screen, ev.Key())

			case tcell.KeyEnter:
				HandleReturn(screen)

			case tcell.KeyRune:
				ch := ev.Rune()
				HandleInsertRune(screen, ch)

			case tcell.KeyBackspace, tcell.KeyBackspace2:
				HandleBackspace(screen)
			}
		}
	}
}

func HandleSave(screen *Screen) {
	tb := screen.tabBuffer
	tb.file.Save(tb)

	screen.WriteDebug("File saved", 4)
	// file := tb.file
	// file.Save(tb)
}

func HandleInsertRune(screen *Screen, r rune) {
	// figure out where the cursor is
	// insert the rune at the cursor
	tb := screen.tabBuffer
	cursor := tb.cursor
	line := tb.GetLine(cursor.y)
	line.Insert(r)
}

func HandleBackspace(screen *Screen) {
	tb := screen.tabBuffer
	cursor := tb.cursor
	line := tb.GetLine(cursor.y)
	line.Delete()
}

func HandleReturn(screen *Screen) {
	tb := screen.tabBuffer

	cursor := tb.cursor
	line := tb.GetLine(cursor.y)
	line.Insert('\n')

	tb.AddLine("")
}

func HandleDirection(screen *Screen, key tcell.Key) {
	tb := screen.tabBuffer
	cursor := tb.cursor

	switch key {
	case tcell.KeyUp:
		cursor.SetPos(cursor.x, cursor.y-1, tb)

	case tcell.KeyDown:
		if cursor.y < tb.Len() {
			cursor.SetPos(cursor.x, cursor.y+1, tb)
		}

	case tcell.KeyLeft:
		cursor.SetPos(cursor.x-1, cursor.y, tb)

	case tcell.KeyRight:
		line := tb.GetLine(cursor.y)
		if cursor.x < line.Len()-1 {
			cursor.SetPos(cursor.x+1, cursor.y, tb)
		}
	}
}
