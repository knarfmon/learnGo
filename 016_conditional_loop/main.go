package main

import (
	"fmt"
)

func main() {
	var y int = 5
	for y > 0 {
		fmt.Printf("y is %v\n", y)
		y--
	}

	fmt.Println("************   infinite for break loop")
	y = 1
	for {
		fmt.Printf("y is %v\n", y)
		if y >= 5 {
			break
		}
		y++
	}

	fmt.Println("************  ; for loop *************")
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Printf("x is %v\n", x)

	}

}
