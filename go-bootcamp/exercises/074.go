package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Marshalling and Unmarshalling with json package
*/

func main() {
	//Creating a value of a type
	type Character struct {
		ID       int
		Nickname string
		Role     string
	}
	turkish := Character{
		ID:       747,
		Nickname: "Turkish",
		Role:     "boxing promoter",
	}
	//Marshalling the data
	b, err := json.Marshal(turkish)
	if err != nil {
		fmt.Println("we have a problem:", err)
	}
	//Output to the console
	os.Stdout.Write(b)
	fmt.Print("\n\n")
	//Creating a variable jsonBlob - a slice of bytes containing a raw string
	var jsonBlob = []byte(`[
	{"ID":47,"Nickname":"The Blade","Role":"Bullet Dodger"},
	{"ID":77,"Nickname":"Brick Top","Role":"Farmer"},
	{"ID":32,"Nickname":"Bullet Tooth","Role":"Negotiations Specialist"},
	{"ID":96,"Nickname":"Gipsy","Role":"Dags Lover"}
]`)
	var cast []Character
	err := json.Unmarshal(jsonBlob, &cast)
	if err != nil {
		fmt.Println("we have a problem:", err)
	}
	cast = append(cast, turkish)
	fmt.Printf("%+v", cast)
}
