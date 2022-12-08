package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("Open File Error ", err)
		return
	}
	var r io.Reader
	r = tty
	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("Hello This is A Test\n"))
}
