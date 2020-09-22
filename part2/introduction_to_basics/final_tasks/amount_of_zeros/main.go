package main

import (
	"fmt"
)

func main() {
	var n, sum, d int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&d)

		if d == 0 {
			sum++
		}
	}

	fmt.Println(sum)
}
