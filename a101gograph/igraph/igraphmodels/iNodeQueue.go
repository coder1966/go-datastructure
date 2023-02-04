package igraphmodels

// 候选顶点队列接口，候选顶点的选择方式不同，决定是深度优先还是广度优先
type iNodeQueue interface {
	Clear()              // 要包含清空方法
	Empty() bool         // 要包含 ？是否队列为空的方法
	Push(node INode)     // 压入队列一个顶点
	Poll() (bool, INode) // 获取并删除Queue中的第一个元素
}
