package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)

	for i, ch := range s {
		if i%2 == 0 {
			fmt.Print(string(ch))
		}
	}

	fmt.Println()
}
