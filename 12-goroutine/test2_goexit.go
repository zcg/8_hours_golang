package main

import (
	"fmt"
	"time"
)

func main() {

	/* 	//go 创建一个承载一个形参为空 返回值为空的一个函数
	   	go func(){
	   		defer fmt.Println("A.defer")
	   		func(){
	   			defer fmt.Println("B.defer")
	   			fmt.Println("B")
	   		}()
	   		fmt.Println("A")
	   	}()

	   	for{
	   		time.Sleep(1 * time.Second)
	   	} */

	go func(a int, b int) bool {
		fmt.Println("a = ", a, ",b= ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}

}
