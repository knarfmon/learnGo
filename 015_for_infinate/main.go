package main

import (
	"fmt"
)

var i int = 0

func main() {
	for {
		fmt.Println("Hello World")
		if i == 5 {
			break
		}
		i++
	}
}
