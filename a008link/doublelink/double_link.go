package doublelink

import (
	"fmt"
)

type DoubleLink interface {
	GetLength() int
	GetHead() *DoubleLinkNode                             //
	GetTail() *DoubleLinkNode                             //
	GetNodeAtIndex(index int) *DoubleLinkNode             //
	InsertNodeFront(node *DoubleLinkNode)                 //
	InsertNodeBack(node *DoubleLinkNode)                  //
	InsertNodeBeforeNode(dest, node *DoubleLinkNode) bool // dest 前面插入 node
	InsertNodeAfterNode(dest, node *DoubleLinkNode) bool  // dest 后面插入 node
	DeleteNode(node *DoubleLinkNode) bool                 //
	DeleteAtIndex(index int) bool                         //
	String() string                                       //
	// DoubleLinkFindMiddle() *DoubleLinkNode
	// DoubleLinkFindAny(x, y int) *DoubleLinkNode
	// LinkReverse()

	// Clear()                                    // 清空
	// Size() int                                 // 大小
	// Pop() interface{}                          // 推出元素
	// Push(newVal interface{}) error             // 压入元素
	// IsFull() bool                              // 满了
	// IsEmpty() bool                             // 空了
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

// 链表的结构体
type DoubleLinkList struct {
	head   *DoubleLinkNode // 链表的头指针，prev永远=nil
	tail   *DoubleLinkNode // 链表的尾巴指针，next永远=nil
	length int             // 链表长度
}

// 创建链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	tail := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, tail, 0}
}

// 实现接口的所有方法
func (l *DoubleLinkList) GetLength() int           { return l.length }
func (l *DoubleLinkList) GetHead() *DoubleLinkNode { return l.head.next }
func (l *DoubleLinkList) GetTail() *DoubleLinkNode { return l.tail.prev }
func (l *DoubleLinkList) GetNodeAtIndex(index int) *DoubleLinkNode {
	// 第1个节点的编号是1不是0，0是head不带有数据
	if index < 1 {
		return nil
	}
	if l.length == 0 {
		return nil
	}
	bak := l.head
	for i := 0; i < index; i++ {
		bak = bak.next
		if bak == l.tail {
			return nil
		}
	}
	return bak
}
func (l *DoubleLinkList) InsertNodeFront(node *DoubleLinkNode) {
	if l.length == 0 {
		// 空的
		node.next = l.tail
		node.prev = l.head
		l.head.next = node
		l.tail.prev = node
		l.length = 1
		return
	}
	// 新 上-> 左， 新 下-> 原右，左 下-> 新，原右 上-> 新
	node.prev, node.next, l.head.next, l.head.next.prev = l.head, l.head.next, node, node
	l.length++
}
func (l *DoubleLinkList) InsertNodeBack(node *DoubleLinkNode) {
	if l.length == 0 {
		// 空的
		node.next = l.tail
		node.prev = l.head
		l.head.next = node
		l.tail.prev = node
		l.length = 1
		return
	}
	// 新 上-> 原尾的右， 新 下-> 尾，左(尾.左) 下-> 新，尾巴 上-> 新
	node.prev, node.next, l.tail.prev.next, l.tail.prev = l.tail.prev, l.tail, node, node
	l.length++
}
func (l *DoubleLinkList) InsertNodeBeforeNode(dest, node *DoubleLinkNode) bool {
	if dest == nil {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for bak.next != dest {
		bak = bak.next
		if bak.next == nil {
			return false
		}
	}
	// 到这里，要被插入的是bak的下一个
	node.next = bak.next
	bak.next = node
	l.length++
	return true
}
func (l *DoubleLinkList) InsertNodeAfterNode(dest, node *DoubleLinkNode) bool {
	if dest == nil {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for bak != dest {
		bak = bak.next
		if bak == nil {
			return false
		}
	}
	// 到这里，要被插入的是bak
	node.next = bak.next
	bak.next = node
	l.length++
	return true
}
func (l *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	// 第1个节点的编号是1不是0，0是head不带有数据
	if l.length == 0 {
		return false
	}

	bak := l.head
	for bak.next != node {
		bak = bak.next
		if bak == l.tail {
			return false
		}
	}
	// 到这里，要删除的是bak的下一个
	bak.next = bak.next.next
	l.length--
	return true
}
func (l *DoubleLinkList) DeleteAtIndex(index int) bool {
	// 第1个节点的编号是1不是0，0是head不带有数据
	if index < 1 {
		return false
	}
	if l.length < index {
		return false
	}
	bak := l.head
	for i := 0; i < index; i++ {
		bak = bak.next
		if bak == l.tail {
			return false
		}
	}
	// 到这里，要删除的是bak本身
	// 删左 下-> 删右， 删右 上-> 删左
	bak.prev.next, bak.next.prev = bak.next, bak.prev

	bak.next = bak.next.next

	l.length--
	return true
}
func (l *DoubleLinkList) String() string {
	p := l.head

	listStr := fmt.Sprintf("len=%v , ", l.length)

	for p != nil {

		if p.prev == nil {
			listStr += fmt.Sprintf("nil|")
		} else {
			listStr += fmt.Sprintf("%v|", p.prev.value)
		}

		if p == nil {
			listStr += fmt.Sprintf("nil|")
		} else {
			listStr += fmt.Sprintf("%v|", p.value)
		}

		if p.next == nil {
			listStr += fmt.Sprintf("nil;")
		} else {
			listStr += fmt.Sprintf("%v;", p.next.value)
		}

		p = p.next
	}
	listStr += fmt.Sprintf("Tail.Prev:%v;", l.tail.prev.value)

	return listStr
}
