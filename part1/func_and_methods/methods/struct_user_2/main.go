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

func (u *User) DoubleAge() int {
	return u.Age * 2
}

func (u *User) MyName() string {
	return u.Name
}

func (u *User) SexToInt() int {
	if u.Sex == true {
		return 1
	} else {
		return 0
	}
}
