package main

import "fmt"

func main() {
	fmt.Println("********* Make a Map, String is the Key *******************")
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("The number two is", m["two"])
	fmt.Println("********* Make a Map, Int is the Key *******************")
	m2 := make(map[int]string)
	m2[1] = "one"
	m2[2] = "two"
	m2[3] = "three"
	fmt.Println(m2[3])
	fmt.Println("************ Length of Map ****************")
	fmt.Println(len(m2))
	fmt.Println("************ Adding to the Map ****************")
	m2[5] = "Frank"
	fmt.Println(m2)
	fmt.Println("************ Ranging Over the Map ****************")
	c := 0
	for k, v := range m2 {
		c++
		fmt.Println("Iteration is ", c, " Key is ", k, " Value is", v)
	}
	fmt.Println("************ Delete an Element from the Map ****************")
	delete(m2, 3)
	fmt.Println(m2)
	fmt.Println("************ Comma Ok on the Map ****************")
	v, ok := m["one"]
	if ok {
		fmt.Println("Found the key, value is", v)
	} else {
		fmt.Println("Key Not found")
	}
	fmt.Println("**** More common version of Comma Ok on the Map ****************")
	if v, ok := m["one"]; ok {
		fmt.Println("Found the key, value is", v)
	} else {
		fmt.Println("Key Not found")
	}
	fmt.Println("****Even More common version of Comma Ok on the Map ****************")
	if v, ok := m["one"]; !ok {
		fmt.Println("Key Not found")
	} else {
		fmt.Println("Using v here", v)
	}
}
