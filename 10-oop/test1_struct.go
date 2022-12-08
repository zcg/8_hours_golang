package main

import "fmt"

// 声明一种数据类型
type myint int

//定义一个结构体
type Book struct {
	title string
	auth  string
}

func printBook(book Book) {
	//传递一个副本
	book.auth = "666"
}

func changeBook(book *Book) {
	//传递一个指针
	book.auth = "777"

}
func main() {
	var book Book
	book.title = "Java"
	book.auth = "Tom"
	fmt.Printf("%v\n", book)
	printBook(book)
	fmt.Printf("%v\n", book)
	changeBook(&book)
	fmt.Printf("%v\n", book)
}
