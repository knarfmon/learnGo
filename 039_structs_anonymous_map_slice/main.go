package main

import (
	"fmt"
)

func main() {

	f := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Frank",
		friends: map[string]int{
			"James":  32,
			"Casey":  34,
			"Maggie": 34,
		},
		favDrinks: []string{
			"Water",
			"Root Beer",
			"Cream Soda",
		},
	}

	fmt.Println(f.first)

	for k, _ := range f.friends {
		fmt.Println("--", k)
	}
	for _, v := range f.favDrinks {
		fmt.Println("----", v)
	}
}
