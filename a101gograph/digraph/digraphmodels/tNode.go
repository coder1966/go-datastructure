package digraphmodels

import (
	"fmt"
	"strings"
)

// tNode 顶点的实现
type tNode struct {
	id       string  // 顶点ID，字符串型
	children []INode // 顶点的孩子们（可联通的其他顶点）
}

// ID 方法，返回顶点ID
func (t *tNode) ID() string {
	return t.id
}

// Append 方法 本顶点加一个可联通到的孩子
func (t *tNode) Append(child INode) {
	t.children = append(t.children, child)
}

// Children 方法，返回被顶点所有孩子
func (t *tNode) Children() []INode {
	return t.children
}

// NewNode 构造函数。创建顶点，返回类型是接口，return的是实现，逼迫实现一系列方法
func NewNode(id string) INode {
	return &tNode{
		id:       id,
		children: make([]INode, 0),
	}
}

// String 方法，把本节点ID+本节点所有孩子ID展示出来
func (t *tNode) String() string {
	items := make([]string, len(t.children))
	for i, it := range t.children {
		items[i] = it.ID()
	}
	return fmt.Sprintf("%v-[%s]", t.id, strings.Join(items, ","))
}
