package graph

import "godatastructure/a101gograph/graph/graphmodels"

func DemoAdd() {
	g := graphmodels.Graph{}
	n1, n2, n3, n4, n5 := graphmodels.NewNode(1), graphmodels.NewNode(2), graphmodels.NewNode(3), graphmodels.NewNode(4), graphmodels.NewNode(5)
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

	g.String()
}
