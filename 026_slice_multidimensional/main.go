package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini, Chocolate"}
	jp := []string{"Jenny", "Moneypenny", "Guiness", "Vanilla"}

	xxs := [][]string{jb, jp}
	fmt.Println(xxs)
	fmt.Println(xxs[0])
	fmt.Println(xxs[1])
	fmt.Println(xxs[1][0])
	fmt.Println("--------------------------------")

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Im 008"}

	xx := [][]string{x1, x2}
	for _, v := range xx {
		fmt.Println(v)
		for _, b := range v {
			fmt.Println(b)
			
		}
	}

}
