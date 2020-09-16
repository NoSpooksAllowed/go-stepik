package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"zarplata"`
	Worked bool   `json:"iswork,omitempty"`
}

func main() {
	var workers = []User{
		{Name: "Александр", Salary: 22500, Worked: true},
		{Name: "Маргарита", Salary: 32200, Worked: true},
		{Name: "Инесса", Salary: 0, Worked: false},
	}

	data, _ := json.Marshal(workers)

	fmt.Printf("%s\n", data)
}
