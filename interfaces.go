package main

import "fmt"

const (
	DOUBLE = "====================================================================================================="
	SINGLE = "-----------------------------------------------------------------------------------------------------"
)

type Struct struct{}

func (strukt Struct) Function() {}

type Interface interface {
	Function()
}

func InitType() Struct {
	var s Struct
	return s
}
func InitPointer() *Struct {
	var s *Struct
	return s
}
func InitEfaceType() interface{} {
	var s Struct
	return s
}
func InitEfacePointer() interface{} {
	var s *Struct
	return s
}
func InitIfaceType() Interface {
	var s Struct
	return s
}
func InitIfacePointer() Interface {
	var s *Struct
	return s
}

func main() {

	fmt.Println("        \tVALUE\t\tGO-SYNTAX REPRESENTATION\t\tTYPE\t\t\t==nil")
	fmt.Println(DOUBLE)
	s := InitType()
	fmt.Printf("    Struct\t%v\t\t%#v\t\t\t\t%T\t\t%v\n", s, s, s, "ERROR")
	fmt.Printf(" ^pointer^\t%v\t\t%#v\t\t\t\t%T\t\t%v\n", &s, &s, &s, &s == nil)
	fmt.Println(SINGLE)
	ps := InitPointer() //ps != &s
	fmt.Printf("   *Struct\t%v\t\t%#v\t\t\t%T\t\t%v\n", ps, ps, ps, ps == nil)
	fmt.Printf(" ^pointer^\t%v\t%#v\t\t%T\t\t%v\n", &ps, &ps, &ps, &ps == nil)
	fmt.Println(SINGLE)
	e := InitEfaceType()
	fmt.Printf(" (S)inface\t%v\t\t%#v\t\t\t\t%T\t\t%v\n", e, e, e, e == nil)
	fmt.Printf(" ^pointer^\t%v\t%#v\t\t%T\t\t%v\n", &e, &e, &e, &e == nil)
	fmt.Println(SINGLE)
	ep := InitEfacePointer()
	fmt.Printf("(*S)inface\t%v\t\t%#v\t\t\t%T\t\t%v\n", ep, ep, ep, ep == nil)
	fmt.Printf(" ^pointer^\t%v\t%#v\t\t%T\t\t%v\n", &ep, &ep, &ep, &ep == nil)
	fmt.Println(SINGLE)
	i := InitIfaceType()
	fmt.Printf(" (S)Interf\t%v\t\t%#v\t\t\t\t%T\t\t%v\n", i, i, i, i == nil)
	fmt.Printf(" ^pointer^\t%v\t%#v\t\t%T\t\t%v\n", &i, &i, &i, &i == nil)
	fmt.Println(SINGLE)
	ip := InitIfacePointer()
	fmt.Printf("(*S)Interf\t%v\t\t%#v\t\t\t%T\t\t%v\n", ip, ip, ip, ip == nil)
	fmt.Printf(" ^pointer^\t%v\t%#v\t\t%T\t\t%v\n", &ip, &ip, &ip, &ip == nil)
}
