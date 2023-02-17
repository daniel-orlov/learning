package main

import "fmt"

func main() {
	file1 := &File{Name: "my-env-vars.env"}
	file2 := &File{Name: "go.mod"}
	file3 := &File{Name: "go.sum"}

	folder1 := &Folder{
		Name:     "my-go-project",
		Children: []inode{file1},
	}

	folder2 := &Folder{
		Name:     "module",
		Children: []inode{folder1, file2, file3},
	}

	fmt.Printf("Folder '%v' hierarchy:\n", folder2.Name)
	folder2.represent("|___")

	clonedFolder := folder2.clone()
	fmt.Printf("Folder '%v' hierarchy:\n", folder2.Name+"_clone") //This line is a disappointment
	clonedFolder.represent("|___")

}
