package main

import (
	"fmt"
	"runtime"
)

func GCTest() {
	strings := randomStrings(10000)
	fmt.Println(strings)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
