package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len =%d,cap=%d,numbers=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len =%d,cap=%d,numbers=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len =%d,cap=%d,numbers=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)
	fmt.Printf("len =%d,cap=%d,numbers=%v\n", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len =%d,cap=%d,numbers2=%v\n", len(numbers2), cap(numbers2), numbers2)
	
	numbers2=append(numbers2,1)
	fmt.Printf("len =%d,cap=%d,numbers2=%v\n", len(numbers2), cap(numbers2), numbers2)

}
