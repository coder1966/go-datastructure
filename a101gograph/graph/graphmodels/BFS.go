package graphmodels

import (
	"sync"
)

/*
BFS：广度优先搜索
BFS（Breadth First Search）：广度优先搜索，广度指的是从一个节点开始 发散性地遍历 周围节点。从某个节点出发，访问它的所有邻接节点，再从这些节点出发，访问它们未被访问过得邻接节点…直到所有节点访问完毕。
有点类似树的层序遍历，但图存在成环的情形，访问过的节点可能会再次访问，所以需要用一个辅助队列来存放待访问的邻接节点。
*/

// 所有节点队列
type NodeQueue struct {
	nodes []Node
	lock  sync.RWMutex
}

// 构造函数
func NewNodeQueue() *NodeQueue {
	q := NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = []Node{}
	return &q
}

// 入队列
func (q *NodeQueue) Enqueue(n Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	defer q.lock.RLock()
	node := q.nodes[0]
	q.nodes = q.nodes[1:] // 剪掉头一个元素
	return &node
}

// 空否
func (q *NodeQueue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.nodes) == 0
}

// BFS 遍历
func (g *Graph) BFS(f func(node *Node)) {
	g.lock.Lock()
	defer g.lock.Unlock()

	// 初始化队列
	q := NewNodeQueue()
	// 取图的第一个节点入队列
	head := g.nodes[0]
	q.Enqueue(*head)
	// 表示这个节点是否访问过
	visited := make(map[*Node]bool)
	visited[head] = true
	// 遍历所有节点，知道队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()     // 出一个队列
		visited[node] = true    // 标记访问过
		nexts := g.edges[*node] // 这个节点所有的边
		// 所有没访问过的邻居节点放入队列
		for _, next := range nexts {
			// 访问过就不进
			if visited[next] {
				continue
			}
			q.Enqueue(*next)     // 这个入队列
			visited[next] = true // 标记已经访问
		}

		// 对每个正在遍历的节点执行回调
		if f != nil {
			f(node)
		}
	}
}
