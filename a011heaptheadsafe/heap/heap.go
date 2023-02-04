package heap

import "sync"

type Int int

func (i Int) Less(than Item) bool {
	return i < than.(Int)
}

// func (i Int) Value() Int {
// 	return i
// }

type Item interface {
	Less(than Item) bool // 比大小
}

// 实现最小堆
type Heap struct {
	lock *sync.Mutex
	data []Item
	min  bool // 是否最小堆
}

func New() *Heap    { return &Heap{new(sync.Mutex), make([]Item, 0), true} }  // 标准堆
func NewMin() *Heap { return &Heap{new(sync.Mutex), make([]Item, 0), true} }  // 最小堆
func NewMax() *Heap { return &Heap{new(sync.Mutex), make([]Item, 0), false} } // 最大堆

func (h *Heap) IsEmpty() bool { return len(h.data) == 0 }
func (h *Heap) Len() int      { return len(h.data) }
func (h *Heap) Get(index int) *Item {
	if index < len(h.data) {
		return &h.data[index]
	} else {
		return nil
	}
}

// 追加|插入数据
func (h *Heap) Append(it Item) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, it)
	h.ShiftMax() // 追加|插入数据 都要排序一次
}

// 根据大小堆类型返回比大小
func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

// 压缩，弹出一个
func (h *Heap) Extract() (el Item) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.Len() == 0 {
		return // 长度0，不需要处理
	}
	el = h.data[0]
	last := h.data[h.Len()-1] // 最后一个
	if h.Len() == 1 {
		h.data = nil // 弹出唯一的数据
		return

	}

	h.data = append([]Item{last}, h.data[1:h.Len()-1]...) // 最后一个，调到第一个给我用，原来第一个切掉
	// h.ShiftMax()                                          // 都要排序一次
	return
}

// 弹出极大值
func (h *Heap) ShiftMax() {
	// h.lock.Lock()
	// defer h.lock.Unlock()

	// 堆排序的循环过程 n,2n+1这样一个过程
	// 循环冒泡，每次取出来一个极小值
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i / 2
		if h.Less(*h.Get(i), *h.Get(parent)) {
			// 比我小,交换
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}

}

// 弹出极小值
func (h *Heap) ShiftMin() {
	// h.lock.Lock()
	// defer h.lock.Unlock()

	// 堆排序的循环过程 n,2n+1这样一个过程
	// 循环冒泡，每次取出来一个极小值
	for i, child := 0, 1; i < h.Len() && i*2+1 < h.Len(); i = child {
		child = i*2 + 1
		// 对比左右孩子
		if child+1 <= h.Len()-1 && h.Less(*h.Get(child + 1), *h.Get(child)) {
			child++ //循环左右节点
		}

		if h.Less(*h.Get(i), *h.Get(child)) {
			// 比我小
			break
		} else {
			h.data[child], h.data[i] = h.data[i], h.data[child]
		}
	}

}

func (h *Heap) String() []Item {
	return h.data
}
