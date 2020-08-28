package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplePriorityQueue_Push(t *testing.T) {
	q := newSimplePriorityQueue()
	q.push(newNode("A"), 10)
	q.push(newNode("B"), 9)
	i := q.pop()
	assert.Equal(t, "B", i.name)
}

func TestSimplePriorityQueue_Pop(t *testing.T) {
	q := newSimplePriorityQueue()
	nodeA := newNode("A")
	q.push(nodeA, 10)
	q.push(newNode("B"), 9)
	q.update(nodeA, 8)
	i := q.pop()
	assert.Equal(t, "A", i.name)
}
