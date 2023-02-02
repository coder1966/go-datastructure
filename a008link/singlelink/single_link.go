package singlelink

import (
	"fmt"
)

type SingleLink interface {
	GetHead() *SingleLinkNode                             //
	GetNodeAtIndex(index int) *SingleLinkNode             //
	InsertNodeFront(node *SingleLinkNode)                 //
	InsertNodeBack(node *SingleLinkNode)                  //
	InsertNodeBeforeNode(dest, node *SingleLinkNode) bool // dest 前面插入 node
	InsertNodeAfterNode(dest, node *SingleLinkNode) bool  // dest 后面插入 node
	DeleteNode(node *SingleLinkNode) bool                 //
	DeleteAtIndex(index int) bool                         //
	String() string                                       //
	SingleLinkFindMiddle() *SingleLinkNode
	SingleLinkFindAny(x, y int) *SingleLinkNode
	LinkReverse()

	// Clear()                                    // 清空
	// Size() int                                 // 大小
	// Pop() interface{}                          // 推出元素
	// Push(newVal interface{}) error             // 压入元素
	// IsFull() bool                              // 满了
	// IsEmpty() bool                             // 空了
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

// 链表的结构体
type SingleLinkList struct {
	head   *SingleLinkNode // 链表的头指针
	length int             // 链表长度
}

// 创建链表
func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)
	return &SingleLinkList{head, 0}
}

// 实现接口的所有方法
func (l *SingleLinkList) GetHead() *SingleLinkNode { return l.head.pNext }
func (l *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	// 第1个节点的编号是1不是0，0是head不带有数据
	if index < 1 {
		return nil
	}
	if l.head == nil {
		return nil
	}
	bak := l.head
	for i := 0; i < index; i++ {
		bak = bak.pNext
		if bak == nil {
			return nil
		}
	}
	return bak
}
func (l *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if l.head == nil {
		l.head = node
		node.pNext = nil
		l.length = 1
		return
	}
	bak := l.head
	node.pNext = bak.pNext
	bak.pNext = node
	l.length++
}
func (l *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if l.head == nil {
		l.head = node
		node.pNext = nil
		l.length = 1
		return
	}
	bak := l.head
	for bak.pNext != nil {
		bak = bak.pNext
	}
	node.pNext = nil
	bak.pNext = node
	l.length++
}
func (l *SingleLinkList) InsertNodeBeforeNode(dest, node *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for bak.pNext != dest {
		bak = bak.pNext
		if bak.pNext == nil {
			return false
		}
	}
	// 到这里，要被插入的是bak的下一个
	node.pNext = bak.pNext
	bak.pNext = node
	l.length++
	return true
}
func (l *SingleLinkList) InsertNodeAfterNode(dest, node *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for bak != dest {
		bak = bak.pNext
		if bak == nil {
			return false
		}
	}
	// 到这里，要被插入的是bak
	node.pNext = bak.pNext
	bak.pNext = node
	l.length++
	return true
}
func (l *SingleLinkList) DeleteNode(node *SingleLinkNode) bool {
	if node == nil {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for bak.pNext != node {
		bak = bak.pNext
		if bak.pNext == nil {
			return false
		}
	}
	// 到这里，要删除的是bak的下一个
	bak.pNext = bak.pNext.pNext
	l.length--
	return true
}
func (l *SingleLinkList) DeleteAtIndex(index int) bool {
	// 第1个节点的编号是1不是0，0是head不带有数据
	if index < 1 {
		return false
	}
	if l.head == nil {
		return false
	}
	bak := l.head
	for i := 0; i < index-1; i++ {
		bak = bak.pNext
		if bak.pNext == nil {
			return false
		}
	}
	// 到这里，要删除的是bak的下一个
	bak.pNext = bak.pNext.pNext
	l.length--
	return true
}
func (l *SingleLinkList) String() string {
	p := l.head

	listStr := fmt.Sprintf("len=%v , ", l.length)

	for p.pNext != nil {
		listStr += fmt.Sprintf("%v->", p.pNext.value)
		p = p.pNext
	}
	listStr += fmt.Sprintf("nil")
	return listStr
}

// SingleLinkFindMiddle 查找中间节点
func (l *SingleLinkList) SingleLinkFindMiddle() *SingleLinkNode {
	/*
		设计2个指针，A走一步，B走两步。当B=nil时候，就是a了
	*/

	tmpA := l.head
	tmpB := l.head

	for tmpB != nil {
		if tmpB.pNext == nil {
			// B是尾巴
			return tmpA
		}
		// B不是尾巴
		tmpB = tmpB.pNext.pNext
		tmpA = tmpA.pNext
	}

	return tmpA
}

// SingleLinkFindAny 查找任意比例中间节点
func (l *SingleLinkList) SingleLinkFindAny(x, y int) *SingleLinkNode {
	/*
		设计2个指针，A走x步，B走y步。当B=nil时候，就是a了
	*/
	if x > y {
		return nil
	}
	tmpA := l.head
	tmpB := l.head

	for tmpB != nil {
		// B 走动
		for i := 0; i < y; i++ {
			tmpB = tmpB.pNext
			if tmpB == nil {
				return tmpA
			}
		}
		// A 走动
		for i := 0; i < x; i++ {
			tmpA = tmpA.pNext
		}

	}

	return tmpA
}

// LinkReverse 链表翻转。【臭名昭著的面试题】
func (l *SingleLinkList) LinkReverse() {
	/*
		前序      当前      后续
		pre->cur cur->cNe  cNe       // curNext := cur.pNext
		pre->cur cur->cNe  cNe->pre // cur.pNext = pre 第二指向第一
		pre->cNe cur->cNe  cNe->pre // pre = cur       第一换成第二
		pre->cNe cur->pre  cNe->pre // cur = curNext   第二换成第三

	*/
	if l.head == nil || l.head.pNext == nil {
		return
	}

	// var pre *SingleLinkNode                // 前一个节点
	// var cur *SingleLinkNode = l.head.pNext // 当前的节点 current
	// for cur != nil {
	// 	curNext := cur.pNext // 后续节点 第三
	// 	cur.pNext = pre      // 第二指向第一
	// 	pre = cur            // 第一换成第二，向右推进
	// 	cur = curNext        // 第二换成第三，向右推进
	// }
	// l.head.pNext = pre

	/*
		a=1 b=2 c=3
		A->B B->C
		A<-B B<-C
		a=2 b=3 c=4

		l.head.pNext = c
	*/

	var a *SingleLinkNode
	var b *SingleLinkNode = l.head.pNext
	// c := b.pNext
	for b != nil {
		//平行赋值语法  可读性差
		b.pNext, a, b = a, b, b.pNext
	}
	l.head.pNext = a
}
