package a010queuethreadsafe

import (
	"sync"
)

// type Queuer interface {
// 	Clear()                     // 清空
// 	Size() int                  // 大小
// 	Front() interface{}         // 第一个元素
// 	End() interface{}           // 最后一个元素
// 	IsEmpty() bool              // 空了
// 	EnQueue(newVal interface{}) // 入队
// 	DeQueue() interface{}       // 出队
// 	String() string             // 转换
// 	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
// }

type Queue struct {
	dataStore []interface{} // 数组存储
	len       int           // 内容大小
	lock      *sync.Mutex   // 锁
}

func New() *Queue {
	// queue := new(Queue)
	queue := &Queue{}
	queue.dataStore = make([]interface{}, 0, 16)
	queue.len = 0
	queue.lock = new(sync.Mutex)
	return queue
}

// 实现各种方法
func (q *Queue) Len() int {
	q.lock.Lock() // 解决了线程安全
	defer q.lock.Unlock()
	return q.len
}
func (q *Queue) IsEmpty() bool {
	q.lock.Lock() // 解决了线程安全
	defer q.lock.Unlock()
	return q.len == 0
}

// 弹出数据
func (q *Queue) Shift() (el interface{}) {
	q.lock.Lock() // 解决了线程安全
	defer q.lock.Unlock()
	el, q.dataStore = q.dataStore[0], q.dataStore[1:]
	q.len--
	return
}

// 压入数据
func (q *Queue) Push(el interface{}) {
	q.lock.Lock() // 解决了线程安全
	defer q.lock.Unlock()
	q.dataStore = append(q.dataStore, el)
	q.len++
	return
}

// 偷看，只是看元素
func (q *Queue) Peek() interface{} {
	q.lock.Lock() // 解决了线程安全
	defer q.lock.Unlock()
	return q.dataStore[0]
}
