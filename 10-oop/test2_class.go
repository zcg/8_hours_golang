package main

import "fmt"


//定义一个结构体
// 如果类名的属性首字母大写 表示其它包也可以访问

type Hero struct {
	// 如果类的首字符大写,表示该属性是能够对外访问的,否则是不能够对外访问的
	Name string
	Ad  int
	Level string
}

// func (this Hero) show(){
// 	fmt.Printf("name = %v\n",this.Name)
// 	fmt.Printf("ad = %v\n",this.Ad)
// 	fmt.Printf("level = %v\n",this.Level)
// }

// func (this Hero) GetName() string{
// 	fmt.Println("Name ="+this.Name)
// 	return this.Name
// }
// func (this Hero) SetName(newName string) {
// 	// this调用的是该对象的一个副本 也即是无法修改原来的值
// 	this.Name = newName
// }

func (this *Hero) show(){
	fmt.Printf("name = %v\n",this.Name)
	fmt.Printf("ad = %v\n",this.Ad)
	fmt.Printf("level = %v\n",this.Level)
}

func (this *Hero) GetName() string{
	fmt.Println("Name ="+this.Name)
	return this.Name
}
func (this *Hero) SetName(newName string) {
	// this调用的是该对象的一个副本 也即是无法修改原来的值
	this.Name = newName
}


func main() {
	//创建一个对象
	hero := Hero{Name: "小丑",Ad: 100,Level: "1"}

	//展示
	hero.show()

	hero.SetName("蝙蝠侠")

	// 无法修改 改成指针传递即可修改 *p为取值 &p为取地址
	hero.show()

}
