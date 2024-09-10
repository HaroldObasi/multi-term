package main

import (
	"os"
)

type File struct {
	path string
}

func NewFile(path string) *File {
	file := File{path: path}
	return &file
}

func (f *File) ReadFile() []byte {
	data, err := os.ReadFile(f.path)

	if err != nil {
		panic(err)
	}

	return (data)
}

func (f *File) Save(tb *TabBuffer) {
	// read through all the lines in the tab buffer and add thier runes to a buffer array
	// then write the buffer array to the file

	validLines := tb.GetValidLines()

	var buffer []byte

	for _, line := range validLines {
		runes := line.GetRunes()
		buffer = append(buffer, []byte(runes)...)
	}

	os.WriteFile(f.path, buffer, 0644)
}
