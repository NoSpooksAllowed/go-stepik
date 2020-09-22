package main

import (
	"fmt"
)

func main() {
	fmt.Println(Mul(2, 7, 3))
}

func Mul(x ...int) int {
	var res = 1

	for _, item := range x {
		res *= item
	}

	return res
}
