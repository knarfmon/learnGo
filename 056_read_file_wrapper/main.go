package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("Error in main in reading file: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("there was an error in reading file: %s", err)
	}
	return xb, nil
}
