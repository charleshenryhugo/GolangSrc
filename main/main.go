package main

import (
	"fmt"
	structs "structpack"
)

func show(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	structs.GDTest()
}
