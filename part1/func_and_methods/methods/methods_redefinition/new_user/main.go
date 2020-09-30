package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
	rate float64
}

type NewUser struct {
	User
}

func main() {
	usr := NewUser{}
	usr.UserMethod()
}

func (u *User) UserMethod() {
	fmt.Println("Метод определен")
}

func (n *NewUser) UserMethod() {
	fmt.Println("Метод переопределен")
}
