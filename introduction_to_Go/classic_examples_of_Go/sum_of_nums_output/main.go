package main

import (
	"fmt"
)

func main() {
	var y int

	for x := 0; x <= 6; x++ {
		y += x
	}

	fmt.Println(y)
}
