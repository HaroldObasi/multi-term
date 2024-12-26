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

// func TestLineGoTo(t *testing.T) {
// 	win, err := window.NewTestWindow()

// 	if err != nil {
// 		t.Fatalf("Error creating new window: %v", err)
// 	}

// 	tests := []struct {
// 		name  string
// 		goToX int
// 	}{
// 		{
// 			name:  "go to 0",
// 			goToX: 0,
// 		},
// 	}

// 	fmt.Println("Tests: ", tests)

// 	tab := buffer.NewTabBuffer(win.Screen)

// 	tab.Cursor.GoTo(0, 0)

// 	tab.InsertRune('a')
// 	tab.InsertRune('b')
// 	tab.InsertRune('c')

// 	tab.Cursor.GoTo(0, 0)

// 	tab.InsertRune('x')
// 	tab.InsertRune('y')
// 	tab.InsertRune('z')

// 	fmt.Println(tab.Lines[0])
// }

func TestInsert(t *testing.T) {
	win, err := window.NewTestWindow()

	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}

	tests := []struct {
		name           string
		initialText    string
		gapSize        int
		insertPosX     int // cursor position where we'll insert
		insertPosY     int // cursor position where we'll insert
		insertString   string
		wantText       string
		wantCursorX    int
		originalString string
	}{
		// {
		// 	name:         "insert 'x' in empty buffer",
		// 	insertPosX:   0,
		// 	insertPosY:   0,
		// 	insertString: "x",
		// 	wantText:     "x",
		// 	wantCursorX:  1,
		// },
		// {
		// 	name:         "insert '1234567890' in empty buffer",
		// 	insertPosX:   0,
		// 	insertPosY:   0,
		// 	insertString: "1234567890",
		// 	wantText:     "1234567890",
		// 	wantCursorX:  1,
		// },
		// {
		// 	name:         "insert '1234567890a' in empty buffer",
		// 	insertPosX:   0,
		// 	insertPosY:   0,
		// 	insertString: "1234567890a",
		// 	wantText:     "1234567890a",
		// 	wantCursorX:  1,
		// },
		{
			name:           "insert 'x' in 'ab' at pos 0",
			insertPosX:     0,
			insertPosY:     0,
			insertString:   "x",
			wantText:       "xab",
			wantCursorX:    1,
			originalString: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)

			// Insert original string first
			if tt.originalString != "" {
				for _, ch := range tt.originalString {
					tab.InsertRune(ch)
				}
			}

			tab.Cursor.GoTo(tt.insertPosX, tt.insertPosY)

			line := tab.Lines[0]

			fmt.Println("Line Before Insert: ", line)

			for _, ch := range tt.insertString {
				tab.InsertRune(ch)
			}

			fmt.Println("Line After Insert: ", line)

			// check if
		})
	}

}
