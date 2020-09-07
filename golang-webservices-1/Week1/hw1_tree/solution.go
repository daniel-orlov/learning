package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "path"
    "path/filepath"
)

func main() {
    tree := BuildTree(os.Args[1])
    fmt.Println(tree)
}

type File struct {
    Name string
}

type Folder struct {
    Name    string
    Files   []*File
    Folders map[string]*Folder
}

func (f *Folder) String() string {
    j, _ := json.Marshal(f)
    return string(j)
}

func BuildTree(dir string) *Folder {
    dir = path.Clean(dir)
    var tree *Folder
    var nodes = map[string]interface{}{}
    var walkFun filepath.WalkFunc = func(p string, info os.FileInfo, err error) error {
        if info.IsDir() {
            nodes[p] = &Folder{path.Base(p), []*File{}, map[string]*Folder{}}
        } else {
            nodes[p] = &File{path.Base(p)}
        }
        return nil
    }
    err := filepath.Walk(dir, walkFun)
    if err != nil {
        log.Fatal(err)
    }

    for key, value := range nodes {
        var parentFolder *Folder
        if key == dir {
            tree = value.(*Folder)
            continue
        } else {
            parentFolder = nodes[path.Dir(key)].(*Folder)
        }

        switch v := value.(type) {
        case *File:
            parentFolder.Files = append(parentFolder.Files, v)
        case *Folder:
            parentFolder.Folders[v.Name] = v
        }
    }

    return tree
}
