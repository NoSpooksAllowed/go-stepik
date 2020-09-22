package main

import (
	"fmt"
)

func main() {
	var n float64

	fmt.Scan(&n)

	switch {
	case n <= 0:
		fmt.Printf("число %.2f не подходит\n", n)
	case n > 10000:
		fmt.Printf("%e\n", n)
	default:
		fmt.Printf("%.4f\n", n*n)
	}
}
