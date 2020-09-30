package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumInt(1, 2, 3, 4, 5))
}

func sumInt(n ...int) (int, int) {
	var sum int

	for _, elem := range n {
		sum += elem
	}

	return len(n), sum
}
