package graphmodels

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

type Graph struct {
	nodes []*Node          // 节点集
	edges map[Node][]*Node // 邻接表表示的无向图
	lock  sync.RWMutex     // 安全锁
}

// AddNode 增加节点
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

// AddEdge 增加边
func (g *Graph) AddEdge(u, v *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*u] = append(g.edges[*u], v) // 建立u-->v的边
	g.edges[*v] = append(g.edges[*v], u) // 无向图，双向边都要建立
}

// 输出图
func (g *Graph) String() {
	g.lock.Lock()
	defer g.lock.Unlock()
	str := ""
	for _, iNode := range g.nodes { // 遍历节点
		str += iNode.String() + " -> " // 领衔节点
		nexts := g.edges[*iNode]       // 本领衔节点所有可联通的，边
		for _, next := range nexts {   // 遍历本领衔节点的变
			str += next.String() + " " // 加上 通达 节点
		}
		str += "\n"
	}

	fmt.Println(str)
}

// 输出节点
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}
