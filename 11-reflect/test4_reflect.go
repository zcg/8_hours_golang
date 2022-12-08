package main

import (
	"fmt"
	"reflect"
)

func reflectNum(args interface{}) {
	fmt.Println("type :", reflect.TypeOf(args))
	fmt.Println("value :", reflect.ValueOf(args))

}

func main() {
	num := 3.14
	reflectNum(num)

}
