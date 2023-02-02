package queue

import "fmt"

type Queuer interface {
	Clear()                     // 清空
	Size() int                  // 大小
	Front() interface{}         // 第一个元素
	End() interface{}           // 最后一个元素
	IsEmpty() bool              // 空了
	EnQueue(newVal interface{}) // 入队
	DeQueue() interface{}       // 出队
	String() string             // 转换
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

type Queue struct {
	dataStore []interface{} // 数组存储
	theSize   int           // 内容大小
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.dataStore = make([]interface{}, 0, 16)
	queue.theSize = 0
	return queue
}

// 实现所有接口方法
func (q *Queue) Size() int     { return q.theSize }
func (q *Queue) IsEmpty() bool { return q.theSize == 0 }
func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0, 16)
	q.theSize = 0
}
func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[0]
}
func (q *Queue) End() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[q.theSize-1]
}
func (q *Queue) EnQueue(newVal interface{}) {
	q.dataStore = append(q.dataStore, newVal)
	q.theSize++
}
func (q *Queue) DeQueue() interface{} {
	if q.Size() == 0 {
		return nil
	}
	data := q.dataStore[0]
	q.dataStore = q.dataStore[1:q.theSize]
	q.theSize--
	return data
}
func (q *Queue) String() string {
	// fmt.Println("q.theSize : ", q.theSize)
	// fmt.Println("q.dataStore : ", q.dataStore)
	str := ""
	for i := 0; i < q.theSize; i++ {
		if i != 0 {
			str = str + " , "
		}
		str = str + fmt.Sprint(q.dataStore[i])
	}
	return str
}
