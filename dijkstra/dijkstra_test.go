package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	g := setUpGraph()
	m := dijkstra(g, g.nodes[0])
	ref := m[g.nodes[6]]
	assert.Equal(t, weight(7), ref.dist)
	assert.Equal(t, g.nodes[4], ref.prev)
}

func BenchmarkDijkstra(b *testing.B) {
	g := setUpGraph()
	for n := 0; n < b.N; n++ {
		dijkstra(g, g.nodes[0])
	}
}

func setUpGraph() graph {
	g := graph{}

	// Nodes
	nodeA := g.addNewNode("A")
	nodeB := g.addNewNode("B")
	nodeC := g.addNewNode("C")
	nodeD := g.addNewNode("D")
	nodeE := g.addNewNode("E")
	nodeF := g.addNewNode("F")
	nodeG := g.addNewNode("G")

	// Edges
	nodeA.addNewEdge(nodeB, 2)
	nodeA.addNewEdge(nodeC, 4)
	nodeB.addNewEdge(nodeC, 1)
	nodeB.addNewEdge(nodeD, 5)
	nodeB.addNewEdge(nodeE, 3)
	nodeC.addNewEdge(nodeD, 2)
	nodeC.addNewEdge(nodeF, 3)
	nodeD.addNewEdge(nodeF, 1)
	nodeD.addNewEdge(nodeG, 5)
	nodeE.addNewEdge(nodeG, 2)
	nodeF.addNewEdge(nodeE, 3)
	nodeF.addNewEdge(nodeG, 4)

	return g
}
