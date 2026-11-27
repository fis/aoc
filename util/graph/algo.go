package graph

// TopoSort returns a topological sort order for the graph.
func (g *DenseG[T]) TopoSort(keepEdges bool) []int {
	var adjCopy []bool
	if keepEdges {
		adjCopy = make([]bool, len(g.adj))
		copy(adjCopy, g.adj)
	}

	var stack []int
	for v := range g.Len() {
		if g.NumPred(v) == 0 {
			stack = append(stack, v)
		}
	}
	var order []int
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for v := range g.Succ(u) {
			g.DelEdge(u, v)
			if g.NumPred(v) == 0 {
				stack = append(stack, v)
			}
		}
	}

	if keepEdges {
		g.adj = adjCopy
	}
	return order
}

// TopoSort returns a topological sort order for the graph.
func (g *SparseWG[T]) TopoSort(keepEdges bool) []int {
	var copyEdges, copyEdgeW [][]int
	if keepEdges {
		copyEdges, copyEdgeW = make([][]int, len(g.edges)), make([][]int, len(g.edges))
		copy(copyEdges, g.edges)
		copy(copyEdgeW, g.edgeW)
	}

	np := make([]int, g.Len())
	for _, e := range g.edges {
		for _, v := range e {
			np[v]++
		}
	}
	var stack []int
	for v, n := range np {
		if n == 0 {
			stack = append(stack, v)
		}
	}
	var order []int
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		vs := g.edges[u]
		g.edges[u], g.edgeW[u] = nil, nil
		for _, v := range vs {
			np[v]--
			if np[v] == 0 {
				stack = append(stack, v)
			}
		}
	}

	if keepEdges {
		g.edges, g.edgeW = copyEdges, copyEdgeW
	}
	return order
}

// N.B. this wasn't actually needed for toposort (since it only drops entire edge lists), but could come useful some day

func (g *SparseWG[T]) deepEdgeCopy() (edges [][]int, edgeW [][]int) {
	totalEdges := 0
	for _, e := range g.edges {
		totalEdges += len(e)
	}
	allEdges, allEdgeW := make([]int, totalEdges), make([]int, totalEdges)
	edges, edgeW = make([][]int, len(g.edges)), make([][]int, len(g.edges))
	at := 0
	for i, e := range g.edges {
		n := len(e)
		edges[i] = allEdges[at : at+n : at+n]
		edgeW[i] = allEdgeW[at : at+n : at+n]
		copy(edges[i], g.edges[i])
		copy(edgeW[i], g.edgeW[i])
		at += n
	}
	return edges, edgeW
}
