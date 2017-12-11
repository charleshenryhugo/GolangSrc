package main

import (
	"fmt"
)

func intPointer() {
	i1 := 5
	fmt.Printf("An Integer: %d, it's location in memory: %p\n", i1, &i1)
	intP := &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

}

func strPointer() {
	s := "good bye"
	strP := &s
	*strP = "ciao"
	fmt.Printf("here is the pointer strP: %p\n", strP)
	fmt.Printf("here is the string *strP: %s\n", *strP)
	fmt.Printf("here is the string s: %s\n", s)
}

/*func WrongPointer() {
	const i int = 5
	ptr := &i   //error
	ptr2 := &10 //error

	var p *int = nil
	*p = 0 //runtime error
}*/
