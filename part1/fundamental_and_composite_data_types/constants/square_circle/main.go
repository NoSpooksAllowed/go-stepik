package main

import (
	"fmt"
	"math"
)

const (
	pi = 3.14159
)

func main() {
	var r float64

	fmt.Scan(&r)

	fmt.Println(pi * math.Pow(r, 2.0))
}
