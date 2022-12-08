package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4}

	// 打印原始切片
	fmt.Printf("原始切片%v\n", slice)

	// 切片赋值 从索引值0到2
	s1 := slice[0:2]
	fmt.Println(s1)

	// 切片赋值 索引下限为0
	s2 := slice[:3]
	fmt.Println(s2)

	// 切片赋值 索引上限为len-1
	s3 := slice[2:]
	fmt.Println(s3)


	s1[0] =100
	fmt.Printf("s1=%v\n", s1)
	fmt.Printf("原始切片%v\n", slice)


	//copy 可以将底层数组的slice一起进行拷贝 多的用0补
	s4 :=make([]int, 6)
	copy(s4, slice)
	fmt.Printf("%v\n",s4)
}
