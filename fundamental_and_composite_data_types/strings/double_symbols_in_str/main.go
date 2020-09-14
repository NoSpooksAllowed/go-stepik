package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)

	for _, ch := range s {
		fmt.Print(string(ch), string(ch))
	}

	fmt.Println()
}
