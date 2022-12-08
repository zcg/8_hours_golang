package main

import "fmt"

func PrintfArr(myArr []int) {
	myArr[0] = 100
	for _, value := range myArr {
		fmt.Println("value is", value)
	}
}

func main() {
	myArr := []int{1, 2, 3, 4, 5, 6} //动态数组 切片slice
	fmt.Printf("myArr type is %T\n", myArr)
	for _, value := range myArr {
		fmt.Println("value is", value)
	}
	fmt.Println("--------")
	PrintfArr(myArr)
}
