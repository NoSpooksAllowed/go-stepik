package main

import (
	"fmt"
	"strings"
)

func main() {
	var numsStr1, numsStr2 string

	fmt.Scan(&numsStr1, &numsStr2)

	for _, ch := range numsStr1 {
		if strings.ContainsRune(numsStr2, rune(ch)) {
			fmt.Print(string(ch), " ")
		}
	}

	fmt.Println()
}
