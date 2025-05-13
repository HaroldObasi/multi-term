package buffer_test

import (
	"bytes"
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

type lineBufferGrowTest struct {
	name             string
	originalBuffer   []byte
	wantBuffer       []byte
	originalGapStart int
	wantGapStart     int
	originalGapEnd   int
	wantGapEnd       int
}

type goToTest struct {
	name             string
	targetPosX       int
	originalBuffer   []byte
	wantBuffer       []byte
	originalGapStart int
	wantGapStart     int
	originalGapEnd   int
	wantGapEnd       int
}

type testSingleInsert struct {
	name             string
	originalBuffer   []byte
	originalGapStart int
	originalGapEnd   int
	wantBuffer       []byte
	targetValue      byte
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

var lineBufferGrowTests = []lineBufferGrowTest{
	{
		name:             "grow buffer at end position",
		originalBuffer:   []byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 0},
		wantBuffer:       []byte{49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		originalGapStart: 9,
		wantGapStart:     9,
		originalGapEnd:   9,
		wantGapEnd:       19,
	},
	{
		name:             "grow buffer at start position",
		originalBuffer:   []byte{0, 49, 50, 51, 52, 53, 54, 55, 56, 57},
		wantBuffer:       []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 54, 55, 56, 57},
		originalGapStart: 0,
		wantGapStart:     0,
		originalGapEnd:   0,
		wantGapEnd:       10,
	},
}

var goToTests = []goToTest{
	{
		name:             "go to start position from end",
		targetPosX:       0,
		originalBuffer:   []byte{104, 97, 114, 111, 108, 100, 0, 0, 0, 0},
		wantBuffer:       []byte{0, 0, 0, 0, 104, 97, 114, 111, 108, 100},
		originalGapStart: 6,
		wantGapStart:     0,
		originalGapEnd:   9,
		wantGapEnd:       3,
	},
	{
		name:             "go to end position from start",
		targetPosX:       9,
		originalBuffer:   []byte{0, 0, 0, 0, 104, 97, 114, 111, 108, 100},
		wantBuffer:       []byte{104, 97, 114, 111, 108, 100, 0, 0, 0, 0},
		originalGapStart: 0,
		wantGapStart:     6,
		originalGapEnd:   3,
		wantGapEnd:       9,
	},
	{
		name:             "go to middle position from start",
		targetPosX:       5,
		originalBuffer:   []byte{0, 0, 0, 0, 104, 97, 114, 111, 108, 100},
		wantBuffer:       []byte{104, 97, 114, 111, 0, 0, 0, 0, 108, 100},
		originalGapStart: 0,
		wantGapStart:     5,
		originalGapEnd:   3,
		wantGapEnd:       8,
	},
}

var testSingleInsertTests = []testSingleInsert{
	{
		name:             "insert '48' at 1",
		originalBuffer:   []byte{49, 0, 0, 0, 0, 0, 116, 101, 115, 116},
		wantBuffer:       []byte{49, 48, 0, 0, 0, 0, 116, 101, 115, 116},
		originalGapStart: 1,
		originalGapEnd:   5,
		targetValue:      48,
	},
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

func TestGrowBuffer(t *testing.T) {
	win, err := window.NewTestWindow()

	if err != nil {
		t.Fatalf("Error creating new window: %v", err)
	}

	for _, tt := range lineBufferGrowTests {
		fmt.Println()
		t.Run(tt.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)

			line := tab.Lines[0]

			// Insert original string first
			line.Buffer = tt.originalBuffer
			line.GapStart = tt.originalGapStart
			line.GapEnd = tt.originalGapEnd

			fmt.Println("Line before Grow: ", line)
			line.GrowBuffer()
			fmt.Println("Line After Grow: ", line)

			// check if the text is as expected
			if !bytes.Equal(line.Buffer, tt.wantBuffer) {
				t.Errorf("got '%q', want '%q'", line.Buffer, tt.wantBuffer)
			}

			// check if the cursor is at the right position
			if line.GapStart != tt.wantGapStart {
				t.Errorf("gap start at x=%d, want x=%d", line.GapStart, tt.wantGapStart)
			}

			if line.GapEnd != tt.wantGapEnd {
				t.Errorf("gap end at x=%d, want x=%d", line.GapEnd, tt.wantGapEnd)
			}
		})
	}
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

			line := tab.Lines[0]

			// Insert original string first
			line.Buffer = tt.originalBuffer
			line.GapStart = tt.originalGapStart
			line.GapEnd = tt.originalGapEnd

			line.GoTo(tt.targetPosX)

			// check if the text is as expected
			if !bytes.Equal(line.Buffer, tt.wantBuffer) {
				t.Errorf("got '%q', want '%q'", line.Buffer, tt.wantBuffer)
			}

			// check if the cursor is at the right position
			if line.GapStart != tt.wantGapStart {
				t.Errorf("gap start at x=%d, want x=%d", line.GapStart, tt.wantGapStart)
			}

			if line.GapEnd != tt.wantGapEnd {
				t.Errorf("gap end at x=%d, want x=%d", line.GapEnd, tt.wantGapEnd)
			}
		})
	}
}

func TestSingleInsert(t *testing.T) {
	win, err := window.NewTestWindow()

	if err != nil {
		t.Fatalf("Error creating test window")
	}

	for _, test := range testSingleInsertTests {
		fmt.Println()

		t.Run(test.name, func(t *testing.T) {
			tab := buffer.NewTabBuffer(win.Screen)
			line := tab.Lines[0]

			line.Buffer = test.originalBuffer
			line.GapStart = test.originalGapStart
			line.GapEnd = test.originalGapEnd

			line.InsertRune(rune(test.targetValue))

			if !bytes.Equal(line.Buffer, test.wantBuffer) {
				fmt.Println(line.Buffer)
				t.Errorf("got '%q', want %q", line.Buffer, test.wantBuffer)
				fmt.Println(test.wantBuffer)
			}

		})

	}
}
