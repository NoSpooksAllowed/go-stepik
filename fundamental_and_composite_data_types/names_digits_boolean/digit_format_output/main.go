package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Printf("%d %T\n%x\n%b\n", n, n, n, n)
}
