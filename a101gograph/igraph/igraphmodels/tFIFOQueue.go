package igraphmodels

// FIFO 队列，实现INodeQueue 接口
type tFIFOQueue struct {
	nodes    []INode // 队列内所有顶点
	capacity int     // 容量
	rindex   int     // 读的头位置
	windex   int     // 写的尾巴位置
}

// Clear 清空队列方法
func (t *tFIFOQueue) Clear() {
	t.nodes = make([]INode, 0) // 所有顶点清空
	t.capacity = 0             // 容量
	t.rindex = -1              // 索引
	t.windex = -1              // 索引
}

// 队列尺寸
func (t *tFIFOQueue) size() int {
	return t.windex - t.rindex
}

// Empty 队列是否为空
func (t *tFIFOQueue) Empty() bool {
	return t.size() <= 0
}

// 确保空间
func (t *tFIFOQueue) ensureSpace(size int) {
	for t.capacity < t.windex+size+1 {
		t.nodes = append(t.nodes, nil)
		t.capacity++
	}
}

// Push 压入栈 顶点
func (t *tFIFOQueue) Push(node INode) {
	t.ensureSpace(1)         // 确保空间 1
	t.windex++               // 写索引
	t.nodes[t.windex] = node // 写索引位置的顶点内容
}

// Poll 获取并删除Queue中的第一个元素
func (t *tFIFOQueue) Poll() (bool, INode) {
	if t.Empty() { // 空，返回nil
		return false, nil
	}
	t.rindex++              // 读指针
	it := t.nodes[t.rindex] // 准备返回的顶点
	t.nodes[t.rindex] = nil // 返回走了，队列里的就删掉

	if t.rindex > t.capacity/2 { // 读指针过总容量一半（浪费了一半的空间），排挤掉前面的空闲空间
		size := t.size()            // 目前队列尺寸
		offset := t.rindex + 1      // 抵消、补偿=读索引+1
		for i := 0; i < size; i++ { // 跨越offset个逐个左移
			t.nodes[i], t.nodes[i+offset] = t.nodes[i+offset], nil
		}
		t.rindex -= offset // 读写指针调整
		t.windex -= offset // 读写指针调整
	}
	return true, it
}

// 构造函数。函数返回参数是“接口”，return 的却是实现接口的队列。（逼迫这个实现队列加上一系列方法）
func newFIFOQueue() iNodeQueue {
	it := &tFIFOQueue{}
	it.Clear()
	return it
}
