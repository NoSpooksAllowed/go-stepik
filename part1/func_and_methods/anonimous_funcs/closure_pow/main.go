package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	myPow := pow
	fmt.Print(myPow()(i))
}

func pow() func(x int) int {
	return func(x int) int {
		return x * x
	}
}
