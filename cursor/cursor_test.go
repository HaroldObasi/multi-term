package cursor_test

import (
	"fmt"
	"testing"

	"github.com/HaroldObasi/multi-term/buffer"
	"github.com/HaroldObasi/multi-term/window"
)

type goToTest struct {
	name string
	x, y int
}

var goToTests = []goToTest{
	{
		name: "go to 0, 0",
		x:    0,
		y:    0,
	},
	{
		name: "go to 1, 1",
		x:    1,
		y:    1,
	},
	{
		name: "go to 10, 10",
		x:    10,
		y:    10,
	},
}

func TestGoTo(t *testing.T) {
	win, err := window.NewTestWindow()

	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}

	for _, tt := range goToTests {
		fmt.Println()

		t.Run(tt.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)

			fmt.Println("Cursor before: ", tab.Cursor)

			tab.Cursor.GoTo(tt.x, tt.y)

			fmt.Println("Cursor After: ", tab.Cursor)

			// check if the cursor is at the right position
			if tab.Cursor.X != tt.x {
				t.Errorf("cursor at x=%d, want x=%d", tab.Cursor.X, tt.x)
			}

			if tab.Cursor.Y != tt.y {
				t.Errorf("cursor at y=%d, want y=%d", tab.Cursor.Y, tt.y)
			}
		})
	}
}
