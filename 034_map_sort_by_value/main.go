package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Key   string
	Value int
}

func main() {
	m := map[string]int{
		"banana": 3,
		"apple":  5,
		"cherry": 2,
	}

	items := make([]Item, 0, len(m))
	for k, v := range m {
		items = append(items, Item{k, v}) //making a slice of item structs
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})

	for _, item := range items {
		fmt.Println(item.Key, item.Value)
	}
}
