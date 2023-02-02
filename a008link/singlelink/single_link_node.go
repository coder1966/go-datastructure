package singlelink

type SingleLinkNode struct {
	value interface{}
	pNext *SingleLinkNode
}

func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

// 实现接口的各种方法
func (n *SingleLinkNode) GetValue() interface{} { return n.value }
func (n *SingleLinkNode) Next() *SingleLinkNode { return n.pNext }
