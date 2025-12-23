package main

import "fmt"

func main() {

	// assigning a function to variable, passing as an argument to other function, returned as a value from functions, functions are a type

	x := func() {
		fmt.Println("Anonymous function")
	}

	y := func(s string) {
		fmt.Println(s)
	}

	x()
	y("An anonymous function called with a argument")

}

//func (r receiver) identifier (p parameter(s) (r return(s)) { code }
