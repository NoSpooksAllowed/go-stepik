package main

import (
	"fmt"
)

func main() {
	var d, s int

	fmt.Scan(&d, &s)

	if s > d {
		fmt.Println("S больше D")
	} else if s < d {
		fmt.Println("D больше S")
	} else {
		fmt.Println("Оба равны")
	}
}
