package main

import "fmt"

func main() {
	// 写入defer关键字 类似java中的finally
	defer fmt.Println("main end")
	fmt.Println("main:hello go1")
	fmt.Println("main:hello go2")

}
