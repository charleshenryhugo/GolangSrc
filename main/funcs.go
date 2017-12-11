package main

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"time"
)

func varnumpar() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimun is %d\n", x)

	arr := []int{7, 9, 3, 5, 1}
	x = min(arr...)
	fmt.Printf("The minimun in the array arr is %d\n", x)
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min

}

// check the types of unknown-type&number inputs
// with use of interfaces{}
func typeCheck(values ...interface{}) {
	for _, value := range values {
		fmt.Printf("%v is type of: ", value)
		fmt.Println(reflect.TypeOf(value))
	}
}

func anonFuncDemo() {
	/*f1 := retFunc1()
	fmt.Println(f1(5))
	f2 := retFunc2(3)
	fmt.Println(f2(5))
	*/
	f3 := retFunc3()
	fmt.Println(f3(1))
	fmt.Println(f3(20))
	fmt.Println(f3(300))

}

func retFunc1() func(int) int {
	return func(b int) int { return b + 2 }
}

func retFunc2(a int) func(int) int {
	return func(b int) int { return b + a }
}

func retFunc3() func(int) int {
	var x int //x = 0
	return func(delta int) int {
		x += delta //x ia remembered even if outer-func have returned
		return x
	}
}

func retfunc4(v0 int, v1 int) func() (int, int) {
	return func() (int, int) { //v0, v1 will be kept
		v0 = v0 + v1
		v1 = v1 + v0
		//fmt.Println(v0, " ", v1)
		return v0, v1
	}
}

//using closure function to solve fibonacci is faster than recursive-method
func fibonacciWithClosureFunc(a int) (int, time.Duration) {
	start := time.Now()
	if a == 0 || a == 1 {
		return 1, 0
	}

	a0, a1 := 1, 1
	f := retfunc4(a0, a1)
	for i := 2; i <= a; i += 2 {
		//fmt.Print(i, ": ")
		a0, a1 = f() //v0 and v1 increased iteratively
	}
	if a%2 == 0 {
		return a0, time.Now().Sub(start)
	}
	return a1, time.Now().Sub(start)
}

//makeAddSuffix is a factory-function that produces a series of functions,
//which can add specific suffix after an input string
func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return (name + suffix)
		}
		return name
	}
}

func anonFuncDemo2() {
	addJpeg := makeAddSuffix(".JPEG")
	addBmp := makeAddSuffix(".BMP")

	file := "image"
	fmt.Println(addJpeg(file))
	fmt.Println(addBmp(file))
}

func mapFunc(fn func(int) int, intList ...int) []int {
	resList := make([]int, 0, len(intList))
	for _, v := range intList {
		resList = append(resList, fn(v))
	}
	return resList
}

func reboot() error {
	const UNIX_REBOOT_MAGIC1 uintptr = 0xfee1dead
	const UNIX_REBOOT_MAGIC2 uintptr = 672274793
	const UNIX_REBOOT_CMD_RESTART uintptr = 0x1234567
	_, _, err := syscall.Syscall(syscall.SYS_REBOOT,
		UNIX_REBOOT_MAGIC1, UNIX_REBOOT_MAGIC2,
		UNIX_REBOOT_CMD_RESTART)
	return err
}
