package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	slc := make([]int, n, n)

	for i := range slc {
		fmt.Scan(&slc[i])
	}

	fmt.Println(slc[3])
}
