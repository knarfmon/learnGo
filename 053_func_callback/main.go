package main

import "fmt"

func main() {
	//passing a function in as an argument is a callback
	x := doMath(1, 1, add)
	fmt.Println(x)

	x = doMath(5, 1, subtract)
	fmt.Println(x)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
