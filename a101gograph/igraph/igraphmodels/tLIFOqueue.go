package igraphmodels

// LIFO堆栈，实现INodeQueue 接口
type tLIFOqueue struct {
	nodes    []INode
	capacity int
	size     int
}

// Clear 方法，清空
func (t *tLIFOqueue) Clear() {
	t.nodes = make([]INode, 0)
	t.capacity = 0
	t.size = 0
}

// Empty 方法，是否是空队列
func (t *tLIFOqueue) Empty() bool {
	return t.size <= 0
}

// Push 方法，压栈一个节点到队列
func (t *tLIFOqueue) Push(node INode) {
	t.ensureSpace(1)
	t.nodes[t.size] = node
	t.size++
}

// 方法，开辟保障空间
func (t *tLIFOqueue) ensureSpace(space int) {
	for t.capacity < t.size+space {
		t.nodes = append(t.nodes, nil)
		t.capacity++
	}
}

// Poll 方法，队列里弹出一个节点
func (t *tLIFOqueue) Poll() (bool, INode) {
	if t.Empty() {
		return false, nil
	}

	t.size--
	it := t.nodes[t.size]
	t.nodes[t.size] = nil

	return true, it
}

// NewLIFOQueue 构造函数，返回类型是接口，return的是实现，逼迫成功实现所有方法
func NewLIFOQueue() iNodeQueue {
	it := &tLIFOqueue{}
	it.Clear()
	return it
}
