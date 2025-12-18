package main

import "fmt"

func main() {

	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)
	fmt.Println("************ Inclusive - Exclusive *******************")

	//Includes 4 elements
	fmt.Println(" xi[0:4] Includes 4 elements", xi[0:4])

	//Includes 4 elements
	fmt.Println(" xi[:4] Includes 4 elements", xi[:4])

	//Starts at 4th, ends at 5th element
	fmt.Println("xi[4:5] Starts at 4th, ends at 5th element", xi[4:5])

	//Everything
	fmt.Println("xi[:] Everything", xi[:])

}
