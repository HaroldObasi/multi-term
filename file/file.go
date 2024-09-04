package file

import "os"

type File struct {
	path string
}

//write an array of runes to the file
func (f *File) Write(chars []rune) {
	//write the runes to the file

	err := os.WriteFile(f.path, []byte(string(chars)), 0644)

	if err != nil {
		panic(err)
	}

}

// func (f *File) Save(tb *TabBuffer) {
// 	//save the tab buffer to the file
// 	//get the serialized data from the tab buffer
// 	serialized := tb.Serialize()
// 	//write the serialized data to the file
// 	f.Write(serialized)
// }