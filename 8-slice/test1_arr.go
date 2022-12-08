package main

import (
	"fmt"
)

func printfArr(myArr [10]int) {

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

}

func main() {
	var myArr [10]int
	myArr2 := [10]int{1, 2, 3, 4, 5}
	myArr3 := [10]string{"haha", "hahahaaaa"}

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

	for index, value := range myArr2 {
		fmt.Println("index=", index, " value=", value)
	}

	for _, value := range myArr2 {
		fmt.Println(" value=", value)
	}

	fmt.Printf("arr1 type = %T\n", myArr)
	fmt.Printf("arr2 type = %T\n", myArr2)
	fmt.Printf("arr3 type = %T\n", myArr3)
	
	fmt.Println("-----------------------")
	printfArr(myArr2)
	
	
	fmt.Println("-----------------------")

}
