package main

import (
	"fmt"
	structs "structpack"
)

//"math"

func show(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	structs.TypeAssertTest()
}
