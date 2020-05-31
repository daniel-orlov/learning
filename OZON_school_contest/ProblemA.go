package main

import (
	"bufio"
	"fmt"
	"os"
)

var target int

func main()  {
	//read from file
	fi, err := os.Open("input-201.txt")
	check(err)
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanWords)

	//prepare file to write to, creating a Writer
	fo, err := os.Create("input-201.a.txt")
	check(err)
	defer fo.Close()
	w := bufio.NewWriter(fo)

	//create a slice of string with all the data from the file
	var ss []string
	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}

	//create a map for constant time lookups
	mi := map[string]int{}

	//map form each unique number to its count in input
	for _, v := range ss {
		mi[v]++
	}
	fmt.Println(mi)
	for k,v := range mi{
		if v%2 != 0 {
			_, err = w.WriteString(k) //writing k into a file
			check(err)
			_ = w.Flush()
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}