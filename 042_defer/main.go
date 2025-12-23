package main

import "fmt"

func main() {
	defer foo()
	bar()
	//defer used for closing services, so deferred function runs after everything else in block, all surrounding functions return.
}

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
