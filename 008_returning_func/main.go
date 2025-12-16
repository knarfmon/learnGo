package main

import (
	"fmt"
)

func main() {
	//can put in main()
	b := func(z string) {
		fmt.Println(z)
	}
	b("James")

	//or an anonymous string
	func(k string) {
		fmt.Println(k)
	}("Alfred")

	x := foo()
	y := foo()
	fmt.Printf("%T\t %T\n", x, y)

	x("jenny")
	y("frank")

}

func foo() func(string) {
	return func(z string) { //returning a value of a type, function are first class citizens
		fmt.Println(z)
	}

}
