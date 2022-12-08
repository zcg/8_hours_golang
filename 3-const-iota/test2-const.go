package main

import (
	"fmt"
)

// const用来定义枚举类型
const(
	// 可以在const()添加一个iota 默认是0 下面每行都会自动加1
	// iota 只能在const()下使用
	BEIJING=iota*10
	SHANGHAI
	SHENGZHEN
)

const(
	a,b=iota+1,iota+2 // iota =0 a=1 b=2
	c,d				//  c=2 d=3
	e,f  		    // e= 3 f=4

	g,h=iota*2,iota*3 //g=6 h=9
	i,k             // i=8 k=12
)

func main() {
	// const 常量 具有只读属性
	const length int = 10
	fmt.Println("length=", length)
	// length = 100 常量不允许修改

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENGZHEN=", SHENGZHEN)


	fmt.Println("a=",a,"b=",b,"c=",c,"d=",d,"e=",e,"f=",f,"g=",g,"h=",h,"i=",i,"k=",k)
}
