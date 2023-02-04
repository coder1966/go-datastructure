package doublelink

import "fmt"

type DoubleLinkNode struct {
	value interface{}
	prev  *DoubleLinkNode
	next  *DoubleLinkNode
}

func NewDoubleLinkNode(data interface{}) *DoubleLinkNode {
	return &DoubleLinkNode{data, nil, nil}
}

// 实现接口的各种方法
func (n *DoubleLinkNode) GetValue() interface{} { return n.value }
func (n *DoubleLinkNode) Prev() *DoubleLinkNode { return n.prev }
func (n *DoubleLinkNode) Next() *DoubleLinkNode { return n.next }
func (n *DoubleLinkNode) String() string {
	listStr := ""

	if n == nil {
		return "nil|nil|nil"
	}

	if n.prev == nil {
		listStr += fmt.Sprintf("nil|")
	} else {
		listStr += fmt.Sprintf("%v|", n.prev.value)
	}

	if n == nil {
		listStr += fmt.Sprintf("nil|")
	} else {
		listStr += fmt.Sprintf("%v|", n.value)
	}

	if n.next == nil {
		listStr += fmt.Sprintf("nil;")
	} else {
		listStr += fmt.Sprintf("%v;", n.next.value)
	}

	return listStr
}
