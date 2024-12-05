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

import "iter"

// Dense is a compatibility alias for DenseG[string].
type Dense = DenseG[string]

// DenseG represents a dense graph using a boolean adjacency matrix.
type DenseG[T comparable] struct {
	graphLabels[T]
	adj []bool
}

func (DenseG[T]) hasWeights() bool { return false }

// DenseW is a compatibility alias for DenseWG[string]
type DenseW = DenseWG[string]

// DenseWG represents a dense weighted graph using an integer adjacency matrix.
type DenseWG[T comparable] struct {
	graphLabels[T]
	adj []int
}

func (DenseWG[T]) hasWeights() bool { return true }

func (g *DenseG[T]) e(u, v int) int  { return u*len(g.labels) + v }
func (g *DenseWG[T]) e(u, v int) int { return u*len(g.labels) + v }

// E returns true if the edge (u, v) exists in the graph.
func (g *DenseG[T]) E(u, v int) bool {
	return g.adj[g.e(u, v)]
}

// W returns 1 if an edge exists between the two vertices, or 0 otherwise.
func (g *DenseG[T]) W(u, v int) int {
	if g.E(u, v) {
		return 1
	} else {
		return 0
	}
}

// E returns true if the edge (u, v) exists and has a non-zero weight in the graph.
func (g *DenseWG[T]) E(u, v int) bool { return g.W(u, v) != 0 }

// W returns the weight of an edge between two vertices.
func (g *DenseWG[T]) W(u, v int) int {
	return g.adj[g.e(u, v)]
}

// Succ returns an iterator for the successors of vertex u.
func (g *DenseG[T]) Succ(u int) iter.Seq[int] {
	return func(yield func(int) bool) {
		e, n := g.e(u, 0), g.Len()
		for v := range n {
			if g.adj[e] {
				if !yield(v) {
					return
				}
			}
			e++
		}
	}
}

// Succ returns an iterator for the successors of vertex u.
func (g *DenseWG[T]) Succ(u int) iter.Seq[int] {
	return func(yield func(int) bool) {
		e, n := g.e(u, 0), g.Len()
		for v := range n {
			if g.adj[e] != 0 {
				if !yield(v) {
					return
				}
			}
			e++
		}
	}
}

// SuccW returns an iterator for the successors of vertex u, together with their weights.
func (g *DenseWG[T]) SuccW(u int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		e, n := g.e(u, 0), g.Len()
		for v := range n {
			if w := g.adj[e]; w != 0 {
				if !yield(v, w) {
					return
				}
			}
			e++
		}
	}
}

// Pred returns an iterator for the predecessors of vertex v.
func (g *DenseG[T]) Pred(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		e, n := g.e(0, v), g.Len()
		for u := range n {
			if g.adj[e] {
				if !yield(u) {
					return
				}
			}
			e += n
		}
	}
}

// Pred returns an iterator for the predecessors of vertex v.
func (g *DenseWG[T]) Pred(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		e, n := g.e(0, v), g.Len()
		for u := range n {
			if g.adj[e] != 0 {
				if !yield(u) {
					return
				}
			}
			e += n
		}
	}
}

// NumSucc returns the total number of successors of u. This is an O(|V|) operation.
func (g *DenseG[T]) NumSucc(u int) (n int) {
	for e, eN := g.e(u, 0), g.e(u+1, 0); e < eN; e++ {
		if g.adj[e] {
			n++
		}
	}
	return n
}

// NumPred returns the total number of successors of u. This is an O(|V|) operation.
func (g *DenseG[T]) NumPred(v int) (n int) {
	for e, eN := g.e(0, v), g.e(g.Len(), v); e < eN; e += g.Len() {
		if g.adj[e] {
			n++
		}
	}
	return n
}

// DelEdge removes the edge (u, v) from the graph. If the edge did not exist, the call does nothing.
func (g *DenseG[T]) DelEdge(u, v int) {
	g.adj[g.e(u, v)] = false
}
