package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		res := int(math.Pow(2, float64(i)))

		if res > n {
			break
		}

		fmt.Print(res, " ")
	}

	fmt.Println()
}
