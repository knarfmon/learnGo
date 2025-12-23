package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

type count int

// attach this function as a method to the type, using a receiver
func (b book) String() string {
	return fmt.Sprintf("The title of the book is %s", b.title)
}

func (c count) String() string {
	return fmt.Sprint("This is the number", strconv.Itoa(int(c)))
	//c is of type count, needs to be converted to int, then to a string.
}
func main() {
	b := book{title: "Go Programming Language"}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}
