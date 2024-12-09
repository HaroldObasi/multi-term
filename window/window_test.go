package window_test

import (
	"fmt"
	"testing"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/HaroldObasi/multi-term/window"
)

func TestNewWindow(t *testing.T) {
	_, err := window.NewTestWindow()
	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}
}

func TestInsert(t *testing.T){
	win, err := window.NewTestWindow()


	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}

	
	tests := []struct {
        name          string
        initialText   string
        gapSize      int
        insertPos    int  // cursor position where we'll insert
        insertRune   rune
        wantText     string
        wantCursorX  int
    }{
        {
            name:        "insert 'x' in empty buffer",
            insertPos:   0,
            insertRune:  'x',
            wantText:    "x",
            wantCursorX: 1,
        },
    }

	for _, tt := range tests {	
		t.Run(tt.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)
			tab.InsertRune(tt.insertRune)
			line := tab.Lines[0]
			fmt.Println("Line: ", line)
		})
	}

}