package main

import (
	"bufio"
	"os"
	"strconv"
)

var target int

func main() {
	//read from file
	fi, err := os.Open("input.txt")
	check(err)
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanWords)

	//prepare file to write to, creating a Writer
	fo, err := os.Create("output.txt")
	check(err)
	defer fo.Close()
	w := bufio.NewWriter(fo)

	//create a slice of string with all the data from the file
	var ss []string
	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}

	//convert each item of ss into int, create a map for constant time lookups
	mi := map[int]int{}
	for i, v := range ss {
		v, err := strconv.Atoi(v)
		check(err)
		//define target value
		if i == 0 {
			target = v
		}
		//populate the map only with relevant values
		if v < target {
			mi[v] = target - v

			//check for breaking conditions
			if _, found := mi[v]; found {
				if _, alsoFound := mi[mi[v]]; alsoFound {
					_, err = w.WriteRune(0x31) //writing 1 into a file
					check(err)
					_ = w.Flush()
					return
				}
			}
		}
	}
	//output in case of all lookups being unsuccessful
	_, err = w.WriteRune(0x30) //writing 0 into a file
	check(err)
	_ = w.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
