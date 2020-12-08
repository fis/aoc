// Copyright 2019 Google LLC
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

package util

// Graph models a dense directed graph of string-labeled nodes.
type Graph struct {
	verts     map[string]int
	vertNames []string
	edges     [][]bool
	w         [][]int
}

// Len returns the number of vertices in the graph.
func (g *Graph) Len() int {
	return len(g.edges)
}

// V returns the vertex number of the vertex with the given name.
func (g *Graph) V(name string) int {
	if g.verts == nil {
		g.verts = make(map[string]int)
	}
	if v, ok := g.verts[name]; ok {
		return v
	}
	v := len(g.verts)
	g.verts[name] = v
	g.vertNames = append(g.vertNames, name)
	for i := range g.edges {
		g.edges[i] = append(g.edges[i], false)
	}
	g.edges = append(g.edges, make([]bool, len(g.verts)))
	if g.w != nil {
		for i := range g.w {
			g.w[i] = append(g.w[i], 0)
		}
		g.w = append(g.w, make([]int, len(g.verts)))
	}
	return v
}

// Name returns the name of the vertex with the given number.
func (g *Graph) Name(v int) string {
	return g.vertNames[v]
}

// Names converts a list of vertex numbers to the corresponding names.
func (g *Graph) Names(vs []int) []string {
	names := make([]string, len(vs))
	for i, v := range vs {
		names[i] = g.vertNames[v]
	}
	return names
}

// AddEdge adds an edge between from and to (if it didn't exist already), creating the vertices if necessary.
func (g *Graph) AddEdge(from, to string) {
	g.AddEdgeV(g.V(from), g.V(to))
}

// AddEdgeV adds an edge between from and to (if it didn't exist already).
func (g *Graph) AddEdgeV(fromV, toV int) {
	g.edges[fromV][toV] = true
}

// AddEdgeW adds a weighted edge between from and to, creating the vertices if necessary.
// If an edge already existed, its weight is updated.
func (g *Graph) AddEdgeW(from, to string, w int) {
	g.AddEdgeWV(g.V(from), g.V(to), w)
}

// AddEdgeWV adds a weighted edge between from and to. If an edge already existed, its weight is updated.
func (g *Graph) AddEdgeWV(fromV, toV, w int) {
	g.AddEdgeV(fromV, toV)
	if g.w == nil {
		g.w = make([][]int, len(g.verts))
		for i := range g.w {
			g.w[i] = make([]int, len(g.verts))
		}
	}
	g.w[fromV][toV] = w
}

// DelEdge removes an edge between from and to (if it existed), creating the vertices if necessary.
func (g *Graph) DelEdge(from, to string) {
	g.DelEdgeV(g.V(from), g.V(to))
}

// DelEdgeV removes an edge between from and to (if it existed).
func (g *Graph) DelEdgeV(fromV, toV int) {
	g.edges[fromV][toV] = false
}

// W returns the weight of an edge between two vertices. The call is valid only if your graph does have weights.
func (g *Graph) W(fromV, toV int) int {
	return g.w[fromV][toV]
}

// Range calls the callback for each of the graph's vertex names.
func (g *Graph) Range(cb func(name string)) {
	g.RangeV(func(v int) {
		cb(g.vertNames[v])
	})
}

// RangeV calls the callback for each of the graph's vertex numbers.
func (g *Graph) RangeV(cb func(v int)) {
	for v := 0; v < len(g.edges); v++ {
		cb(v)
	}
}

// NumSucc returns the number of successors of the given vertex.
func (g *Graph) NumSucc(from string) int {
	return g.NumSuccV(g.V(from))
}

// NumSuccV returns the number of successors of the given vertex.
func (g *Graph) NumSuccV(fromV int) int {
	n := 0
	for i := 0; i < len(g.edges); i++ {
		if g.edges[fromV][i] {
			n++
		}
	}
	return n
}

// RangeSucc calls the callback for each of the given vertex's successors.
func (g *Graph) RangeSucc(from string, cb func(to string) bool) {
	g.RangeSuccV(g.V(from), func(toV int) bool { return cb(g.vertNames[toV]) })
}

// RangeSuccV calls the callback for each of the given vertex's successors.
func (g *Graph) RangeSuccV(fromV int, cb func(toV int) bool) {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[fromV][i] {
			if !cb(i) {
				break
			}
		}
	}
}

// NumPred returns the number of predecessors of the given vertex.
func (g *Graph) NumPred(to string) int {
	return g.NumPredV(g.V(to))
}

// NumPredV returns the number of predecessors of the given vertex.
func (g *Graph) NumPredV(toV int) int {
	n := 0
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i][toV] {
			n++
		}
	}
	return n
}

// RangePred calls the callback for each of the given vertex's predecessors.
func (g *Graph) RangePred(to string, cb func(from string) bool) {
	g.RangePredV(g.V(to), func(fromV int) bool { return cb(g.vertNames[fromV]) })
}

// RangePredV calls the callback for each of the given vertex's predecessors.
func (g *Graph) RangePredV(toV int, cb func(fromV int) bool) {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i][toV] {
			if !cb(i) {
				break
			}
		}
	}
}

// TopoSort returns the graph's vertices in topological order (which must exist).
func (g *Graph) TopoSort() []string {
	return g.Names(g.TopoSortV())
}

// TopoSortV returns the graph's vertices in topological order (which must exist).
func (g *Graph) TopoSortV() []int {
	var stack []int
	g.RangeV(func(v int) {
		if g.NumPredV(v) == 0 {
			stack = append(stack, v)
		}
	})
	var order []int
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, n)
		g.RangeSuccV(n, func(m int) bool {
			g.DelEdgeV(n, m)
			if g.NumPredV(m) == 0 {
				stack = append(stack, m)
			}
			return true
		})
	}
	return order
}
