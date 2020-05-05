package main

import (
	"fmt"
)

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWrite interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read Data")
}
func (f *File) Write() {
	fmt.Println("Write Data")
}

func Test(rw ReaderWrite) {
	rw.Read()
	rw.Write()
}

func main() {
	var f *File
	Test(f)
}
