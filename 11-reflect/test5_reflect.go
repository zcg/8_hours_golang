package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (t User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", t)

}
func doFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("Type of input is", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("Type of input is", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v = %v\n", filed.Name, filed.Type, value)
	}
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)

		fmt.Printf("%s:%v\n", m.Name, m.Type)

	}
}

func main() {
	user := User{1, "hehe", 18}
	doFieldAndMethod(user)

}
