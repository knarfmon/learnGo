package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range m {
		fmt.Println(k)
		for count, vv := range v {
			fmt.Println(count, "---", vv)
			// When ranging over map the k represents the key provided, when ranging over a slice, the key value is the index position.
		}

	}
	fmt.Println("-------------- Adding to the Map-----------------------------")
	m["fleming-ian"] = []string{"steaks", "cigars", "espionage"}
	for k1, v1 := range m {
		fmt.Println(k1)
		for count, vv1 := range v1 {
			fmt.Println(count, "---", vv1)
			// When ranging over map the k represents the key provided, when ranging over a slice, the key value is the index position.
		}
	}
	fmt.Println("-------------- Deleting from the Map-----------------------------")
	delete(m, "no_dr")
	for k1, v1 := range m {
		fmt.Println(k1)
		for count, vv1 := range v1 {
			fmt.Println(count, "---", vv1)
			// When ranging over map the k represents the key provided, when ranging over a slice, the key value is the index position.
		}
	}
}
