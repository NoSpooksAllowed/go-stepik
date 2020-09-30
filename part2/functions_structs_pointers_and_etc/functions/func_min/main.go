package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var min int
	var n int

	fmt.Scan(&n)

	min = n
	for i := 0; i < 3; i++ {
		fmt.Scan(&n)

		if n < min {
			min = n
		}
	}

	return min
}
