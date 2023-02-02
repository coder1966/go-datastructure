package doublelink

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
