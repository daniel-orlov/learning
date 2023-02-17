package main

/*
Given the following code,
Unmarshal the JSON into a Go data structure
*/


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

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("\nWe have a problem:", err)
	}
	//fmt.Println(bs)
	fmt.Println(string(bs))

	// your code goes here
	type People []struct {
		First string `json:"First"`
		Age   int    `json:"Age"`
	}
	p := People{}
	err = json.Unmarshal(bs, &p)

	fmt.Println(p)
}