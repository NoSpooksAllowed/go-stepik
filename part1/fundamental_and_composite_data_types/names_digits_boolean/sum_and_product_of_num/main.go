package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Printf("%d\n%d\n", a+b, a*b)
}
