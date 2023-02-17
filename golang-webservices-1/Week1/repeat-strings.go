package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//empty string
	var str string
	fmt.Println(str)

	//to create string with escape chars use double quotes ""
	var escapeString string = "HA-HA\tHA-HA\n"
	fmt.Println(escapeString)

	//to create a string literal use back ticks``
	var literal string = `I almost forgot
		how to do this`
	fmt.Println(literal)

	//use single quotes for type 'byte' (alias for uint8)
	var rawBinary byte = '\x42' //I have no idea what symbol it's going to be
	fmt.Println(rawBinary)


	//use type 'rune' for single UTF-8 symbol (alias for uint32)
	var realRune rune = '⌘'
	fmt.Println(realRune)

	//string concatenation
	var classics = "Hello World!"
	addCheesyLine := classics + " We met again -_- "
	fmt.Println(addCheesyLine)

	//counting length of a string
	myFullName := "Daniel Orlov"
	lengthInBytes := len(myFullName)							//12 bytes
	lengthInSymbols := utf8.RuneCountInString(myFullName)		//12 runes
	fmt.Println(lengthInBytes, lengthInSymbols)

	//slicing the string
	name := myFullName[:6]
	surname := myFullName[7:]
	fmt.Println(name)
	fmt.Println(surname)

	//converting a string into a slice of bytes: back and forth
	byteString := []byte(myFullName)
	backIntoASring := string(byteString)
	fmt.Println(byteString, backIntoASring)


}
