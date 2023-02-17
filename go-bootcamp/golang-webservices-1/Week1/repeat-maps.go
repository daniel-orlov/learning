package main

import "fmt"

func main() {
	//initiating using literal
	var user map[string]string = map[string]string{
		"name":     "James",
		"lastName": "Bond",
	}

	//initiating with defined capacity
	profile := make(map[string]string, 10)

	//get length
	mapLength := len(user)
	fmt.Printf("%d %#v\n", mapLength, profile)

	//access a value by key
	fmt.Println(user["name"])

	//if no such key, returns default value
	fmt.Println(user["motto"])

	//how to check if the key exists:
	motto, mottoExists := user["motto"]
	fmt.Println("motto =", motto, "key_exists =", mottoExists)
	//alternatively: just checking if the key exists
	_, keyExists := user["motto"]
	fmt.Println("key_exists =", keyExists)

	//deleting a key from map
	delete(user, "name")
	fmt.Printf("%#v\n", user)
}
