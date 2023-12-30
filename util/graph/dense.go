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

type Dense struct {
	graphLabels
	adj []bool
}

func (Dense) hasWeights() bool { return false }

type DenseW struct {
	graphLabels
	adj []int
}

func (DenseW) hasWeights() bool { return true }

func (g *Dense) e(u, v int) int  { return u*len(g.labels) + v }
func (g *DenseW) e(u, v int) int { return u*len(g.labels) + v }

// E returns true if the edge (u, v) exists in the graph.
func (g *Dense) E(u, v int) bool {
	return g.adj[g.e(u, v)]
}

// W returns 1 if an edge exists between the two vertices, or 0 otherwise.
func (g *Dense) W(u, v int) int {
	if g.E(u, v) {
		return 1
	} else {
		return 0
	}
}

// E returns true if the edge (u, v) exists and has a non-zero weight in the graph.
func (g *DenseW) E(u, v int) bool { return g.W(u, v) != 0 }

// W returns the weight of an edge between two vertices.
func (g *DenseW) W(u, v int) int {
	return g.adj[g.e(u, v)]
}

// Succ returns an iterator for the successors of vertex u.
func (g *Dense) Succ(u int) DenseIt {
	return g.Next(DenseIt{u: u, v: -1, du: 0, dv: 1})
}

// Succ returns an iterator for the successors of vertex u.
func (g *DenseW) Succ(u int) DenseIt {
	return g.Next(DenseIt{u: u, v: -1, du: 0, dv: 1})
}

// Pred returns an iterator for the predecessors of vertex v.
func (g *Dense) Pred(v int) DenseIt {
	return g.Next(DenseIt{u: -1, v: v, du: 1, dv: 0})
}

// Pred returns an iterator for the predecessors of vertex v.
func (g *DenseW) Pred(v int) DenseIt {
	return g.Next(DenseIt{u: -1, v: v, du: 1, dv: 0})
}

// Next moves an edge iterator (from Succ or Pred) one step further.
func (g *Dense) Next(it DenseIt) DenseIt {
	for {
		it.u, it.v = it.u+it.du, it.v+it.dv
		if it.u >= g.Len() || it.v >= g.Len() {
			return DenseIt{u: -1}
		}
		if g.adj[g.e(it.u, it.v)] {
			return it
		}
	}
}

// Next moves an edge iterator (from Succ or Pred) one step further.
func (g *DenseW) Next(it DenseIt) DenseIt {
	for {
		it.u, it.v = it.u+it.du, it.v+it.dv
		if it.u >= g.Len() || it.v >= g.Len() {
			return DenseIt{u: -1}
		}
		if g.adj[g.e(it.u, it.v)] != 0 {
			return it
		}
	}
}

// DenseIt is an iterator over the edges of a dense graph.
type DenseIt struct {
	u, v   int
	du, dv int
}

func (it DenseIt) Valid() bool    { return it.u >= 0 }
func (it DenseIt) At() (u, v int) { return it.u, it.v }
func (it DenseIt) Tail() int      { return it.u }
func (it DenseIt) Head() int      { return it.v }

// ForSucc iterates over the successors of u, returning true if the end was reached.
// Return false from the callback to stop short; that is then returned to the caller.
func (g *Dense) ForSucc(u int, cb func(v int) bool) bool {
	for v, e := 0, g.e(u, 0); v < g.Len(); v, e = v+1, e+1 {
		if g.adj[e] {
			if !cb(v) {
				return false
			}
		}
	}
	return true
}

// ForSuccW is like ForSucc, but the weight (always 1) will also be passed to the callback.
func (g *Dense) ForSuccW(u int, cb func(v, w int) bool) bool {
	for v, e := 0, g.e(u, 0); v < g.Len(); v, e = v+1, e+1 {
		if g.adj[e] {
			if !cb(v, 1) {
				return false
			}
		}
	}
	return true
}

// ForSucc iterates over the successors of u, returning true if the end was reached.
// Return false from the callback to stop short; that is then returned to the caller.
func (g *DenseW) ForSucc(u int, cb func(v int) bool) bool {
	for v, e := 0, g.e(u, 0); v < g.Len(); v, e = v+1, e+1 {
		if g.adj[e] != 0 {
			if !cb(v) {
				return false
			}
		}
	}
	return true
}

// ForSuccW is like ForSucc, but the weight will also be passed to the callback.
func (g *DenseW) ForSuccW(u int, cb func(v, w int) bool) bool {
	for v, e := 0, g.e(u, 0); v < g.Len(); v, e = v+1, e+1 {
		if w := g.adj[e]; w != 0 {
			if !cb(v, w) {
				return false
			}
		}
	}
	return true
}

// NumSucc returns the total number of successors of u. This is an O(|V|) operation.
func (g *Dense) NumSucc(u int) (n int) {
	for e, eN := g.e(u, 0), g.e(u+1, 0); e < eN; e++ {
		if g.adj[e] {
			n++
		}
	}
	return n
}

// NumPred returns the total number of successors of u. This is an O(|V|) operation.
func (g *Dense) NumPred(v int) (n int) {
	for e, eN := g.e(0, v), g.e(g.Len(), v); e < eN; e += g.Len() {
		if g.adj[e] {
			n++
		}
	}
	return n
}

// DelEdge removes the edge (u, v) from the graph. If the edge did not exist, the call does nothing.
func (g *Dense) DelEdge(u, v int) {
	g.adj[g.e(u, v)] = false
}
