package main

import (
	"encoding/json"
	"fmt"
)

type Vote struct {
	Id       int     `json:"id"`
	Username string  `json:"user_name"`
	AvgVote  float32 `json:"avg_vote"`
}

func main() {
	voters := []Vote{
		{Id: 12, Username: "Mikhail", AvgVote: 4.9},
		{Id: 2, Username: "Jul", AvgVote: 4.8},
		{Id: 17, Username: "Alex", AvgVote: 4.2},
	}

	data, _ := json.Marshal(voters)
	fmt.Printf("%s\n", data)
}
