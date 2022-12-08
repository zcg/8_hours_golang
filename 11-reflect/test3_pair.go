package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {}

func (t *Book)ReadBook(){
	fmt.Println("read a Book")
}

func (t *Book)WriteBook(){
	fmt.Println("write a Book")
}

func main() {
	b := &Book{}
	var r Reader 

	r = b
	r.ReadBook()

	var w Writer

	w = r.(Writer)

	w.WriteBook()

}
