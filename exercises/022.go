package main

//Printing out the years I lived without coding on Go

import "fmt"

//First solution

//func main() {
//	bd := 1991
//	for bd < 2020 {
//		fmt.Println(bd)
//		bd++
//	}
//}

//Second solution

//func main() {
//	for bd := 1991; bd < 2020; bd++ {
//		fmt.Println(bd)
//	}
//}

//Third solution

func main() {
	bd := 1991
	for {
		if bd == 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}