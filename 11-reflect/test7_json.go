package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 100, []string{"张飞", "刘德华", "张艺谋", "刘备"}}

	// 编码过程:  结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Println(string(jsonStr))

	// 解码过程: json->结构体
	myMovie :=Movie{}
	err =json.Unmarshal(jsonStr, &myMovie)
	if err!= nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Println(myMovie)
}
