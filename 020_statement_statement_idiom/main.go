package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 { //a quick if statement
			fmt.Println("x is ", x, " a quick if statement")
		}
	}
}
