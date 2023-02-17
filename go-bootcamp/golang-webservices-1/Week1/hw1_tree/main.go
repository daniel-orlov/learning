package main

import (
	"fmt"
	es "github.com/pkg/errors"
	"os"
	"sort"
	"strconv"
)

// Используй константы, так их проще в коде редактировать
const (
	MIDDLEBRANCH = "├───"
	ENDBRANCH    = "└───"
	FILEFLAG     = "-f"
	USAGE        = "usage go run main.go . [-f]"
)

// объявляй типы в начале файла, так читать удобнее
type Elem struct {
	Name, Path, Prefix string
	Size               int64
	IsDir              bool
	Parent             *Elem
	Children           []*Elem
}

func (e Elem) FormatSize() string {
	switch {
	case e.IsDir:
		return ""
	case e.Size > 0:
		return " (" + strconv.Itoa(int(e.Size)) + "b)"
	case e.Size == 0:
		return " (empty)"
	default:
		return ""
	}
}

func main() {
	out := os.Stdout //where the output will be stored -> printed out
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic(USAGE)
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == FILEFLAG
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

func Represent(elems ...*Elem) string {
	var repr string
	branch := MIDDLEBRANCH
	lastIndex := len(elems) - 1
	for index, l := range elems {
		//Checking branch type
		if index == lastIndex {
			branch = ENDBRANCH
		}
		//Adding prefixes to children
		for _, child := range l.Children {
			if index != lastIndex {
				child.Prefix += child.Parent.Prefix + "│\t"
				continue
			}
			child.Prefix += child.Parent.Prefix + "\t"
		}
		//Adding size
		size := l.FormatSize()
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

func StudyDirectory(parent *Elem, printFiles bool) (res []*Elem, err error) {
	//Open/close file
	parentDir, err := os.Open(parent.Path)
	if err != nil {
		err = es.Wrap(err, "failed to open parentDir")
		return nil, err
	}
	defer func() {
		//defer вызывается перед возвращением значения из функции, но после присвоения возвращаемых значений,
		//используя именованные возвращаемые значения можно такой трюк проворачивать
		err = parentDir.Close()
	}()

	//Os-independent separator
	separator := string(os.PathSeparator)

	//Reading from root dir and sorting the output
	names := make([]os.FileInfo, 0)
	names, err = parentDir.Readdir(0)
	if err != nil {
		err = es.Wrap(err, "failed to Readdir from parentDir")
		return nil, err
	}
	sort.Slice(names, func(i, j int) bool { return names[i].Name() < names[j].Name() })

	//Iterating over the dir contents and pouplating a slice of pointers to Elem
	children := make([]*Elem, 0, len(names))
	for _, file := range names {
		// Checking the flag
		if !file.IsDir() && !printFiles {
			continue
		}
		el := Elem{
			Name:   file.Name(),
			Path:   parent.Path + separator + file.Name(),
			Parent: parent,
			Size:   file.Size(),
			IsDir:  file.IsDir(),
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
