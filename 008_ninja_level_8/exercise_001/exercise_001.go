package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	users := []user{u1}

	users = append(users, user{
		First: "Andrew",
		Age:   31,
	}, user{
		First: "Bob",
		Age:   22,
	})
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
