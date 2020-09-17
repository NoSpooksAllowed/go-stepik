package main

import (
	"fmt"
)

func main() {
	var n int
	var b int

	fmt.Scan(&n)

	for n > 0 {
		b = n % 10
		n /= 10
	}
	
	fmt.Println(b)
}
