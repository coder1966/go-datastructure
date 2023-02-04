package digraphmodels

import (
	"errors"
	"sync"
)

type DiGraph struct {
	vertexes []*Vertex    // 顶点集
	lock     sync.RWMutex // 安全锁
}

// AddVertex 加顶点
func (d *DiGraph) AddVertex(vertex *Vertex) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == vertex.name {
			return errors.New("您想增加的边已经存在")
		}
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes = append(d.vertexes, vertex)
	return nil
}

func (d *DiGraph) RenameVertex(vertex *Vertex, name string) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == vertex.name {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes[i].name = name
	return nil
}

func (d *DiGraph) DeleteVertex(vertex *Vertex) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == vertex.name {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes = append(d.vertexes[0:i], d.vertexes[i+1:]...)
	return nil
}

func (d *DiGraph) AddEdge(tail, head *Vertex, weight int) error {
	return tail.AddEdge(head, weight)
}

func (d *DiGraph) ModifyEdge(tail, head *Vertex, weight int) error {
	return tail.ModifyEdge(head, weight) // 不用解释吧
}

func (d *DiGraph) DeleteEdge(tail, head *Vertex) error {
	return tail.DeleteEdge(head) // 不用解释吧
}

/***********************************************************/

// AddVertexByName 加顶点
func (d *DiGraph) AddVertexByName(name string) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == name {
			return errors.New("您想增加的边已经存在")
		}
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes = append(d.vertexes, &Vertex{name: name, edges: make([]*Edge, 0)})
	return nil
}

func (d *DiGraph) RenameVertexByName(oldName, name string) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == oldName {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes[i].name = name
	return nil
}

func (d *DiGraph) DeleteVertexByName(name string) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == name {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes = append(d.vertexes[0:i], d.vertexes[i+1:]...)
	return nil
}

func (d *DiGraph) AddEdgeByName(tailName, headName string, weight int) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == tailName {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	return d.vertexes[i].AddEdgeByName(headName, weight)
}

func (d *DiGraph) ModifyEdgeByName(tailName, headName string, weight int) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == tailName {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	return d.vertexes[i].ModifyEdgeByName(headName, weight)
}

func (d *DiGraph) DeleteEdgeByName(tailName, headName string) error {
	i := 0
	for ; i < len(d.vertexes); i++ {
		if d.vertexes[i].name == tailName {
			break
		}
	}
	if i >= len(d.vertexes) {
		return errors.New("没有找到您要的起点")
	}
	return d.vertexes[i].DeleteEdgeByName(headName)
}

/***********************************************************/

func (d *DiGraph) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.vertexes = make([]*Vertex, 0)
}

func (d *DiGraph) Empty() bool {
	return len(d.vertexes) == 0
}

// NewDiGraph 构造函数
func NewDiGraph() IDigraph {
	return &DiGraph{
		vertexes: make([]*Vertex, 0),
	}
}

//
// // 输出图
// func (g *Graph) String() {
// 	g.lock.Lock()
// 	defer g.lock.Unlock()
// 	str := ""
// 	for _, iNode := range g.nodes { // 遍历节点
// 		str += iNode.String() + " -> " // 领衔节点
// 		nexts := g.edges[*iNode]       // 本领衔节点所有可联通的，边
// 		for _, next := range nexts {   // 遍历本领衔节点的变
// 			str += next.String() + " " // 加上 通达 节点
// 		}
// 		str += "\n"
// 	}
//
// 	fmt.Println(str)
// }
//
// // 输出节点
// func (n *Node) String() string {
// 	return fmt.Sprintf("%v", n.value)
// }
