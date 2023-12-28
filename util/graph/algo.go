// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

// TopoSort returns a topological sort order for the graph.
func (g *Dense) TopoSort(keepEdges bool) []int {
	var adjCopy []bool
	if keepEdges {
		adjCopy = make([]bool, len(g.adj))
		copy(adjCopy, g.adj)
	}

	var stack []int
	for v := 0; v < g.Len(); v++ {
		if g.NumPred(v) == 0 {
			stack = append(stack, v)
		}
	}
	var order []int
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for it := g.Succ(u); it.Valid(); it = g.Next(it) {
			v := it.Head()
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
func (g *SparseW) TopoSort(keepEdges bool) []int {
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

func (g *SparseW) deepEdgeCopy() (edges [][]int, edgeW [][]int) {
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
