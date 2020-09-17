package main

import (
	"fmt"
)

func main() {
	var l, n, sum int

	fmt.Scan(&l)

	for i := 0; i < l; i++ {
		fmt.Scan(&n)

		if (n%8 == 0) && (n < 100 && n >= 10) {
			sum += n
		}
	}

	fmt.Println(sum)
}
