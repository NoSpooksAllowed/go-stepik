package main

import (
	"fmt"
)

func main() {
	var s string
	var n int

	fmt.Scan(&s, &n)

	for i := 0; i < n; i++ {
		fmt.Print(s, " ")
	}

	fmt.Println()
}
