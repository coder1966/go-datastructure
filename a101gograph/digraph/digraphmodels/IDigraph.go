package digraphmodels

// IDigraph 有向（加权）图接口
type IDigraph interface {
	AddVertex(vertex *Vertex) error                  // 加顶点方法
	RenameVertex(vertex *Vertex, name string) error  // 改名字
	DeleteVertex(vertex *Vertex) error               // 加顶点方法
	AddEdge(tail, head *Vertex, weight int) error    // 加边方法
	ModifyEdge(tail, head *Vertex, weight int) error // 修改边权重方法
	DeleteEdge(tail, head *Vertex) error             // 减边方法

	AddVertexByName(name string) error                            // 加顶点方法
	RenameVertexByName(oldName, name string) error                // 改名字
	DeleteVertexByName(name string) error                         // 加顶点方法
	AddEdgeByName(tailName, headName string, weight int) error    // 加边方法
	ModifyEdgeByName(tailName, headName string, weight int) error // 修改边权重方法
	DeleteEdgeByName(tailName, headName string) error             // 减边方法

	Clear()      // 清空方法
	Empty() bool // 是否为空的方法
}

type IVertex interface {
	Rename(name string)                        // 改名字
	AddEdge(head *Vertex, weight int) error    // 加边方法
	ModifyEdge(head *Vertex, weight int) error // 修改边权重方法
	DeleteEdge(head *Vertex) error             // 减边方法

	AddEdgeByName(headName string, weight int) error    // 加边方法
	ModifyEdgeByName(headName string, weight int) error // 修改边权重方法
	DeleteEdgeByName(headName string) error             // 减边方法
}
