package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[3:4:4]
	fmt.Println(s, cap(s))
}
