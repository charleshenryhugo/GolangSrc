package main

import (
	"fmt"
)

func show(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	goroutineTest()
}
