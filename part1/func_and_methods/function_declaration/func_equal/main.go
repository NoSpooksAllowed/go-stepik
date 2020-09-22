package main

import (
	"fmt"
)

func main() {
	fmt.Println(equal(2, 1))
}

func equal(x int, y int) bool {
	return x == y
}
