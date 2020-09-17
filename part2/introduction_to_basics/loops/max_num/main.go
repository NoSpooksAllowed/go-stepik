package main

import (
	"fmt"
)

func main() {
	var max, n, sumMax int

	for {
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		if max < n {
			max = n
			sumMax = 0
		}

		if n == max {
			sumMax++
		}

	}

	fmt.Println(sumMax)
}
