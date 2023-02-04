package heap

type PriorityItem struct {
	value    interface{} // 数据
	Priority int         // 优先级

}

// 构造函数，构造队列中间的一个元素
func NewPriorityItem(v interface{}, p int) *PriorityItem { return &PriorityItem{v, p} }

// 实现比大小 对比的是优先级
func (x PriorityItem) Less(than Item) bool {
	return x.Priority < than.(PriorityItem).Priority // 对比的是优先级
}

// 优先队列，基于堆来实现
type PriorityQueue struct {
	data *Heap
}

// 最大堆
func NewMinPriorityQueue() *PriorityQueue { return &PriorityQueue{NewMin()} }

// 最小堆
func NewMaxPriorityQueue() *PriorityQueue { return &PriorityQueue{NewMax()} }

func (pq *PriorityQueue) Len() int { return pq.data.Len() }

func (pq *PriorityQueue) Append(el Item) {
	pq.data.Append(el)
}

func (pq *PriorityQueue) Extract() PriorityItem {
	return pq.data.Extract().(PriorityItem)
}

func (pq *PriorityQueue) ChangPriority(val interface{}, pri int) {
	storage := NewQueue()  // 一个新队列，把这个备份存储一下
	popped := pq.Extract() // 拿出一个最小的
	for val != popped.value {
		if pq.Len() == 0 {
			return // 出错了
		}
		storage.Push(popped) // 没找到的，压入数据

		popped = pq.Extract() // 再拿出一个最小的
	}

	// 找到了
	popped.Priority = pri // 修改了优先级

	pq.data.Append(popped)

	// 其余数据，放入重新队列
	for storage.Len() > 0 {
		pq.data.Append(storage.Shift().(Item))
	}
}
