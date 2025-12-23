package main

import "fmt"

func main() {

	// func (receiver) identifier (parameters) (returns) {code...}
	// receivers give types their methods.
	// everything is passed by value
	foo()

	bar("-- one param")

	fmt.Println(aloha("Mr Tibbs"))
	s := aloha("Mr Tibbs")
	fmt.Println(s)

	s1, n := dogYears("Frank", 65)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("foo has no parameters and no return")
}

func bar(s string) {
	fmt.Println("bar one parameter and no return", s)
}
func aloha(s string) string {
	return fmt.Sprint("One param, one return, My name is ", s)
}
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}
