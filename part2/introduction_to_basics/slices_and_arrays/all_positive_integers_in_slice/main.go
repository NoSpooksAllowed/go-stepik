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

	var sum int

	for _, elem := range slice {
		if elem > 0 {
			sum++
		}
	}

	fmt.Println(sum)
}
