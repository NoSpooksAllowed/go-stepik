package main

import (
	"fmt"
	"math"
)

func main() {
	mySlice := make([]int, 3)

	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = 1 * int(math.Pow(10, float64(i)))
	}

	fmt.Println(mySlice)
	fmt.Printf("%T\n", mySlice)
}
