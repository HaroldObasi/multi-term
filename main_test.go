package main

import (
	"fmt"
	"os"
	"testing"
)

func GetTestApplication() Screen {
	testingFileName := "testing.txt"
	file, err := os.Create(testingFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello world")
	if err != nil {
		panic(err)
	}
	argv := []string{"multi-term", testingFileName}
	screen, _ := NewTestScreen(argv)
	return *screen
}

// func TestLineBuffer_Delete(t *testing.T) {
// 	screen	:= GetTestApplication()
// 	tabBuffer := screen.tabBuffer
// 	lineBuffer := screen.tabBuffer.lines[0]
// 	cursor := lineBuffer.cursor

// 	// text := lineBuffer.GetText()
// 	// lenText := len(text)

// 	cursor.SetPos(11, 0, tabBuffer)
	
// 	lineBuffer.Delete()
// 	fmt.Println(lineBuffer.GetText())
// }
func TestInsert(t *testing.T) {
	fmt.Println("Test Insert")
    screen := GetTestApplication()
	tabBuffer := screen.tabBuffer
	lineBuffer := tabBuffer.lines[0]
	cursor := lineBuffer.cursor
 
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
            name:        "insert 'x' at start",
            initialText: "hello",
            gapSize:     5,
            insertPos:   0,
            insertString:  "x",
            wantText:    "xhello",
            wantCursorX: 1,
        },
        {
            name:        "insert 'x' in middle",
            initialText: "hello",
            gapSize:     5,
            insertPos:   2,
            insertString:  "x",
            wantText:    "hexllo",
            wantCursorX: 3,
        },
        {
            name:        "insert 'x' at end",
            initialText: "hello",
            gapSize:     5,
            insertPos:   5,
            insertString:  "x",
            wantText:    "hellox",
            wantCursorX: 6,
        },
        {
            name:        "insert 'This is a test' at begining",
            initialText: "Hello world",
            gapSize:     5,  // Will trigger Grow
            insertPos:   0,
            insertString:  "This is a test",
            wantText:    "This is a testHello world",
            wantCursorX: 14,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            lb := NewLineBuffer(tt.initialText, tt.gapSize, &screen, cursor)

			//Print test name
			fmt.Println(tt.name)
            
            // Print initial state
            fmt.Printf("Initial state - Text: %q, GapStart: %d, GapEnd: %d\n", 
                   lb.GetText(), lb.gapStart, lb.gapEnd)
            
            // Set cursor position
            lb.cursor.SetPos(tt.insertPos, 0, screen.tabBuffer)
            
            // Do insert
			for _, r := range tt.insertString {
				lb.Insert(r)
			}
            
            // Print final state
            fmt.Printf("Final state - Text: %q, GapStart: %d, GapEnd: %d, CursorX: %d\n", 
                   lb.GetText(), lb.gapStart, lb.gapEnd, lb.cursor.x)
            
            // Check resulting text
            if got := lb.GetText(); got != tt.wantText {
                err := fmt.Errorf("after Insert(%s) at pos %d = %q, want %q", 
                        tt.insertString, tt.insertPos, got, tt.wantText)

				t.Error(err)
            }

            
            // Check cursor position
            if lb.cursor.x != tt.wantCursorX {
                err := fmt.Errorf("cursor position after insert = %d, want %d", 
                        lb.cursor.x, tt.wantCursorX)

				t.Error(err)
            }
            
            // Verify gap start is at cursor position
            if lb.gapStart != lb.cursor.x {
                err := fmt.Errorf("gap start (%d) doesn't match cursor position (%d)", 
                        lb.gapStart, lb.cursor.x)

				t.Error(err)
            }
			fmt.Println()
        })
    }
}