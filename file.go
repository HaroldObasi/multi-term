package main

import "os"

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