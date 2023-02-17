package main

import "fmt"

func main() {
	//simple condition
	boolVar := true
	if boolVar {
		fmt.Println("boolVar is true")
	}

	mapVal := map[string]string{"name": "lenovo"}
	//condition with init block
	if keyValue, keyExists := mapVal["name"]; keyExists { //first initialize, then execute condition "keyExists == true"
		fmt.Println("name =", keyValue)
	}

	//only checking if key exists
	if _, keyExists := mapVal["name"]; keyExists {
		fmt.Println("key exists")
	}

	someInt := 1
	//multiple if else
	if someInt == 1 {
		fmt.Println("someInt equals 1")
	} else if someInt == 0 {
		fmt.Println("someInt equals 0")
	} else {
		fmt.Println("someInt equals something unexpected")
	}

	//switch operator
	var hand string
	switch hand {
	case "left":
		fmt.Println("user is lefty")
	case "right":
		fmt.Println("user is righty")
	case "both", "ambidextrous":
		fmt.Println("user is alrighty")
	default:
		fmt.Println("hand is not chosen")
	}

	//using switch instead of multiple if-else statements
	var val1, val2 = 1, 2
	switch {
	case val1 > 1 || val2 < 10:
		fmt.Println("First case worked")
	case val2 > 9:
		fmt.Println("Second case worked")
	}

	//exiting the loop from switch
Loop:
	for key, val := range mapVal {
		fmt.Println("switch in loop", key, val)
		switch {
		case key == "lastname":
			break
			fmt.Println("this code is unreachable")
		case key == "name" && val == "lenovo":
			fmt.Println("switch will break loop here, we found the target")
			break Loop
		}
	} //here is from where the code execution will continue
}
