package main

import (
	"fmt"
	"strconv"
)

func main() {
	var digits [6]int
	var s string

	fmt.Scan(&s)

	for i := 0; i < len(digits); i++ {
		digits[i], _ = strconv.Atoi(string(s[i]))
	}

	sum1 := digits[0] + digits[1] + digits[2]
	sum2 := digits[3] + digits[4] + digits[5]

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
