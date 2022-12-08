package main

import "fmt"

// :=只能生成方法内局部变量 不能作为全局变量
// gc:=234

func main() {
	// 1. 声明一个变量 默认值为0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a= %T\n", a)
	// 2. 声明一个变量 赋值为10
	var b int = 10
	fmt.Println("b=", b)
	fmt.Printf("type of b= %T\n", b)
	// 3. 声明一个变量 省去数据类型 让go自己去推导
	var c = "hello"
	fmt.Println("c=", c)
	fmt.Printf("type of c= %T\n", c)
	// 4. 声明一个变量 省去var关键字 直接自动匹配
	d := 100
	fmt.Println("d=", d)
	fmt.Printf("type of d= %T\n", d)
	e := "hello"
	fmt.Println("e=", e)
	fmt.Printf("type of e= %T\n", e)
	f := 3.14
	fmt.Println("f=", f)
	fmt.Printf("type of f= %T\n", f)

	// 5. 声明多个变量
	var xx , yy =100,200
	fmt.Println("xx=", xx,"yy=", yy)

	// 6. 声明多个变量 括号写法
	var(
		vv int=123
		jj string ="haha"
	)
	fmt.Println("vv=", vv,"jj=", jj)
	
}
