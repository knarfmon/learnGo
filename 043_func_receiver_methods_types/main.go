package main

import "fmt"

type person struct {
	name string
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func main() {
	// func (receiver) identifier (parameters) (returns) {code...}
	// receivers give types their methods.

	p1 := person{
		name: "James",
	}
	p2 := person{
		name: "Jenny",
	}
	p1.speak() //method speak attached to the type
	p2.speak()
}
