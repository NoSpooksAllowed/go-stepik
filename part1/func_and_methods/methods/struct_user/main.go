package main

import (
	"fmt"
)

type User struct {
	Name string
	Sex  bool
	Age  int
}

func main() {
	p := User{}
	p.SetAge(10)
	p.SetName("Alex")
	p.SetSex(true)
	fmt.Println(p)
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) SetName(nm string) {
	u.Name = nm
}

func (u *User) SetSex(male bool) {
	u.Sex = male
}
