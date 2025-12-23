package main

import "fmt"

// Interface defines a set of method signatures, Hey baby if you got these methods then you are my type.
// Polymorphism is when a value can be of more than one type

type person struct {
	name string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agemt", sa.name)
}

// Interfaces allow us to have polymorphism
type human interface {
	speak() //If you have these methods then you are my type, so both secret agent and person are of this type
}

func saySomething(h human) {
	h.speak() //This function can take in any value of type human
}

func main() {
	// func (receiver) identifier (parameters) (returns) {code...}
	// receivers give types their methods.

	sa1 := secretAgent{
		person: person{
			name: "James",
		},
		ltk: true,
	}

	p2 := person{
		name: "Jenny",
	}
	//sa1.speak() //method speak attached to the type
	//p2.speak()
	saySomething(&sa1) //polymorphism, value can be of more than one type
	saySomething(&p2)
}
