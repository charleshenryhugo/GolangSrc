package main

import (
	"fmt"
	"structpack"
)

//"math"

func show(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	s := new(structpack.StackList)
	s.Init()
	s.Push(234)
	s.Push("haha")
	s.Pop()
	s.Push(false)
	s.Push(456)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
