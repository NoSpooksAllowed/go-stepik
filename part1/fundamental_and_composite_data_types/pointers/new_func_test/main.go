package main

import (
	"fmt"
)

func main() {
	ptr := new(int)
	*ptr = 56743

	fmt.Printf("%T %[1]d\n", *ptr)
}
