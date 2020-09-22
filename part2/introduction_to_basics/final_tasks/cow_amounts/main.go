package main

import "fmt"

func main() {
	var cow_amount int

	fmt.Scan(&cow_amount)

	if cow_amount%100 >= 11 && cow_amount%100 <= 19 {
		fmt.Printf("%d korov\n", cow_amount)
	} else {
		switch cow_amount % 10 {
		case 1:
			fmt.Printf("%d korova\n", cow_amount)
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fmt.Printf("%d korovy\n", cow_amount)
		default:
			fmt.Printf("%d korov\n", cow_amount)
		}
	}
}
