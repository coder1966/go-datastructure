package a005link

import "fmt"

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	Length() int                // 大小
	EnQueue(newVal interface{}) // 入列
	DeQueue() interface{}       // 出列
	String() string             // 转
}

func NewQueueLink() *QueueLink {
	return &QueueLink{}
}

// 实现接口的方法
func (q *QueueLink) Length() int {
	tmp := q.front
	length := 0
	for tmp.pNext != nil {
		tmp = tmp.pNext
		length++
	}
	return length
}
func (q *QueueLink) EnQueue(newVal interface{}) {
	// 头部插入，头部删除，头一个算句柄不算节点
	newNode := &Node{data: newVal, pNext: nil}
	if q.front == nil {
		// 头是空，那就是全空
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.pNext = newNode // 挂上新节点
		q.rear = q.rear.pNext  // 尾部指针后移
	}
}
func (q *QueueLink) DeQueue() interface{} {
	if q.front == nil {
		// 头是空，那就是全空
		return nil
	}
	newNode := q.front // 记录头部位置

	if q.front == q.rear {
		// 说明只剩下一个
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.pNext // 相当于删除了一个
	}

	return newNode.data
}
func (q *QueueLink) String() string {
	tmp := q.front
	str := ""
	for tmp != nil {
		str = str + " , " + fmt.Sprint(tmp.data)
		tmp = tmp.pNext
	}

	return str
}
