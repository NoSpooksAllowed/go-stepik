package main

import (
	"fmt"
)

func main() {
	x, y := ptrVar(1, 2)
	fmt.Printf("%T %T\n", x, y)
}

func ptrVar(x int, y int) (*int, *int) {
	return &x, &y
}
