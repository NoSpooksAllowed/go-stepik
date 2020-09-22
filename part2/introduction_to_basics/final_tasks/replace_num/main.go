package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n string

	fmt.Scan(&s)
	fmt.Scan(&n)

	fmt.Println(strings.Replace(s, n, "", -1))
}
