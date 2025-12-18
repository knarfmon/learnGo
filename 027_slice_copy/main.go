package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make([]int, 10)
	copy(b, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("******** Copied into a different underlying array ***********")
	a[0] = 100
	fmt.Println("a ", a)
	fmt.Println("b ", b)
}
