package a005link

import "fmt"

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool           // 空了
	Pop() interface{}        // 推出元素
	Push(newVal interface{}) // 压入元素
	Size() int               // 大小
	String() string          // 转
	// Clear()                        // 清空
	// IsFull() bool                  // 满了
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

func NewStack() *Node {
	return &Node{}
}

// 实现接口的方法
func (n *Node) IsEmpty() bool { return n.pNext == nil }
func (n *Node) Pop() interface{} {
	// 头部插入，头部删除，头一个算句柄不算节点
	if n.IsEmpty() {
		return nil
	}
	value := n.pNext.data   // 要弹出的数据
	n.pNext = n.pNext.pNext // 删除刚弹出的数据
	return value
}
func (n *Node) Push(newVal interface{}) {
	// 头部插入，头部删除，头一个算句柄不算节点
	newNode := &Node{data: newVal, pNext: n.pNext}
	n.pNext = newNode
}
func (n *Node) Size() int {
	tmp := n
	length := 0
	for tmp.pNext != nil {
		tmp = tmp.pNext
		length++
	}
	return length
}
func (n *Node) String() string {
	tmp := n
	str := ""
	for tmp != nil {
		str = str + " , " + fmt.Sprint(tmp.data)
		tmp = tmp.pNext
	}

	return str
}
