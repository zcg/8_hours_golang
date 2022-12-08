package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i

		}
		// close可以关闭一个channel
		close(c)
	}()

	// for {
	// 	if data, ok := <-c; ok {
	// 		fmt.Println("data = ", data)
	// 	} else {
	// 		break
	// 	}
	// }


	// 使用range 迭代操作channel
	for data := range c{
		fmt.Println("data = ", data)
	}
	fmt.Println("main finised")
}
