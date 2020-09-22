package main

import (
	"fmt"
)

func main() {
	fmt.Println(concat("Мне нравится ", "программировать!"))
}

func concat(s1 string, s2 string) (res string) {
	res = s1 + s2

	return
}
