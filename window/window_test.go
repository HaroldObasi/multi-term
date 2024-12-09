package window_test

import (
	"fmt"
	"testing"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/HaroldObasi/multi-term/window"
)

// func TestNewWindow(t *testing.T) {
// 	_, err := window.NewTestWindow()
// 	if err != nil {
// 		t.Fatalf("Error creating new window: %v", err)
// 	}
// }

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
        insertString   string
        wantText     string
        wantCursorX  int
    }{
        {
            name:        "insert 'x' in empty buffer",
            insertPos:   0,
            insertString:  "x",
            wantText:    "x",
            wantCursorX: 1,
        },
        {
            name:        "insert '1234567890' in empty buffer",
            insertPos:   0,
            insertString:  "1234567890",
            wantText:    "1234567890",
            wantCursorX: 1,
        },
    }

	for _, tt := range tests {	
		t.Run(tt.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)

			for _, ch := range tt.insertString{
				tab.InsertRune(ch)
			}
			
			line := tab.Lines[0]
			fmt.Println("Line: ", line)
		})
	}

}