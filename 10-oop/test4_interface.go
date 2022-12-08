package main

import (
	"fmt"
)

// 接口 本质是一个指针

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// -------
func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}

// -------

func main() {

	// var animal AnimalIF //接口父类类型 父类指针
	// animal =&Cat{"white"}
	// animal.Sleep() //调用的就是cat的sleep
	// color := animal.GetColor()
	// fmt.Println(color)
	// animalType := animal.GetType()
	// fmt.Println(animalType)
	// animal =&Dog{"black"}
	// animal.Sleep() //调用的就是cat的sleep
	// color1 :=animal.GetColor()
	// fmt.Println(color1)
	// animalType1 := animal.GetType()
	// fmt.Println(animalType1)

	cat := &Cat{"white"}
	dog := &Dog{"black"}

	showAnimal(cat)
	showAnimal(dog)

}
