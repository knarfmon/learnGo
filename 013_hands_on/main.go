package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(251) // 251 to include 250
	fmt.Printf("x = %d\n", x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("Between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 100 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("This is the default statement")
	}
	for i := 0; i < 40; i++ { // only includes 0,1,2
		fmt.Println(rand.Intn(3))
	}
}
