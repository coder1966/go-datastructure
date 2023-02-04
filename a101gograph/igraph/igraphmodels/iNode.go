package igraphmodels

/*
https://segmentfault.com/a/1190000039295001
INode: 顶点接口
IGraphVisitor: 图的遍历器接口
tNode: 顶点的实现
iNodeQueue: 候选顶点队列. 候选顶点的选择方式不同, 决定了是深度优先还是广度优先.
tLIFOQueue: LIFO堆栈, 实现iNodeQueue接口
tFIFOQeuue: FIFO队列, 实现iNodeQueue接口
tGraphVisitor: 遍历器, 实现IGraphVisitor接口,
	从指定顶点出发, 访问所有可到达的顶点, 然后返回顶点数组
	使用哈希表记录已访问顶点, 防止重复访问
*/

// INode 顶点的接口
type INode interface {
	ID() string         // 要有ID这个方法
	Append(child INode) // 增加一个顶点
	Children() []INode  // 孩子，本顶点连向的其他顶点
}
