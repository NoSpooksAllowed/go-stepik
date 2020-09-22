package main

import (
	"fmt"
)

func main() {
	workArray := [10]uint8{}

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	var k, j int

	for i := 0; i < 3; i++ {
		fmt.Scan(&k, &j)

		workArray[k], workArray[j] = workArray[j], workArray[k]
	}

	for i := range workArray {
		fmt.Print(workArray[i], " ")
	}

	fmt.Println()
}
