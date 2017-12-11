package main

import (
	"fmt"
)

//recursive fibonacci is a very slow method!
func fibonacci(a int) (pos int, val int) {
	if a == 0 {
		return 0, 1
	} else if a == 1 {
		return 1, 1
	}
	_, v1 := fibonacci(a - 1)
	_, v2 := fibonacci(a - 2)
	return a, v1 + v2
}

func fromTenToOne(a int) {
	if a < 1 {
		return
	}
	fmt.Println(a)
	a--
	fromTenToOne(a)
}

func factorial(a int) (res int) {
	if a <= 1 {
		return 1
	}
	return factorial(a-1) * a
}
