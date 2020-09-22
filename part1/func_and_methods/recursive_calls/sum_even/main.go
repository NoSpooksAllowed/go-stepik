package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumEven(12))
}

func sumEven(x int) int {
	if x == 2 {
		return 2
	}

	return x + sumEven(x-2)
}
