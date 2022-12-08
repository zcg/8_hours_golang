package main

import "fmt"

// 返回一个返回值

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

// 返回多个返回值 匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	d := 200
	return c, d
}

// 返回多个返回值 有形参名称的
func foo3(a string, b int) (r1 int, r2 string) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = 1000
	r2 = "哈哈哈"

	return r1, r2
}

func main() {
	c := foo1("Hello", 10)
	fmt.Println("c=", c)

	ret1, ret2 := foo2("Hello", 10)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	r1, r2 := foo3("hello", 10)

	fmt.Println("r1=", r1, "r2=", r2)

}
