package main

import (
	"fmt"
	es "github.com/pkg/errors"
	"os"
	"path/filepath"
)

/*TODO
- implement proper formating (levels and final elements)
- make dirTree write to *os.File, not just Print
*/

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
	if printFiles {
		//all files version
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				details := fmt.Sprintf("failed accessing a path %q\n", path)
				err = es.Wrap(err, details)
				return err
			}
			if info.IsDir() {
				line := fmt.Sprintf("├─── dir: %v\n", info.Name())
				_, err = fmt.Fprint(out, line)
				if err != nil {
					err = es.Wrap(err, "failed to Fprint to out")
					return err
				}
			} else {
				line := fmt.Sprintf("├─── file: %v (%vb)\n", info.Name(), info.Size())
				_, err := fmt.Fprint(out, line)
				if err != nil {
					err = es.Wrap(err, "failed to Fprint to out")
					return err
				}
			}
			return err
		})
		return err
	} else {
		//only dirs version
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				details := fmt.Sprintf("failed accessing a path %q\n", path)
				err = es.Wrap(err, details)
				return err
			}
			if info.IsDir() {
				line := fmt.Sprintf("├─── dir: %v\n", info.Name())
				_, err = fmt.Fprint(out, line)
				if err != nil {
					err = es.Wrap(err, "failed to Fprint to out")
					return err
				}
			}
			return err
		})
		return err
	}
}
