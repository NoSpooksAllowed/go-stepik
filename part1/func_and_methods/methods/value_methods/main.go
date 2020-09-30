package main

import "fmt"

type Room struct {
	Num     int
	isClean bool
}

func (r Room) CleanRoom(num int) bool {
	return true
}

func main() {
	r := Room{}
	clean := func() func(r Room, num int) bool {
		return func(r Room, num int) bool {
			return true
		}
	}()
	if clean(r, 3) == true {
		fmt.Print("Номер убран")
	}
}
