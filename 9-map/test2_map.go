package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	cityMap["London"] = "London"
	cityMap["Paris"] = "Paris"
	cityMap["New York"] = "New York"
	cityMap["Chicago"] = "Chicago"

	// 遍历
	for key, value := range cityMap {
		fmt.Println( key,":",value)

	}

	fmt.Println("--------")

	//删除
	delete(cityMap, "Chicago")

	//修改
	cityMap["New York"] = "New York22"
	fmt.Println("--------")

	//遍历
	for key, value := range cityMap {
		fmt.Println( key,":",value)
	}
}