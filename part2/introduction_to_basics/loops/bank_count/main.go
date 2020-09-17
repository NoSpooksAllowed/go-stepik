package main

import (
	"fmt"
)

func main() {
	var x, y int
	var p float64
	var years int

	fmt.Scan(&x, &p, &y)
	p /= 100

	for ; x <= y; x += int(float64(x) * p) {
		years++
	}

	fmt.Println(years)
}
