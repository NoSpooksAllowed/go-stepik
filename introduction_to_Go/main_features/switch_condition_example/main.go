package main

import (
	"fmt"
)

func main() {
	var month int

	fmt.Scan(&month)

	switch month {
	case 1:
		fmt.Println("Январь")
	case 2:
		fmt.Println("Февраль")
	case 3:
		fmt.Println("Март")
	case 4:
		fmt.Println("Апрель")
	case 5:
		fmt.Println("Май")
	case 6:
		fmt.Println("Июнь")
	case 7:
		fmt.Println("Июль")
	case 8:
		fmt.Println("Август")
	case 9:
		fmt.Println("Сентябрь")
	case 10:
		fmt.Println("Октябрь")
	case 11:
		fmt.Println("Ноябрь")
	case 12:
		fmt.Println("Декабрь")
	default:
		fmt.Println("Номера такого месяца не существует, введи число в диапазоне от [1:12]")
	}
}
