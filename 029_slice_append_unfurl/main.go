package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)

	for k, v := range x {
		fmt.Println("k:", k, "v:", v)
	}
	fmt.Println("**********************")

	x = append(x, 53, 54, 55)
	for k, v := range x {
		fmt.Println("k:", k, "v:", v)
	}
	fmt.Println("********** Using Range to Append *************")
	y := []int{56, 57, 58, 59, 60}

	for _, v := range y {
		x = append(x, v)
	}

	for k, v := range x {
		fmt.Println("k:", k, "v:", v)
	}
	fmt.Println("****** Unfurled Example *****************")
	x = append(x, y...)
	for k, v := range x {
		fmt.Println("k:", k, "v:", v)
	}
}
