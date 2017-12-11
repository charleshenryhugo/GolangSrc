package main

import (
	"fmt"
	"io"
	"log"
)

func trace(s string) string {
	fmt.Println("Entering: ", s)
	return s
}

func untrace(s string) {
	fmt.Println("Leaving: ", s)
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}

func c() {
	defer func() {
		untrace(trace("c"))
	}()
	fmt.Println("in c")
	a()
}

func deferTracing() {
	c()
	//b()
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func deferLogvalues() {
	func1("GO")
}
