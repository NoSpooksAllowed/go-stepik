package main

import (
	"fmt"
)

func main() {
	var n int
	var prew = 0
	var cur = 1
	var i = 1
	fmt.Scan(&n)

	for {
		prew, cur = cur, prew+cur

		i++

		if cur == n {
			fmt.Println(i)
			break
		}

		if cur > n {
			fmt.Println("-1")
			break
		}
	}
}
