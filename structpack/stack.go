package structpack

import (
	"fmt"
	"math"
)

type Stack interface {
	Init()
	Push(interface{})
	Pop() interface{}
	String() string
}

//StackList is a stack implemented by list of interface{}
type StackList struct {
	bottom *listNode
	top    *listNode
	len    int
	ready  bool
}

//Init inits the top node of Stacklist
func (sta *StackList) Init() {
	sta.top = new(listNode)
	sta.top.val = "Bottom-of-stack"
	sta.bottom = sta.top
	sta.ready = true
	//return sta
}

//Push pushes an interface{} into the StackList
func (sta *StackList) Push(i interface{}) {
	if !sta.ready {
		return
	}
	newNode := listNode{next: sta.top, val: i}
	sta.top = &newNode
	sta.len++
}

//Pop pops and returns an interface{} from the StackList
func (sta *StackList) Pop() interface{} {
	if !sta.ready {
		return "Stack-not-ready"
	}
	if sta.top == sta.bottom {
		return "Bottom-of-stack"
	}
	val := sta.top.val
	sta.top = sta.top.next
	return val
}

//String used for fmt.Println etc.
func (sta *StackList) String() string {
	if !sta.ready {
		return "Stack-not-ready"
	}
	str := "StackList:\n"
	for ptr := sta.top; ptr != nil; ptr = ptr.next {
		str += fmt.Sprintf("%p: %v ", ptr, ptr.val)
	}
	return str
}

//StackArr is a stack implemented by slice of interface{}
type StackArr struct {
	head  int
	arr   []interface{}
	ready bool //whether the stack is ready
}

//Init inits the array of StackArr
func (sta *StackArr) Init() {
	sta.head = -1
	sta.arr = make([]interface{}, 0)
	sta.ready = true
	//return sta
}

//Push pushes an interface{} into the StackArr
func (sta *StackArr) Push(i interface{}) {
	if !sta.ready {
		return
	}
	sta.arr = append(sta.arr, i)
	sta.head++
}

//Pop pops and returns an interface{} from the StackArr
func (sta *StackArr) Pop() interface{} {
	if !sta.ready {
		return nil
	}
	if sta.head < 0 {
		return -math.MaxInt64
	}
	e := sta.arr[sta.head]
	sta.head--
	sta.arr = sta.arr[:sta.head+1]
	return e
}

//At returns the element at position:pos
func (sta *StackArr) At(i int) interface{} {
	if !sta.ready {
		return -math.MaxInt64
	}
	if i > sta.head || i < 0 || sta.head < 0 {
		return -math.MaxInt64
	}
	return sta.arr[i]
}

//String used for fmt.Println etc.
func (sta *StackArr) String() string {
	if !sta.ready {
		return "Stack not ready"
	}
	str := "StackArr: \n"
	for i := 0; i <= sta.head; i++ {
		str += fmt.Sprintf("%d: %v  ", i, sta.arr[i])
	}
	return str
}
