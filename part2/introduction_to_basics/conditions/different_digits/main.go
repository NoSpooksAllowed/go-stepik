package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	a := n % 10
	b := (n / 10) % 10
	c := (n / 100) % 10

	if a == b || a == c || b == c {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
