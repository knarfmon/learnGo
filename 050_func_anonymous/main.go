package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous function")
	}() // here is where you pass the arguments

	func(s string) {
		fmt.Println(s)
	}("An anonymous function with a argument for the parameter")
}

//func (r receiver) identifier (p parameter(s) (r return(s)) { code }
// anonymous function has no identifier or receiver
