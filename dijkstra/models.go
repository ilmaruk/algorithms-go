package dijkstra

type weight float64

// Graph

type graph struct {
	nodes []*node
}

func newGraph() graph {
	return graph{
		nodes: make([]*node, 0),
	}
}

func (g *graph) addNewNode(name string) *node {
	n := newNode(name)
	g.nodes = append(g.nodes, n)
	return n
}

// Node

type node struct {
	name  string
	edges []edge
}

func newNode(name string) *node {
	return &node{
		name:  name,
		edges: make([]edge, 0),
	}
}

func (n *node) addNewEdge(dest *node, weight weight) {
	e := edge{
		Destination: dest,
		Weight:      weight,
	}
	n.edges = append(n.edges, e)
}

// Edge

type edge struct {
	Destination *node
	Weight      weight
}

// Map

type dijkstraMap map[*node]mapItem

type mapItem struct {
	dist weight
	prev *node
}
