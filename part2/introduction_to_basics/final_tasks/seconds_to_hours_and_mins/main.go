package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Scan(&k)

	k -= k % 60

	k /= 60

	fmt.Printf("It is %d hours %d minutes\n", k/60, k%60)
}
