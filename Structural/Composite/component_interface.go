package main

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(query string) {
	fmt.Printf("Searching for '%v' in file '%s'\n", query, f.Name)
}

func (f *File) getName() string {
	return f.Name
}
