package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int `json:"user_id,string"` //this and below are struct tags
	Username string
	Address  string `json:",omitempty"`
	Company string `json:"-"`
}

func main() {
	u := &User{
		ID:       47,
		Username: "Dan",
		Address:  "",
		Company: "Andersen",
	}
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string: %s\n", string(result))
}
