package main

//Using anonymous struct

import "fmt"

//Previously used way - creating a type person
//type person struct{
//	first, last string
//	is_artist bool
//}

func main() {
	//Using anonymous struct
	p1 := struct{
		first, last string
		is_artist bool
	}{
		first: "Frank",
		last: "Sinatra",
		is_artist: true,
	}
	fmt.Println(p1)
}