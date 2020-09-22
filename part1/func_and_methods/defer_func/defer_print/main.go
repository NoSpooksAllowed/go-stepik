package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	defer PrintNum(n)
	PrintMessage()
}

func PrintNum(x int) {
	fmt.Print(x)
}

func PrintMessage() {
	fmt.Print("Ваше число = ")
}
