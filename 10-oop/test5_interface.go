package main

import (
	"fmt"
)

// interface是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)
	//给interface提供类型断言的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string")

	} else {
		fmt.Printf("arg is a string value is %T\n",value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "David"}
	myFunc(book)
	myFunc("100")
	myFunc(1)
	myFunc(true)
	myFunc(3.14)
}
