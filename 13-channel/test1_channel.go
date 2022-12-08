package main

import "fmt"

func main() {

	// 定义一个channel
	c := make(chan int)
	go func() {

		defer fmt.Println("go routine执行结束")
		fmt.Println("go routine正在运行")
		c <- 666 //将666发给c

	}()
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main routine执行结束")
}
