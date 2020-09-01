package main

import "fmt"

type Person struct {
	Id int
	Name string
}

//this method will not change the original struct (method takes a copy of the struct)
//basically useless
func (p Person) UpdateName(name string) {
	p.Name = name
}

//this is a classic setter, which changes the struct (method takes the address)
func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id int
	Name string
	Person
}

type MySlice []int

func (sl *MySlice) Add(val ...int) {
	*sl = append(*sl, val...)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}


func main() {
	pers := Person{7, "James"}
	pers.SetName("James Bond")
	// (&pers).SetName("James, James Bond") <- not like this!
	// But Go sees it that way
	fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id: 1,
		Name: "The First",
		Person: Person{
			Id: 2,
			Name: "Human",
		},
	}
	fmt.Printf("Print: %#v\n", acc)
	acc.SetName("Adam")
	fmt.Printf("Print after upd8:\n %#v\n", acc)
	acc.Person.SetName("Noribman")
	fmt.Printf("Print after upd8 w/ embd struct's method call:\n %#v\n", acc)


	// var months MySlice = MySlice{1,2,3,4,5,6,7,8,9}
	//alternative version
	months := MySlice([]int{1,2,3,4,5,6,7,8,9})
	fmt.Println(months)
	fmt.Println("adding 3 last months")
	months.Add(10,11,12)
	fmt.Println("Full calender now:", months)
	fmt.Println("The total length is:", months.Count())
}
