package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	//full declaration of a struct (some fields can be omitted)
	var acc Account = Account{
		Id:   7,
		Name: "James",
		Person: Person{
			Name:    "Bond",
			Address: "London",
		},
	}
	fmt.Printf("%#v\n", acc)

	//short declaration of a struct (all fileds must be used)
	acc.Owner = Person{1, "Sean Connery", "Edinburgh"}
	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Name)        //upper-level field gets priority
	fmt.Println(acc.Person.Name) //addressing the field directly

}
