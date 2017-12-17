package structpack

import (
	"fmt"
	"reflect"
)

//Node is the interface for all nodes
type node interface {
	info()
}

type listNode struct {
	val  interface{}
	prev *listNode
	next *listNode
}

func (n *listNode) info() {
	fmt.Printf("val: %v  prev: %p  next: %p ",
		n.val, n.prev, n.next)
}

type binaryNode struct {
	val   interface{}
	left  *binaryNode
	right *binaryNode
}

func (n *binaryNode) info() {
	fmt.Printf("val: %v  left: %p  right: %p ",
		n.val, n.left, n.right)
}

//NodeTest shows the testcase fot Node
func NodeTest() {
	node := new(listNode)
	tType := reflect.TypeOf(*node)
	fmt.Println(tType)
	nField := tType.NumField()
	for ix := 0; ix < nField; ix++ {
		ixField := tType.Field(ix)
		fmt.Println(ixField.Name)
	}
}
