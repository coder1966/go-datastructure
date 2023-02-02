package arraylist

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex() int              // 得到索引

}

type Iterable interface {
	Iterator() Iterator // 构造的时候，需要初始化一个接口。这个函数，返回迭代器接口。
}

// 构造指针，访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

// ArrayList 实现这个接口
func (l *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // 构造迭代器
	it.currentIndex = 0
	it.list = l
	return it
}

// ArrayListIterator 实现接口
func (it *ArrayListIterator) HasNext() bool { return it.currentIndex < it.list.theSize }
func (it *ArrayListIterator) GetIndex() int { return it.currentIndex }
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个了")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}
