package main

import "fmt"

const (
	N       = 3
	DOUBLE1 = "=============================================================="
)

func main() {
	m := make(map[int]*int)
	fmt.Println("   VAR  |\tVALUE\t\tGO-SYNTAX\t\tTYPE")
	fmt.Println(DOUBLE1)
	for i := 0; i < N; i++ {
		m[i] = &i
		fmt.Printf("i (key) |\t%v\t\t%#v\t\t\t%T\n", i, i, i)
		fmt.Printf("m[i]=&i |\t%v\t%#v\t%T\n", &i, &i, &i)
		fmt.Printf("______________________Iteration #%v ended______________________\n", i)
	}

	fmt.Printf("\n\nThe whole map m now looks like this:\nDef. Repr - %v\nGo-sintax - %#v\n", m, m)

	for k, v := range m {
		fmt.Println("KEY:", k, "VAL:", *v)
	}

	fmt.Println(
		"\nAll values are pointing to the same memory address\n" +
			"which currently holds the last value 'i' had during the iterations while populating the map.",
	)
}
