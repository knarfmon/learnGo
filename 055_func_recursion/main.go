package main

import "fmt"

func main() {
	//recursion is a function that calls itself

	//a factorial
	//4!
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(5))
}
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
