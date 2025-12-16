package main

import (
	"fmt"
	"math/rand"
)

//Sequence, Conditional, Iterative

func init() {
	fmt.Println("func init() begins initialization before func main()")
}

func main() {
	//Sequence
	fmt.Println("***********The statements are SEQUENTIAL *************")
	fmt.Println("this is the first statement")
	fmt.Println("this is the second statement")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v \n", x, y)

	//Conditional
	//if statements
	// switch statements
	fmt.Println("***********The statements are CONDITIONAL *************")
	if x < 42 {
		fmt.Println("x is less than 42")
	}
	if x < 42 {
		fmt.Println("x is less than 42")
	} else {
		fmt.Println("x is greater than or equal to 42")
	}

	if x < 42 {
		fmt.Println("x is less than 42")
	} else if x == 42 {
		fmt.Println("x is equal to 42")
	} else {
		fmt.Println("x is greater than 42")
	}
	fmt.Println("***********+ CONDITIONAL OPERATORS   *************")

	if x < 42 && y < 42 {
		fmt.Println("&& requires both sides of equality to be true")
	}
	if x > 30 || x < 42 {
		fmt.Println("|| requires only one side of the equality to be true")
	}
	if x != 42 && y != 10 {
		fmt.Println("! reverses the sign, means not")
	}
	fmt.Println("***********+ Statement-Statement Idiom   *************")
	//The true/false conditional expression is preceded by a simple statement in an IF statement
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER than or equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS than or equal to x which is %v\n", z, x)
	}
	fmt.Println("***********+ Comma ok Idiom   *************")
	// used for checking a value in a map

	fmt.Println("***********+ Switch Statements   *************")

	switch {
	case x < 42:
		fmt.Println("x < 42")
	case x == 42:
		fmt.Println("x == 42")
	case x > 42:
		fmt.Println("x > 42")
	default:
		fmt.Println("This is the default statement")
	}

	switch x { // No fall through
	case 40:
		fmt.Println("x = 40")
		fallthrough // Not sure why you would use this
	case 41:
		fmt.Println("x = 41")
	case 42:
		fmt.Println("x = 42")
	case 43:
		fmt.Println("x = 43")
	default:
		fmt.Println("Default")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("***********+ Loops   *************")

	fmt.Println("for init; condition; post { }")

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y >= 10 {
			break
		}
		y++
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("for condition {}")

	for y < 20 {
		fmt.Printf("y = %v \n", y)
		y++
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("for(;;) {}   Traditional statement Idiom  ")

	for i := 0; i < 40; i++ {
		if i%2 != 0 {
			continue // Notice the use of continue with if statement
		}
		fmt.Printf("%v ", i)

	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("***********+ Nesting Loops   *************")

	for i := 0; i < 5; i++ { //Outer loop
		fmt.Printf("--")
		for j := 0; j < 5; j++ { //Inner loop runs 5 times for each cycle of outer loop
			fmt.Printf("outer loop %v\t inner loop %v\n", i, j)
		}
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("***********+ Range Loops   *************")
	fmt.Println("***********+ Data Structures and Slice   *************")
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("xi[%d] = %v\n", i, v)
	}
	fmt.Println("")
	m := map[string]int{ //string is the lookup value (key, k) and int is value (v) returned.
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m { // Maps do not return in predictable order.
		fmt.Printf("%v: %v\n", k, v)
	}

	fmt.Println("")
	fmt.Println("***********+ Another Range Loop   *************")
	sum := 0
	for _, v := range xi {
		sum += v
	}
	fmt.Println(sum)
} //When func main ends the program stops.
