package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Работа восстановлена!")
		}
	}()

	panic("Error")
}
