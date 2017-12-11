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
	h := new(structpack.HSon)
	h.PR()
}
