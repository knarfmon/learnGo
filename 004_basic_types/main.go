package main

import (
	"fmt"
)

const pi = 3.1415926

func main() {

	//default values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %v \n", i, f, b, s)

	//u := int(t)
	//fmt.Printf("%v %T\n", u, u)

	//Conversions
	var a int = 42
	var c float64 = float64(a)
	fmt.Printf("%v %T\n", c, c)

	t := 34.34
	fmt.Printf("%v %T\n", t, t)
	u := int(t)
	fmt.Printf("%v %T\n", u, u)

	fmt.Println("Happy ", pi, "Day")

}
