package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	fmt.Scan(&name, &age)

	fmt.Printf("Меня зовут %s и мне %d лет\n", name, age)
}
