package main

import "fmt"

type User struct {
	id          int
	name        string
	totalPost   int
	avgRatePost float64
}

func main() {
	user := User{1, "Max", 72, 0.87}
	rating := user.RateUser()
	fmt.Println(rating)
}

func (u *User) RateUser() float64 {
	return float64(u.totalPost) * u.avgRatePost
}
