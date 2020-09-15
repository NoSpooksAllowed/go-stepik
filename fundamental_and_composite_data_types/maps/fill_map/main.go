package main

import (
	"fmt"
)

func main() {
	defCode := make(map[int]string)

	defCode[917] = "МТС-Поволжье"
	defCode[927] = "Мегафон-Поволжье"
	defCode[906] = "Билайн-Поволжье"

	fmt.Println(defCode)
}
