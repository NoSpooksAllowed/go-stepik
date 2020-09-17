package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		fmt.Scan(&n)

		switch {
		case n > 100:
			return
		case n < 10:
			continue
		default:
			fmt.Println(n)
		}
	}
}
