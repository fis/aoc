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

import (
	"slices"
)

type Sparse struct {
	graphLabels
	edges [][]int
}

func (Sparse) hasWeights() bool { return false }

type SparseW struct {
	Sparse
	edgeW [][]int
}

func (SparseW) hasWeights() bool { return true }

// E returns true if the edge (u, v) exists in the graph.
func (g *Sparse) E(u, v int) bool {
	return slices.Contains(g.edges[u], v)
}

// W returns 1 if an edge exists between the two vertices, or 0 otherwise.
func (g *Sparse) W(u, v int) int {
	if g.E(u, v) {
		return 1
	} else {
		return 0
	}
}

// E returns true if the edge (u, v) exists and has a non-zero weight in the graph.
func (g *SparseW) E(u, v int) bool { return g.W(u, v) != 0 }

// W returns the weight of an edge between two vertices.
// Note that this involves a linear scan for the edge in the edge list; if iterating, get the value from the iterator instead.
func (g *SparseW) W(u, v int) int {
	i := slices.Index(g.edges[u], v)
	if i < 0 {
		return 0
	}
	return g.edgeW[u][i]
}

// Succ returns an iterator for the successors of vertex u.
func (g *Sparse) Succ(u int) SparseIt {
	if len(g.edges[u]) == 0 {
		return SparseIt{i: -1}
	}
	return SparseIt{u: u, v: g.edges[u][0], i: 0}
}

// Succ returns an iterator for the successors of vertex u.
func (g *SparseW) Succ(u int) SparseItW {
	if len(g.edges[u]) == 0 {
		return SparseItW{SparseIt: SparseIt{i: -1}}
	}
	return SparseItW{SparseIt: SparseIt{u: u, v: g.edges[u][0], i: 0}, w: g.edgeW[u][0]}
}

// Next moves an edge iterator (from Succ or Pred) one step further.
func (g *Sparse) Next(it SparseIt) SparseIt {
	i := it.i + 1
	if i == 0 || i == len(g.edges[it.u]) {
		return SparseIt{i: -1}
	}
	return SparseIt{u: it.u, v: g.edges[it.u][i], i: i}
}

// Next moves an edge iterator (from Succ or Pred) one step further.
func (g *SparseW) Next(it SparseItW) SparseItW {
	i := it.i + 1
	if i == 0 || i == len(g.edges[it.u]) {
		return SparseItW{SparseIt: SparseIt{i: -1}}
	}
	return SparseItW{SparseIt: SparseIt{u: it.u, v: g.edges[it.u][i], i: i}, w: g.edgeW[it.u][i]}
}

// SparseIt is an iterator over the edges of a sparse unweighted graph.
type SparseIt struct {
	u, v, i int
}

func (it SparseIt) Valid() bool    { return it.i >= 0 }
func (it SparseIt) At() (u, v int) { return it.u, it.v }
func (it SparseIt) Tail() int      { return it.u }
func (it SparseIt) Head() int      { return it.v }

// SparseWIt is an iterator over the edges of a sparse weighted graph.
type SparseItW struct {
	SparseIt
	w int
}

func (it SparseItW) W() int { return it.w }

// ForSucc iterates over the successors of u, returning true if the end was reached.
// Return false from the callback to stop short; that is then returned to the caller.
func (g *Sparse) ForSucc(u int, cb func(v int) bool) bool {
	for _, v := range g.edges[u] {
		if !cb(v) {
			return false
		}
	}
	return true
}

// ForSuccW is like ForSucc, but the weight (always 1) will also be passed to the callback.
func (g *Sparse) ForSuccW(u int, cb func(v, w int) bool) bool {
	for _, v := range g.edges[u] {
		if !cb(v, 1) {
			return false
		}
	}
	return true
}

// ForSuccW is like ForSucc, but the weight will also be passed to the callback.
func (g *SparseW) ForSuccW(u int, cb func(v, w int) bool) bool {
	for i, v := range g.edges[u] {
		if !cb(v, g.edgeW[u][i]) {
			return false
		}
	}
	return true
}

// NumSucc returns the number of successors of u. This is an O(1) operation.
func (g *SparseW) NumSucc(u int) int { return len(g.edges[u]) }

// SuccI returns the i'th successor of u. This is an O(1) operation.
func (g *SparseW) SuccI(u, i int) int { return g.edges[u][i] }

// PredI returns the i'th predecessor of u (ordered by vertex index).
// Note that this is an O(|V|+|E|) operation.
func (g *SparseW) PredI(v, i int) int {
	for u, e := range g.edges {
		for _, v2 := range e {
			if v2 == v {
				if i == 0 {
					return u
				}
				i--
			}
		}
	}
	return -1
}
