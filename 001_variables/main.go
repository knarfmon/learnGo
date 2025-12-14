package main

import "fmt"

// Package level scope
func main() {
	//foo()

	// a Variable hold a Value of a specific Type

	var a int = 5
	println(a)
	fmt.Printf("%v of type %T \n", a, a)
	a = 55
	println(a)

	q := 77.7
	fmt.Printf("%v of type %T \n", q, q)

	b := "Frank is great."
	//Use the short declaration operator
	println(b)

	c, d, e := 3, 4, 5
	println(c, d, e)
	fmt.Printf("%v of type %T \n", c, c)
	h := float64(c) //Conversion
	fmt.Printf("%v of type %T \n", h, h)

	var g int
	println(g) // use when you want a zero value

	g = 1000
	println(g)
	fmt.Printf("Printing the value of g in Binary, %b \n", g)
}

//func foo() {
//	fmt.Println("foo says go run main.go")
//}
