package main

import "fmt"

func main() {

	type person struct {
		first    string
		last     string
		iceCream []string
	}
	p1 := person{
		first:    "Frank",
		last:     "Young",
		iceCream: []string{"vanilla", "chocolate", "strawberry"},
	}
	p2 := person{
		first:    "Casey",
		last:     "Young",
		iceCream: []string{"Neapolitan", "chocolate", "strawberry"},
	}
	xp := []person{p1, p2}
	for _, v := range xp {
		fmt.Println(v.first, v.last, "Favorite Ice Creams")
		for _, p := range v.iceCream {
			fmt.Println("- ", p)
		}
	}
	fmt.Println("*************** structs in a map ********************************")

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last)
		for _, p := range v.iceCream {
			fmt.Println("- ", p)
		}
	}

}
