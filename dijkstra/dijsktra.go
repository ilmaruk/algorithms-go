package dijkstra

func dijkstra(g graph, s *node) dijkstraMap {
	// Initialise map
	dm := make(dijkstraMap)
	q := newSimplePriorityQueue()

	dm[s] = &mapItem{
		dist: 0,
	}
	q.push(s, 0)

	for _, n := range g.nodes {
		if n != s {
			dm[n] = &mapItem{
				dist: infinite,
			}
			q.push(n, infinite)
		}
	}

	// Main loop
	for true {
		u := q.pop()
		if u == nil {
			break
		}

		for _, e := range u.edges {
			v := e.dest
			alt := dm[u].dist + e.weight
			if alt < dm[v].dist {
				dm[v].dist = alt
				dm[v].prev = u
				q.update(v, alt)
			}
		}
	}

	return dm
}
