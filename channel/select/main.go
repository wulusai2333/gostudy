package main

import "fmt"

/*
	select的使用
*/
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select { //select是对满足条件的case随机执行的, 这里输出为0 2 4 6 8
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:

		}
	}
}
