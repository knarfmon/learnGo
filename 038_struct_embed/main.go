package main

import "fmt"

func main() {

	type engine struct {
		electric bool
	}
	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Chevrolet",
		model: "Corvette",
		doors: 2,
		color: "red",
	}
	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Honda",
		model: "CRV",
		doors: 4,
		color: "Black",
	}
	fmt.Println(v1, v2)
	fmt.Println(v1.make, v2.make)
}
