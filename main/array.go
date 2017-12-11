package main

import (
	"fmt"
	"time"
)

type (
	//Vectorf3D is 3-dim array of float64
	Vectorf3D [3]float64
)

const (
	//ARRMAX is the max-size of an array
	ARRMAX = 5001
)

var fib [ARRMAX]int

func sliceInitials() {
	arr := [5]int{12, 34, 56, 78, 90}
	sl1 := arr[1:5]
	sl1[2] -= 5
	sl2 := arr[0:4]
	sl3 := arr[:]
	arrP := &arr

	sl5 := []int{1, 2, 3}

	//sl4 and sl6 are both slice with (len=5,cap=10,pointing to len-10 array)
	sl4 := new([10]int)[0:5]
	sl6 := make([]int, 5, 10)
	sl7 := make([]int, 6, 9)
	sl4[1] = 3
	slsl7 := [][]int{sl4, sl6, sl7}
	typeCheck(sl1, sl2, sl3, arrP, sl4, sl5, sl6, slsl7)
}

func browseArrs(arrs ...interface{}) {
	for _, arr := range arrs {
		fmt.Println(arr)
	}
}

func arrInitials() {
	arrAge := [5]int{18, 23, 45, 56, 47}
	arrLazy := [...]int{34, 45, 45, 56, 56}
	arrLazy2 := []int{34, 45, 45, 56, 56}
	arrStr := [7]int{3: 66, 4: 89}
	arrStr2 := [...]int{3: 21, 8: 99}
	browseArrs(arrAge, arrLazy, arrLazy2, arrStr, arrStr2)
	fmt.Println(len(arrStr2))
}

func arrTypeCheck() {
	arrP := new([10]int)
	arrP2 := new([10]int)
	var arr [10]int
	//var arr2 [5]int
	var arr3 [10]int
	arr4 := [2][10]int{arr, arr3}
	arr5 := arr4
	arr6 := [...][2][10]int{arr4, arr5}
	arr7 := [...]*[10]int{arrP, arrP2}
	typeCheck(arr4, arr5, arr6, arrP, arrP2, arr7)
}

func fibonacciWithArray(a int) (int, time.Duration) {
	start := time.Now()
	if a >= ARRMAX {
		return -1, 0
	}
	fib[0], fib[1] = 1, 1
	for i := 2; i <= a; i += 2 {
		fib[i] = fib[i-1] + fib[i-2]
		fib[i+1] = fib[i] + fib[i-1]
	}
	return fib[a], time.Now().Sub(start)

}

func scan2DimSlice(slsl [][]int) {
	for _, sls := range slsl {
		for _, v := range sls {
			fmt.Printf(" %d ", v)
		}
		fmt.Println()
	}
}

func accuV(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func accuA(arr []float32) float32 {
	var sum float32
	for _, v := range arr {
		sum += v
	}
	return sum
}

func expandSlice(slice []int, factor int) []int {
	newSlice := make([]int, len(slice), len(slice)*factor)
	copy(newSlice, slice)
	return newSlice
}

func filter(s []int, fn func(int) bool) []int {
	newSlice := make([]int, 0, len(s))
	for _, v := range s {
		if fn(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func insertStringSlice(dst, src []byte, st int) []byte {
	if st > len(dst)-1 {
		return nil
	}
	newSlice := make([]byte, 0, len(src)+len(dst))

	newSlice = append(newSlice, dst[:st]...)
	newSlice = append(newSlice, src...)
	newSlice = append(newSlice, dst[st:]...)
	return newSlice
}

func removeStringSlice(slice []byte, start, end int) []byte {
	if start < 0 || end > len(slice)-1 {
		return nil
	}
	return append(slice[:start], slice[end+1:]...)
}
