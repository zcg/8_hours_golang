package main

import "fmt"

type Human struct {
	name string
	sex  string
}

type SuperMan struct {
	Human // SuperMan 类继承了Human类的方法
	level int
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()")
}

// --------

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()")
}

func (this *SuperMan) Print() {
	fmt.Println("name is ", this.name)
	fmt.Println("sex is ", this.sex)
	fmt.Println("level is ", this.level)
}

// --------

func main() {
	// 定义一个对象human
	h := Human{"Tom", "Male"}
	h.Eat()
	h.Walk()
	// 定义一个子对象SuperMan
	// s := SuperMan{Human{"su", "feMale"},88}
	var s SuperMan
	s.Human = Human{"su", "feMale"}
	s.level = 88
	s.Eat()
	s.Walk()
	s.Fly()
	s.Print()
}
