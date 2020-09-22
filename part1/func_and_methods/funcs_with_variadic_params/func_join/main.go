package main

import (
	"fmt"
)

func main() {
	fmt.Println(Join("Hello ", "world!"))
}

func Join(s ...string) string {
	var res string

	for _, item := range s {
		res += item
	}

	return res
}
