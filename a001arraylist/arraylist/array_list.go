package arraylist

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newVal interface{}) error    // 修改第几个元素
	Insert(index int, newVal interface{}) error // 插入元素
	Append(newVal interface{})                  // 追加元素
	Clear()                                     // 清空
	Delete(index int) error                     // 删除元素
	String() string
	Iterator() Iterator // (也要放在这里)构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

// 数据结构。是泛型的，所以用空接口
type ArrayList struct {
	dataStore []interface{} // 数组存储
	theSize   int           // 数组的大小

}

// 构造函数，开辟内存
func NewArrayList() *ArrayList {
	list := new(ArrayList) // 初始化结构体
	list.dataStore = make([]interface{}, 0, 16)
	list.theSize = 0
	return list
}

// 实现接口的所有方法
func (l *ArrayList) Size() int { return l.theSize }
func (l *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.theSize {
		return nil, errors.New("索引越界")
	}
	return l.dataStore[index], nil
}
func (l *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= l.theSize {
		return errors.New("索引越界")
	}
	l.dataStore[index] = newVal
	return nil
}
func (l *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= l.theSize {
		return errors.New("索引越界")
	}
	tmp := append(l.dataStore[:index], l.dataStore[index])
	l.dataStore = append(tmp, l.dataStore[index:]...) // 实现拉开，重新叠加，跳过 index
	l.theSize++
	l.dataStore[index] = newVal
	return nil
}
func (l *ArrayList) Append(newVal interface{}) {
	l.dataStore = append(l.dataStore, newVal)
	l.theSize++
}
func (l *ArrayList) Clear() {
	l.dataStore = make([]interface{}, 0, 16)
	l.theSize = 0
}
func (l *ArrayList) Delete(index int) error {
	if index < 0 || index >= l.theSize {
		return errors.New("索引越界")
	}
	l.dataStore = append(l.dataStore[:index], l.dataStore[index+1:]...) // 实现删除，重新叠加，跳过 index
	l.theSize--
	return nil
}
func (l *ArrayList) String() string { return fmt.Sprint(l.dataStore) }
