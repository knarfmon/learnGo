package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 20; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println(" ")
	fmt.Println("***************  for loop with random number")
	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("Step %v, random number is %v\n", i, x)
	}
	fmt.Println(" ")
	fmt.Println("***************  switch loop with random number")
	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("iteration %v is %v\n", i, x)
		case 1:
			fmt.Printf("iteration %v is %v\n", i, x)
		case 2:
			fmt.Printf("iteration %v is %v\n", i, x)
		case 3:
			fmt.Printf("iteration %v is %v\n", i, x)
		case 4:
			fmt.Printf("iteration %v is %v\n", i, x)
		}
	}

}
