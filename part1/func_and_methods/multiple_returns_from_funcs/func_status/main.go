package main

import (
	"fmt"
)

func main() {
	res, err := status("Alex", -15)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func status(s string, x int) (string, error) {
	if x < 0 {
		return "", fmt.Errorf("Возраст не может быть отрицательным!")
	}

	var res string

	switch {
	case x >= 0 && x <= 6:
		res = fmt.Sprintf("%s малыш", s)
	case x >= 7 && x <= 17:
		res = fmt.Sprintf("%s школьник", s)
	case x >= 18 && x <= 22:
		res = fmt.Sprintf("%s студент", s)
	case x >= 23 && x <= 55:
		res = fmt.Sprintf("%s взрослый", s)
	case x > 55:
		res = fmt.Sprintf("%s пожилой", s)
	}

	return res, nil
}
