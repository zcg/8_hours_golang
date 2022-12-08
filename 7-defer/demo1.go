package main

import "fmt"

func main() {

	// defer执行顺序按照 栈的顺序 先进后出 后进先出的顺序执行


	// 有return 的先执行return的,后执行defer的
	
	defer func1()
	defer func2()
	defer func3()

}


func func1(){
	fmt.Printf("A")
}

func func2(){
	fmt.Printf("B")
}

func func3(){
	fmt.Printf("C")
}