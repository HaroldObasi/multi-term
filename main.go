package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

// func HandleRender(events chan string, tScreen tcell.Screen, tb *TabBuffer) {
// 	for range events {
// 		Render(tScreen, tb)
// 	}
// }
func HandleRender(events chan string, screen *Screen) {
	for range events {
		Render(screen.tScreen, screen.tabBuffer)
	}
}


func Render(tScreen tcell.Screen, tb *TabBuffer) {
	tb.ClearTabArea()
	x := 0
	y := tb.GetUpperBound()
	for _, line := range tb.GetValidLines(){
		for _, char := range line.GetBufferWithoutGap() {
			if char == '\n'{
				tScreen.SetContent(x, y, char, nil, tcell.StyleDefault)
				x = 0
				y++
			} else if char == '\t'{
				x += 4
			} else {
				tScreen.SetContent(x, y, char, nil, tcell.StyleDefault)
				x++
			}
		}
		y++
		x = 0
	}
	tScreen.Show()
}

func main() {

	argv := os.Args
	screen, err := NewScreen(argv)
	screen.tabBuffer.WriteFileToScreen()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	events := make(chan string)
	

	go HandleRender(events, screen)
	// go HandleRender(events, screen.tScreen, screen.tabBuffer)
	HandleEvents(screen, events)
}
