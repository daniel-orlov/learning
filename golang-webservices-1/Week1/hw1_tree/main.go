package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	es "github.com/pkg/errors"
)

func main() {
	out := os.Stdout //where the output will be stored -> printed out
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f" //boolean flag
	err := dirTree(out, path, printFiles)
	if err != nil {
		fmt.Println(err)
	}
}

func dirTree(out *os.File, path string, printFiles bool) error {
	rootElem := Elem{Path: path, Name: path}
	contents, err := StudyDirectory(&rootElem, printFiles)
	if err != nil {
		err = es.Wrap(err, "failed to StudyDirectory from given path")
		return err
	}
	_, err = fmt.Fprint(out, Represent(contents...))
	if err != nil {
		err = es.Wrap(err, "failed to Represent to out")
		return err
	}
	return nil
}

type Elem struct {
	Name, Path, Prefix string
	Size int64
	IsDir bool
	Parent *Elem
	Children []*Elem
}

func Represent(elems ...*Elem) string {
	var repr string
	branch := "├───"
	var size string
	for index, l := range elems {
		//Checking branch type
		if index == len(elems)-1 {
			branch = "└───"
		}
		//Adding prefixes to children
		if len(l.Children) > 0 && index != len(elems)-1 {
			for _, child := range l.Children {
				child.Prefix += "│\t" + child.Parent.Prefix
			}
		} else if len(l.Children) > 0{
			for _, child := range l.Children {
				child.Prefix += child.Parent.Prefix + "\t"
			}
		}
		//Adding size
		if !l.IsDir && l.Size > 0 {
			size = " (" + strconv.Itoa(int(l.Size)) + "b)"
		} else if !l.IsDir && l.Size == 0{
			size = " (empty)"
		} else if l.IsDir {
			size = ""
		}
		//Representing
		repr += l.Prefix + branch + l.Name + size + "\n"
		//Calling Represent on children
		if len(l.Children) > 0 {
			chldrn := Represent(l.Children...)
			repr += chldrn
		}
	}
	return repr
}

func StudyDirectory(parent *Elem, printFiles bool) ([]*Elem, error) {
	//Open/close file
	parentDir, err := os.Open(parent.Path)
	if err != nil {
		err = es.Wrap(err, "failed to open parentDir")
		return nil, err
	}
	defer parentDir.Close()

	//Os-independent separator
	separator := string(os.PathSeparator)

	//Reading from root dir and sorting the output
	names := make([]os.FileInfo, 0, 15)
	names, err = parentDir.Readdir(0)
	if err != nil {
		err = es.Wrap(err, "failed to Readdir from parentDir")
		return nil, err
	}
	sort.Slice(names, func(i, j int) bool { return names[i].Name() < names[j].Name() })

	//Iterating over the dir contents and pouplating a slice of pointers to Elem
	children := make([]*Elem, 0, 25)
	for _, file := range names{
		// Checking the flag
		if !file.IsDir() && !printFiles {
			continue
		}
		el := Elem{
			Name: file.Name(),
			Path: parent.Path + separator + file.Name(),
			Parent: parent,
			Size: file.Size(),
			IsDir: file.IsDir(),
		}
		//Studying underlying directories encountered
		if el.IsDir {
			el.Children, err = StudyDirectory(&el, printFiles)
			if err != nil {
				err = es.Wrap(err, "failed StudyDirectory recursion")
				return nil, err
			}
		}
		children = append(children, &el)
	}
	return children, nil
}
