package arraylist

import (
	"errors"
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
	myArray *ArrayList // 利用我的数组存储
	capSize int        // 最大大小

}

// 构造函数，开辟内存
func NewArrayListStack(cap int) *Stack {
	s := new(Stack)            // 初始化结构体
	s.myArray = NewArrayList() // 利用我的数组存储
	s.capSize = cap
	return s
}

// 实现接口的所有方法
func (s *Stack) Size() int     { return s.myArray.theSize }
func (s *Stack) IsFull() bool  { return s.myArray.theSize >= s.capSize }
func (s *Stack) IsEmpty() bool { return s.myArray.theSize == 0 }
func (s *Stack) Clear() {
	s.myArray.Clear() // 利用我的
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	last := s.myArray.dataStore[s.myArray.theSize-1]
	s.myArray.Delete(s.myArray.theSize - 1)
	return last
}
func (s *Stack) Push(newVal interface{}) error {
	if s.IsFull() {
		return errors.New("满的，不能压入")
	}
	s.myArray.Append(newVal)
	return nil
}
func (s *Stack) String() string {
	return s.myArray.String()
}
