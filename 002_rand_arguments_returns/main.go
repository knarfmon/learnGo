package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("My new random number is", rand.Intn(10)+1)
	fmt.Println("My new random number is", rand.Intn(10)+1)
	fmt.Println("My new random number is", rand.Intn(10)+1)
	fmt.Println("My new random number is", rand.Intn(10)+1)
	fmt.Println("My new random number is", rand.Intn(10)+1)
	fmt.Println("My new random number is", rand.Intn(10)+1)

	fmt.Println(math.Pi)

	fmt.Println("2 + 3 = ", add(2, 3))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

}

func add(a, y int) int {
	return a + y
}
func swap(x, y string) (string, string) {
	return y, x
}
