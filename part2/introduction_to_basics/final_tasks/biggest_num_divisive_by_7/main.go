package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	for ; b >= a; b-- {
		if b%7 == 0 {
			fmt.Println(b)
			return
		}
	}

	fmt.Println("NO")
}
