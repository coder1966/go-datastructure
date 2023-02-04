package igraphmodels

/*
遍历器，实现IGraphVisitor 接口
从指定顶点出发, 访问所有可到达的顶点, 然后返回顶点数组
使用哈希表记录已访问顶点, 防止重复访问
*/
type tGraphVisitor struct {
}

// NewGraphVisitor 构造函数，返回类型是接口，return 的是接口的实现，逼迫接口的实现配上相应的方法
func NewGraphVisitor() IGraphVisitor {
	return &tGraphVisitor{}
}

// CreatQueue 创建队列
// @Policy 遍历的模式，深度优先、广度优先。返回类型是队列的接口，可以配对多种队列的实现。
func (t *tGraphVisitor) CreatQueue(Policy VisitPolicy) iNodeQueue {
	switch Policy {
	case BFSPolicy:
		return newFIFOQueue()
	case DFSPolicy:
		return NewLIFOQueue()
	default:
		panic("模式没有指定")
	}
}

// Visit 遍历，访问
// @root 入口，树根
// @policy 遍历模式，深度、广度
// @return 遍历的顶点的有序数组切片
func (t *tGraphVisitor) Visit(root INode, policy VisitPolicy) []INode {
	queue := t.CreatQueue(policy) // 队列创建，按指定模式
	queue.Push(root)              // 根顶点入列

	visited := make(map[string]bool, 0) // 创建“已经访问顶点map”，防止重复访问
	result := make([]INode, 0)          // 准备返回的结果
	for !queue.Empty() {                // 队列不空
		_, node := queue.Poll()       // 队列里取出一个顶点
		visited[node.ID()] = true     // map里，设定这个取出的顶点“访问过了”
		result = append(result, node) // 返回结果集加上这个顶点

		children := node.Children() // 本顶点的孩子们（可到达顶点）
		if children != nil {        // 有可达到顶点
			for _, it := range children { // 遍历孩子们
				ok, _ := visited[it.ID()] // map里有？（就是访问过）
				if ok {                   // 访问过
					continue // 短路，回去
				}
				queue.Push(it) // 新的顶点（没访问过），加入返回结果集
			}
		}
	}
	return result
}

// GraphVisitor 测试时候，这个是入口
var GraphVisitor = NewGraphVisitor()
