package main

import (
	"fmt"
)

const (
	degree_per_min = 0.5
)

func main() {
	var degrees int

	fmt.Scan(&degrees)

	minutes := int(float64(degrees) / degree_per_min)

	fmt.Printf("It is %d hours %d minutes\n", minutes/60, minutes%60)
}
