package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums_units = map[int64]string{
		1: "Один",
		2: "Два",
		3: "Три",
		4: "Четыре",
		5: "Пять",
		6: "Шесть",
		7: "Семь",
		8: "Восемь",
		9: "Девять",
	}

	var nums_tens = map[int64]string{
		1: "Десять",
		2: "Двадцать",
		3: "Тридцать",
		4: "Сорок",
		5: "Пятьдесят",
		6: "Шестьдесят",
		7: "Семьдесят",
		8: "Восемьдесят",
		9: "Девяносто",
	}

	var input string

	fmt.Scan(&input)

	tens, _ := strconv.ParseInt(string(input[0]), 0, 32)
	unit, _ := strconv.ParseInt(string(input[1]), 0, 32)

	fmt.Println(nums_tens[tens], nums_units[unit])
}
