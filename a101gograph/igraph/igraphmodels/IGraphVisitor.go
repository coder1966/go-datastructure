package igraphmodels

// IGraphVisitor 图的遍历器接口
type IGraphVisitor interface {
	// Visit 要有遍历这个方法
	Visit(root INode, Policy VisitPolicy) []INode
}

type VisitPolicy int // 遍历的模式

const DFSPolicy VisitPolicy = 1 // 深度优先搜索
const BFSPolicy VisitPolicy = 2 // 广度优先
