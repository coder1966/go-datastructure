package circlequeue

import (
	"errors"
	"fmt"
)

type Queuer interface {
	Clear()                           // 清空
	Size() int                        // 大小
	Front() interface{}               // 第一个元素
	End() interface{}                 // 最后一个元素
	IsEmpty() bool                    // 空了
	IsFull() bool                     // 满了
	EnQueue(newVal interface{}) error // 入队
	DeQueue() interface{}             // 出队
	String() string                   // 转换
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

type Queue struct {
	dataStore []interface{} // 数组存储
	cap       int           // 最大储存
	front     int           // 头部位置，新入队的在这个index下
	rear      int           // 尾部位置，最早入队的在这个index下
	theSize   int           // 内容大小
}

func NewQueue(cap int) *Queue {
	q := new(Queue)
	q.dataStore = make([]interface{}, cap)
	q.cap = cap
	q.front = 0
	q.rear = 0
	q.theSize = 0
	return q
}

// 实现所有接口方法
func (q *Queue) Size() int     { return q.theSize }
func (q *Queue) IsEmpty() bool { return q.theSize == 0 }
func (q *Queue) IsFull() bool  { return q.theSize >= q.cap }
func (q *Queue) Clear() {
	q.front = 0
	q.rear = 0
	q.theSize = 0
}
func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[q.front]
}
func (q *Queue) End() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[q.rear]
}
func (q *Queue) EnQueue(newVal interface{}) error {
	if q.theSize >= q.cap {
		return errors.New("环形队列满了")
	}

	q.dataStore[q.front] = newVal
	q.front, _ = q.iAdd(q.front, 1)
	// 不检查出圈
	// if err != nil {
	// 	return fmt.Errorf("q.front,err = q.iAdd(q.front, 1) err: %v", err)
	// }
	q.theSize++
	return nil
}
func (q *Queue) DeQueue() interface{} {
	if q.Size() == 0 {
		return nil
	}

	data := q.dataStore[q.rear]
	var err error
	q.rear, err = q.iAdd(q.rear, 1)
	// 检查出圈
	if err != nil {
		fmt.Printf("q.rear, err = q.iAdd(q.rear, 1) err: %v", err)
		return nil
	}
	q.theSize--
	return data
}
func (q *Queue) String() string {
	var err error
	fmt.Println("q.theSize : ", q.theSize)
	fmt.Println("q.front   : ", q.front)
	fmt.Println("q.rear    : ", q.rear)
	fmt.Println("q.dataStore : ", q.dataStore)
	str := ""
	c := q.rear
	for i := 0; i < q.theSize; i++ {
		if i != 0 {
			str = str + " , "
			c, err = q.iAdd(c, 1)
			if err != nil {
				fmt.Println("String() q.iAdd(c, 1) error: ", err)
			}
		}
		str = str + fmt.Sprint(q.dataStore[c])
	}
	return str
}

// a 指针行走 b 步的结果指针(b可以小于0)
func (q *Queue) iAdd(a, b int) (int, error) {
	// if !q.isInCircle(a) {
	// 	return 0, fmt.Errorf("指针a 不在队列里, %d|%d|%d", a, q.rear, q.rear)
	// }

	// 一步一步的走
	c := a
	if b >= 0 {
		for i := 0; i < b; i++ {
			c++
			if c >= q.cap {
				c = 0
			}
			if i < b-1 {
				// 最后一步不检查
				if !q.isInCircle(c) {
					return c, fmt.Errorf("指针c 经过的路径不在队列里, %d|%d|%d", c, q.rear, q.rear)
				}
			}

		}
	} else {
		for i := 0; i < -b; i++ {
			c--
			if c == -1 {
				c = q.cap - 1
			}
			if i < -b-1 {
				// 最后一步不检查
				if !q.isInCircle(c) {
					return c, fmt.Errorf("指针c 经过的路径不在队列里, %d|%d|%d", c, q.rear, q.rear)
				}
			}

		}
	}
	return c, nil

}

// a 是否在圈
func (q *Queue) isInCircle(a int) bool {
	// 出大圈了
	if a >= q.cap || a < 0 {
		return false
	}
	// 头=尾，空的
	if q.front == q.rear {
		return false
	}
	// 头>尾，
	if q.front >= q.rear {
		if a >= q.rear && a <= q.front {
			return true
		} else {
			return false
		}
	}
	// 头<尾，说明转圈了。
	if q.front < q.rear {
		if a >= q.front || a <= q.rear {
			return true
		} else {
			return false
		}
	}
	return false // 不可能到这里
}
