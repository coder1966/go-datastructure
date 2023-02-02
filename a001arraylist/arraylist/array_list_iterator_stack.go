package arraylist

import (
	"errors"
)

// 接口
type StackArrayX interface {
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
type StackX struct {
	MyArray *ArrayList // 利用我的数组存储
	capSize int        // 最大大小
	MyIt    Iterator   // 代表一个迭代器
}

// 构造函数，开辟内存
func NewArrayListStackX(cap int) *StackX {
	s := new(StackX)              // 初始化结构体
	s.MyArray = NewArrayList()    // 利用我的数组存储
	s.MyIt = s.MyArray.Iterator() // 迭代的作用
	s.capSize = cap
	return s
}

// 实现接口的所有方法
func (s *StackX) Size() int     { return s.MyArray.theSize }
func (s *StackX) IsFull() bool  { return s.MyArray.theSize >= s.capSize }
func (s *StackX) IsEmpty() bool { return s.MyArray.theSize == 0 }
func (s *StackX) Clear() {
	s.MyArray.Clear() // 利用我的
	s.MyArray.theSize = 0
}
func (s *StackX) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	last := s.MyArray.dataStore[s.MyArray.theSize-1]
	s.MyArray.Delete(s.MyArray.theSize - 1)
	return last
}
func (s *StackX) Push(newVal interface{}) error {
	if s.IsFull() {
		return errors.New("满的，不能压入")
	}
	s.MyArray.Append(newVal)
	return nil
}
func (s *StackX) String() string {
	return s.MyArray.String()
}
