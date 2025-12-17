package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("Outer Loop, Step %v\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("\t\tInner Loop, Step %v\n", j)
		}
	}
}
