package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new GoRoutine : i=%d\n",i)
		time.Sleep(1*time.Second)
	}
}

func main() {

	// 创建一个go程 执行newTask()
	go newTask()
	j :=0
	for {
		j++
		fmt.Printf("main GoRoutine : j=%d\n",j)
		time.Sleep(1*time.Second)
	}

}