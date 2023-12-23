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

import (
	"fmt"
	"io"
)

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

// SuccV returns the idx'th successor of the given vertex.
func (g *Graph) SuccV(fromV, idx int) int {
	n := 0
	for i := 0; i < len(g.edges); i++ {
		if g.edges[fromV][i] {
			if n == idx {
				return i
			}
			n++
		}
	}
	return -1
}

// RangeSucc calls the callback for each of the given vertex's successors.
// It returns false if the callback ever returned false, true otherwise.
func (g *Graph) RangeSucc(from string, cb func(to string) bool) bool {
	return g.RangeSuccV(g.V(from), func(toV int) bool { return cb(g.vertNames[toV]) })
}

// RangeSuccV calls the callback for each of the given vertex's successors.
// It returns false if the callback ever returned false, true otherwise.
func (g *Graph) RangeSuccV(fromV int, cb func(toV int) bool) bool {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[fromV][i] {
			if !cb(i) {
				return false
			}
		}
	}
	return true
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

// PredV returns the idx'th predecessor of the given vertex.
func (g *Graph) PredV(toV, idx int) int {
	n := 0
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i][toV] {
			if n == idx {
				return i
			}
			n++
		}
	}
	return -1
}

// RangePred calls the callback for each of the given vertex's predecessors.
// It returns false if the callback ever returned false, true otherwise.
func (g *Graph) RangePred(to string, cb func(from string) bool) bool {
	return g.RangePredV(g.V(to), func(fromV int) bool { return cb(g.vertNames[fromV]) })
}

// RangePredV calls the callback for each of the given vertex's predecessors.
// It returns false if the callback ever returned false, true otherwise.
func (g *Graph) RangePredV(toV int, cb func(fromV int) bool) bool {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i][toV] {
			if !cb(i) {
				return false
			}
		}
	}
	return true
}

// TopoSort returns the graph's vertices in topological order (which must exist).
// If keepEdges is not set, all the graph edges will be removed.
func (g *Graph) TopoSort(keepEdges bool) []string {
	return g.Names(g.TopoSortV(keepEdges))
}

// TopoSortV returns the graph's vertices in topological order (which must exist).
// If keepEdges is not set, all the graph edges will be removed.
func (g *Graph) TopoSortV(keepEdges bool) []int {
	var edgeCopy [][]bool
	if keepEdges {
		edgeCopy = make([][]bool, len(g.edges))
		for i, e := range g.edges {
			edgeCopy[i] = make([]bool, len(e))
			copy(edgeCopy[i], e)
		}
	}

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

	if keepEdges {
		g.edges = edgeCopy
	}
	return order
}

// MakeUndirected ensures that all edges are bidirectional.
// If the graph has weights and an edge already exists in both directions, the weights sum up.
func (g *Graph) MakeUndirected() {
	N := len(g.edges)
	for u := 0; u < N-1; u++ {
		for v := u + 1; v < N; v++ {
			e := g.edges[u][v] || g.edges[v][u]
			g.edges[u][v], g.edges[v][u] = e, e
		}
	}
	if g.w != nil {
		for u := 0; u < N-1; u++ {
			for v := u + 1; v < N; v++ {
				w := g.w[u][v] + g.w[v][u]
				g.w[u][v], g.w[v][u] = w, w
			}
		}
	}
}

// WriteDOT writes the graph out in GraphViz format. The `nodeAttr` and `edgeAttr` callback
// functions are optional, and can be used to add extra attributes to the node. If the callback
// returns a "label" attribute, it takes precedence over the usual node name / edge weight.
func (g *Graph) WriteDOT(w io.Writer, name string, nodeAttr func(v int) map[string]string, edgeAttr func(fromV, toV int) map[string]string) (err error) {
	fmt.Fprintf(w, "digraph %s {\n", name)
	g.RangeV(func(v int) {
		var attrs map[string]string
		if nodeAttr != nil {
			attrs = nodeAttr(v)
		}
		fmt.Fprintf(w, "  n%d [", v)
		writeAttrs(w, attrs, "label", fmt.Sprintf(`"%s"`, g.Name(v)))
		fmt.Fprintf(w, "];\n")
	})
	g.RangeV(func(fromV int) {
		g.RangeSuccV(fromV, func(toV int) bool {
			var attrs map[string]string
			if edgeAttr != nil {
				attrs = edgeAttr(fromV, toV)
			}
			fmt.Fprintf(w, "  n%d -> n%d [", fromV, toV)
			if g.w != nil {
				writeAttrs(w, attrs, "label", fmt.Sprintf(`"%d"`, g.W(fromV, toV)))
			} else {
				writeAttrs(w, attrs)
			}
			fmt.Fprintf(w, "];\n")
			return true
		})
	})
	_, err = fmt.Fprintln(w, "}")
	return err
}

func writeAttrs(w io.Writer, attr map[string]string, xattr ...string) error {
	i := 0
	for k, v := range attr {
		if err := writeAttr(w, k, v, i); err != nil {
			return err
		}
		i++
	}
	for x := 0; x+1 < len(xattr); x += 2 {
		if _, ok := attr[xattr[x]]; ok {
			continue
		}
		if err := writeAttr(w, xattr[x], xattr[x+1], i); err != nil {
			return err
		}
		i++
	}
	return nil
}

func writeAttr(w io.Writer, k, v string, i int) error {
	if i > 0 {
		_, err := fmt.Fprint(w, ",")
		if err != nil {
			return err
		}
	}
	// TODO: better marshalling
	_, err := fmt.Fprintf(w, "%s=%s", k, v)
	if err != nil {
		return err
	}
	return nil
}
