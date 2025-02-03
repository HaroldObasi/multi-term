package buffer_test

import (
	"fmt"
	"testing"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/HaroldObasi/multi-term/window"
)

type lineInsertTest struct {
	name           string
	insertPosX     int
	insertPosY     int
	insertString   string
	wantText       string
	wantCursorX    int
	originalString string
}

var lineInsertTests = []lineInsertTest{
	{
		name:           "insert 'c' in 'ab' at pos 2",
		insertPosX:     2,
		insertPosY:     0,
		insertString:   "c",
		wantText:       "abc",
		wantCursorX:    3,
		originalString: "ab",
	},
	{
		name:           "insert 'x' in 'ab' at pos 0",
		insertPosX:     0,
		insertPosY:     0,
		insertString:   "x",
		wantText:       "xab",
		wantCursorX:    1,
		originalString: "ab",
	},
	{
		name:           "insert xx in 'hello world' at pos 0",
		insertPosX:     0,
		insertPosY:     0,
		insertString:   "xx",
		wantText:       "xxhello world",
		wantCursorX:    2,
		originalString: "hello world",
	},
}

type lineGoLeftInsert struct {
	name           string
	originalBuffer []byte
	gapStart       int
	gapEnd         int
}

func TestInsert(t *testing.T) {
	win, err := window.NewTestWindow()

	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}

	for _, tt := range lineInsertTests {
		fmt.Println()
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

			// check if the text is as expected
			if line.GetString() != tt.wantText {
				t.Errorf("got '%q', want '%q'", line.GetString(), tt.wantText)
			}

			// check if the cursor is at the right position
			if tab.Cursor.X != tt.wantCursorX {
				t.Errorf("cursor at x=%d, want x=%d", tab.Cursor.X, tt.wantCursorX)
			}
		})
	}

}
