package main

func main() {
	//Creating files
	file1 := &File{"File1"}
	file2 := &File{"File2"}
	file3 := &File{"File3"}

	//Creating folders
	folder1 := &Folder{Name: "Folder1"}
	folder2 := &Folder{Name: "Folder2"}

	//Populating folders
	folder1.Add(file1)
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	//Where's Waldo?
	folder2.Search("Waldo")
}
