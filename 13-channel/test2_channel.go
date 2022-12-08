package main

import "fmt"

func main() {

	c := make(chan int, 3) // 带有缓冲的通道 长度为3

	fmt.Println("len(c)", len(c), " cap(c)", cap(c), "c=", c)

	go func() {
		defer fmt.Println("子go 程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("发送子go程正在运行,发送的元素:i=", i, "len(c)", len(c), "cap(c)", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("接受子go程运行完成,接收到的元素:num=", num, "len(c)", len(c), "cap(c)", cap(c))

	}

	fmt.Println("main 程结束")

}
