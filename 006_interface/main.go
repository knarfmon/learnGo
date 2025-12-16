package main

import (
	"fmt"
)

// Hey Babby if you got these methods then you are my type.

type person struct { // Anything of type person will be of type human
	first string
}
type secretagent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (p person) speak() { // has the speak method so you are of human type
	fmt.Println("Hello, I am, ", p.first)
	//The speak() method is attached to Values of type person, package scope.
	// (p person) is called a receiver
}

func (p secretagent) speak() { // has the speak method so you are of human type
	fmt.Println("Hello, I am secret agent, ", p.first)
}

func talk(h human) { //takes in type human
	h.speak()
}

func main() {
	p := person{
		first: "James",
	}
	talk(&p)

	sa := secretagent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	talk(&sa)
}
