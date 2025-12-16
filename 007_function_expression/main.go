package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("My first func expression")
	}
	f() //calls the generic expression that was assigned to a variable. Function are a certain type.

	g := func(x int) {
		fmt.Println("The year big brother started watching: ", x)
	}
	g(1984)

	foo()
	foo2("foo2")
}
func foo() {
	fmt.Println("go run main.go")
}
func foo2(p string) {
	fmt.Println(p)
}
