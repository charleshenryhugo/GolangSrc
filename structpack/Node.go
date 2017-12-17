package structpack

type listNode struct {
	val  interface{}
	prev *listNode
	next *listNode
}
