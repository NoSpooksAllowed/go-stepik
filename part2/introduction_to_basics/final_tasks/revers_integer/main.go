package main

import (
	"fmt"
)

func main() {
	var n, reverseN int

	fmt.Scan(&n)

	for n > 0 {
		remainder := n % 10
		reverseN = reverseN*10 + remainder
		n /= 10
	}

	fmt.Println(reverseN)
}
