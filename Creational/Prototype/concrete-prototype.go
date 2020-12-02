package main

import "fmt"

type File struct {
	Name string
}

func (f *File) represent(indent string) {
	fmt.Printf("%v %v\n", indent, f.Name)
}

func (f *File) clone() inode {
	return &File{Name: f.Name + "_clone"}
}

type Folder struct {
	Name     string
	Children []inode
}

func (f *Folder) represent(indent string) {
	fmt.Printf("%v %v\n", indent, f.Name)
	for _, i := range f.Children {
		i.represent(indent + indent)
	}
}

func (f *Folder) clone() inode {
	clone := &Folder{Name: f.Name + "_clone"}
	var childNodes []inode
	for _, i := range f.Children {
		clonedChild := i.clone()
		childNodes = append(childNodes, clonedChild)
	}
	clone.Children = childNodes
	return clone
}
