package dijkstra

type weight float64

// Graph

type graph struct {
	Nodes []*node
}

func newGraph() graph {
	return graph{
		Nodes: make([]*node, 0),
	}
}

func (g *graph) addNewNode(name string) *node {
	n := newNode(name)
	g.Nodes = append(g.Nodes, n)
	return n
}

// Node

type node struct {
	Name  string
	Edges []edge
}

func newNode(name string) *node {
	return &node{
		Name:  name,
		Edges: make([]edge, 0),
	}
}

func (n *node) addNewEdge(dest *node, weight weight) {
	e := edge{
		Destination: dest,
		Weight:      weight,
	}
	n.Edges = append(n.Edges, e)
}

// Edge

type edge struct {
	Destination *node
	Weight      weight
}

// Map

type dijkstraMap map[*node]mapItem

type mapItem struct {
	Dist weight
	Prev *node
}
