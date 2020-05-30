package main

import (
	"fmt"
	"sort"
)

/*
Custom sorting structs using pkg sort
*/

type Singer struct {
	First, Last string
	Band string
	Moustache bool
}

var singer0 = Singer{
	First:     "Paul",
	Last:      "McCartney",
	Band:      "The Beatles",
	Moustache: false,
}

var singer1 = Singer{
	First:     "Freddy",
	Last:      "Mercury",
	Band:      "Queen",
	Moustache: true,
}

var singer2 = Singer{
	First:     "Luciano",
	Last:      "Pavarotti",
	Band:      "3 Tenors",
	Moustache: true,
}

var singer3 = Singer{
	First:     "Joe",
	Last:      "Dassin",
	Band:      "",
	Moustache: false,
}

// ByBand implements sort.Interface for []Singer based on the Band field.
type ByBand []Singer

func (a ByBand) Len() int           { return len(a) }
func (a ByBand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBand) Less(i, j int) bool { return a[i].Band < a[j].Band }

// ByFirst implements sort.Interface for []Singer based on the First field.
type ByFirst []Singer

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

// ByLast implements sort.Interface for []Singer based on the Last field.
type ByLast []Singer

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }


func main()  {
	singers := []Singer{
		singer0,
		singer1,
		singer2,
		singer3,
	}

	fmt.Println("\n__SORTING__BY__BAND___NAME__")
	sort.Sort(ByBand(singers))
	fmt.Println(singers)

	fmt.Println("\n__SORTING__BY__FIRST__NAME__")
	sort.Sort(ByFirst(singers))
	fmt.Println(singers)

	fmt.Println("\n__SORTING__BY__LAST___NAME__")
	sort.Sort(ByLast(singers))
	fmt.Println(singers)

}