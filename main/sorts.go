package main

import (
	"fmt"
	"sort"
)

func demo1(a int) {
	sli := randomInts(30)
	sort.Ints(sli)
	fmt.Println("Sorted array: ", sli)
	fmt.Printf("Find %d in array? index: %d", a, sort.SearchInts(sli, a))
}

func demo2(s string) {
	slstr := randomStrings(5)
	sort.Strings(slstr)
	fmt.Println("Sorted Array: ")
	for _, str := range slstr {
		fmt.Printf("%s\n------\n", str)
	}

	fmt.Printf("Find %s in array? index: %d", s, sort.SearchStrings(slstr, s))
}
