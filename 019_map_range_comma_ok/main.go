package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println(" Age of James is :", m["James"])
	if v, ok := m["James"]; ok {
		fmt.Println("James is :", v)
	}

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println(" Age of Q is :", m["Q"])
	age := m["Q"]
	fmt.Println(age)

	fmt.Println("")
	fmt.Println("************ , ok ******************")
	if v, ok := m["Q"]; !ok {
		fmt.Println("Q is nil!, here is the zero value of int", v)
	}

}
