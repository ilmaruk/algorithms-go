package dijkstra

type priorityQueue interface {
	push(n *node, w weight)
	pop() *node
	update(n *node, w weight)
}

type simplePriorityQueue struct {
	items  map[*node]weight
	minKey *node
}

var _ priorityQueue = &simplePriorityQueue{}

func newSimplePriorityQueue() simplePriorityQueue {
	return simplePriorityQueue{
		items: make(map[*node]weight),
	}
}

func (q *simplePriorityQueue) push(n *node, w weight) {
	q.items[n] = w
	if q.minKey == nil || w < q.items[q.minKey] {
		q.minKey = n
	}
}

func (q *simplePriorityQueue) pop() *node {
	if q.minKey == nil {
		return nil
	}

	n := q.minKey
	delete(q.items, q.minKey)

	q.setMinKey()

	return n
}

func (q *simplePriorityQueue) update(n *node, w weight) {
	q.push(n, w)
}

func (q *simplePriorityQueue) setMinKey() {
	q.minKey = nil
	for n, w := range q.items {
		if q.minKey == nil || w < q.items[q.minKey] {
			q.minKey = n
		}
	}
}
