package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, _ :=	ioutil.ReadDir("testdata")
	printDirContent(files...)
	// fmt.Printf("Type = %T\n", files)
	//iterarting over and printing

}

func printDirContent(dir ...os.FileInfo) {
	wd, _ := os.Getwd()
	currentDir := wd + "\\testdata\\"

	for _, file := range dir {
		if file.IsDir() {
			dirName := file.Name()

			childDir := wd  + dirName
			err := os.Chdir(childDir)
			if err != nil {
				panic(err)
			}
			fs, _ := ioutil.ReadDir(dirName)
			printDirContent(fs...)
		} else {
			fmt.Printf("├───%v (%vb)\n", file.Name(), file.Size())
		}
	}
}
