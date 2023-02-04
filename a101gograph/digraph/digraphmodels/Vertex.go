package digraphmodels

import (
	"errors"
	"sync"
)

// Vertex 顶点的结构体
type Vertex struct {
	name  string       // 顶点名字
	edges []*Edge      // 本顶点出发的所有有向有权重边
	lock  sync.RWMutex // 安全锁
}

func (v *Vertex) Rename(name string) {
	v.name = name
}

func (v *Vertex) AddEdge(head *Vertex, weight int) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head == head {
			return errors.New("您想增加的边已经存在")
		}
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges = append(v.edges, &Edge{head: head, weight: weight})
	return nil
}

func (v *Vertex) ModifyEdge(head *Vertex, weight int) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head == head {
			break
		}
	}
	if i >= len(v.edges) {
		return errors.New("没有找到您要修改的边")
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges[i].weight = weight
	return nil
}

func (v *Vertex) DeleteEdge(head *Vertex) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head == head {
			break
		}
	}
	if i >= len(v.edges) {
		return errors.New("没有找到您要删除的边")
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges = append(v.edges[:i], v.edges[i+1:]...)
	return nil
}

/***********************************************************/

func (v *Vertex) AddEdgeByName(headName string, weight int) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head.name == headName {
			return errors.New("您想增加的边已经存在")
		}
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges = append(v.edges, &Edge{head: v.edges[i].head, weight: weight})
	return nil
}

func (v *Vertex) ModifyEdgeByName(headName string, weight int) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head.name == headName {
			break
		}
	}
	if i >= len(v.edges) {
		return errors.New("没有找到您要修改的边")
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges[i].weight = weight
	return nil
}

func (v *Vertex) DeleteEdgeByName(headName string) error {
	i := 0
	for ; i < len(v.edges); i++ {
		if v.edges[i].head.name == headName {
			break
		}
	}
	if i >= len(v.edges) {
		return errors.New("没有找到您要删除的边")
	}
	v.lock.Lock()
	defer v.lock.Unlock()
	v.edges = append(v.edges[:i], v.edges[i+1:]...)
	return nil
}

// Edge 边的结构体
type Edge struct {
	head   *Vertex // head 终点，弧头； tail 起点，弧尾
	weight int     // 权重
}

// NewVertex 构造函数
func NewVertex(name string) IVertex {
	return &Vertex{
		name:  name,
		edges: make([]*Edge, 0),
	}
}
