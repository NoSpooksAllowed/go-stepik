package main

import (
	"fmt"
	"math"
)

func main() {
	const base = 2

	for i := 1; i <= 6; i++ {
		fmt.Println(int(math.Pow(float64(base), float64(i))))
	}
}
