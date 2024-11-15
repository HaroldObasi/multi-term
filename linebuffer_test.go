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

func TestLineBuffer_Delete(t *testing.T) {
	screen	:= GetTestApplication()
	tabBuffer := screen.tabBuffer
	lineBuffer := screen.tabBuffer.lines[0]
	cursor := lineBuffer.cursor

	// text := lineBuffer.GetText()
	// lenText := len(text)

	cursor.SetPos(11, 0, tabBuffer)
	
	lineBuffer.Delete()
	fmt.Println(lineBuffer.GetText())
}