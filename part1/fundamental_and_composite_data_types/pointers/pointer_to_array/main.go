package main

import (
	"fmt"
)

func main() {
	var mainArray [70]int

	ptr := &mainArray

	fmt.Printf("%T\n", ptr)
}
