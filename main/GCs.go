package main

import (
	"fmt"
	"runtime"
)

//GCTest is a simple test for GarbageCollection
func GCTest() {
	strings := randomStrings(10000)
	fmt.Println(strings)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
