package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	slice := make([]int, n)

	for i := range slice {
		fmt.Scan(&slice[i])
	}

	for i, elem := range slice {
		if i%2 == 0 {
			fmt.Print(elem, " ")
		}
	}

	fmt.Println()
}
