package main

import (
	"fmt"
	"math"
)

func main() {
	var n, min, minAmount int

	fmt.Scan(&n)

	slice := make([]int, n)

	for i := range slice {
		fmt.Scan(&slice[i])

	}

	min = slice[0]
	for _, elem := range slice {

		min = int(math.Min(float64(elem), float64(min)))
	}

	for _, elem := range slice {
		if elem == min {
			minAmount++
		}
	}

	fmt.Println(minAmount)
}
