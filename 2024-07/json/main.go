package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json: first_name`
	LastName  string `json: last_name`
	Age       int    `json: age`
}

func main() {
	jio := User{FirstName: "George William", LastName: "Sison", Age: 23}

	jioJson, err := json.Marshal(jio)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jioJson))
}
