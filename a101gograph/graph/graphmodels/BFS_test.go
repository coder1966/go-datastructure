package graphmodels

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	g := Graph{}
	n1, n2, n3, n4, n5 := NewNode(1), NewNode(2), NewNode(3), NewNode(4), NewNode(5)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n5)
	g.AddEdge(n2, n3)
	g.AddEdge(n2, n4)
	g.AddEdge(n2, n5)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.BFS(func(node *Node) {
		fmt.Printf("[Current Traverse Node]: %v\n", node)
	})
}
