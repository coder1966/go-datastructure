package stackarray

import (
	"errors"
	"fmt"
)

// 接口
type StackArray interface {
	Clear()                        // 清空
	Size() int                     // 大小
	Pop() interface{}              // 推出元素
	Push(newVal interface{}) error // 压入元素
	IsFull() bool                  // 满了
	IsEmpty() bool                 // 空了
	String() string
	// Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

// 数据结构。是泛型的，所以用空接口
type Stack struct {
	dataStore   []interface{} // 数组存储
	capSize     int           // 最大大小
	currentSize int           // 当前实际使用大小

}

// 构造函数，开辟内存
func NewStack(cap int) *Stack {
	s := new(Stack) // 初始化结构体
	s.dataStore = make([]interface{}, cap)
	s.capSize = cap
	s.currentSize = 0
	return s
}

// 实现接口的所有方法
func (s *Stack) Size() int     { return s.currentSize }
func (s *Stack) IsFull() bool  { return s.currentSize >= s.capSize }
func (s *Stack) IsEmpty() bool { return s.currentSize == 0 }
func (s *Stack) Clear() {
	s.dataStore = make([]interface{}, s.capSize)
	s.currentSize = 0
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.currentSize-- // 指针偏移回来，就是减掉一个
	return s.dataStore[s.currentSize]
}
func (s *Stack) Push(newVal interface{}) error {
	if s.IsFull() {
		return errors.New("满的，不能压入")
	}
	s.dataStore[s.currentSize] = newVal // 这个位置，就是栈顶
	s.currentSize++                     // 指针涨一个
	return nil
}
func (s *Stack) String() string {
	str := ""
	for i := 0; i < s.currentSize; i++ {
		if i != 0 {
			str = str + " , "
		}
		str = str + fmt.Sprint(s.dataStore[i])
	}
	return str
}
