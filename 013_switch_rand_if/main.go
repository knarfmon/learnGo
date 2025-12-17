package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("init")
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	fmt.Println("main")
	fmt.Println("This is a random integer")
	x := rand.Intn(251) // 251 to include 250
	fmt.Printf("x = %d\n", x)
	fmt.Println("+++++++++ Using a switch/case statement  **********")
	switch {
	case x >= 0 && x <= 100:
		fmt.Println("Between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 100 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("This is more than 250")
	}

	for i := 0; i < 40; i++ { // only includes 0,1,2
		fmt.Println(rand.Intn(3))
	}
	fmt.Println("+++++++++ Using a for loop statement  **********")
	fmt.Printf("")

	if x <= 100 {
		fmt.Println("Less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("This was more than 250")
	}

	fmt.Println("+++++++++ Using an if statement  **********")
	a, b := rand.Intn(10), rand.Intn(10)
	fmt.Printf("a = %d, b = %d\n", a, b)
	if a < 4 && b < 4 {
		fmt.Println("Both less than 4")
	} else if a > 6 && b > 6 {
		fmt.Println("Both greater than 6")
	} else if a >= 4 && a <= 6 {
		fmt.Println("a is greater than or equal to 4 and less than or equal to 6")
	} else if b != 5 {
		fmt.Println("b is not equal to 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}
	fmt.Println("+++++++++ Using Switch again statement  **********")
	fmt.Printf("")

	switch {
	case a < 4 && b < 4:
		fmt.Println("Both less than 4")
	case a > 6 && b > 6:
		fmt.Println("Both greater than 6")
	case a >= 4 && b <= 6:
		fmt.Println("a is greater than or equal to 4 and less than or equal to 6")
	case b != 5:
		fmt.Println("b is not equal to 5")
	default:
		fmt.Println("none of the previous cases were met")
	}

}
