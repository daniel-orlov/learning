package main

import "fmt"

type Folder struct {
	Name     string
	Contents []component
}

func (f *Folder) Search(query string) {
	fmt.Printf("Searching recursively for '%s' in folder '%s'\n", query, f.Name)
	for _, composite := range f.Contents {
		composite.Search(query)
	}
}

func (f *Folder) Add(c component) {
	f.Contents = append(f.Contents, c)
}
