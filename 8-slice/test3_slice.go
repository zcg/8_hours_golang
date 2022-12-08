package main

import (
	"fmt"
)

func main() {
	// 1. 声明一个slice 并初始化 长度为3
	// slice := []int{1, 2, 3}
	// 2. 声明一个slice 并不初始化 长度为0
	// var slice []int
	// 3. 在函数体内定义slice 长度为3 默认是0
	slice := make([]int, 3)
	fmt.Printf("len = %d,slice is %v\n", len(slice), slice)

	if slice == nil {
		fmt.Println("slice是空的")
	} else {
		fmt.Printf("slice是有空间的")
	}
}
