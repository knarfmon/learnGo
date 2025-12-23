package main

import "fmt"

func main() {

	f := incrementor() // increments f each time, same memory location
	fmt.Println(f())
	fmt.Println(&f) //same memory location
	fmt.Println(f())
	fmt.Println(&f) //same memory location
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int { //closing over the returning function
	x := 0
	return func() int { //this is the returning function
		x++ // this gets called over and over
		return x
	}
}

// incrementor is called, returns this:
//		func() int {
//		x++
//		return x
//	}
// assigns to f
// runs the function, adding one to x, returns x, prints
