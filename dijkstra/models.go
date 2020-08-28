package dijkstra

import (
	"fmt"
	"math"
)

type weight float64

const infinite = math.MaxFloat64

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
		dest:   dest,
		weight: weight,
	}
	n.edges = append(n.edges, e)
}

// Edge

type edge struct {
	dest   *node
	weight weight
}

// Map

type dijkstraMap map[*node]*mapItem

func (m dijkstraMap) String() string {
	var s string
	for n, i := range m {
		prev := ""
		if i.prev != nil {
			prev = i.prev.name
		}
		s += fmt.Sprintf("%s:[%s|%.f]\n", n.name, prev, i.dist)
	}
	return s
}

type mapItem struct {
	dist weight
	prev *node
}
