package graph

import (
	"iter"
	"slices"
)

// Sparse is a compatibility alias for SparseG[string]
type Sparse = SparseG[string]

// SparseG represents a sparse graph with adjacency lists.
type SparseG[T comparable] struct {
	graphLabels[T]
	edges [][]int
}

func (SparseG[T]) hasWeights() bool { return false }

// SparseW is a compatibility alias for SparseWG[string].
type SparseW = SparseWG[string]

// SparseWG represents a sparse weighted graphs with adjacency and edge weight lists.
type SparseWG[T comparable] struct {
	SparseG[T]
	edgeW [][]int
}

func (SparseWG[T]) hasWeights() bool { return true }

// E returns true if the edge (u, v) exists in the graph.
func (g *SparseG[T]) E(u, v int) bool {
	return slices.Contains(g.edges[u], v)
}

// W returns 1 if an edge exists between the two vertices, or 0 otherwise.
func (g *SparseG[T]) W(u, v int) int {
	if g.E(u, v) {
		return 1
	} else {
		return 0
	}
}

// E returns true if the edge (u, v) exists and has a non-zero weight in the graph.
func (g *SparseWG[T]) E(u, v int) bool { return g.W(u, v) != 0 }

// W returns the weight of an edge between two vertices.
// Note that this involves a linear scan for the edge in the edge list; if iterating, get the value from the iterator instead.
func (g *SparseWG[T]) W(u, v int) int {
	i := slices.Index(g.edges[u], v)
	if i < 0 {
		return 0
	}
	return g.edgeW[u][i]
}

// Succ returns an iterator for the successors of vertex u.
func (g *SparseG[T]) Succ(u int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range g.edges[u] {
			if !yield(v) {
				return
			}
		}
	}
}

// SuccW returns an iterator for the successors of vertex u, together with their weights.
func (g *SparseWG[T]) SuccW(u int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, v := range g.edges[u] {
			if !yield(v, g.edgeW[u][i]) {
				return
			}
		}
	}
}

// NumSucc returns the number of successors of u. This is an O(1) operation.
func (g *SparseWG[T]) NumSucc(u int) int { return len(g.edges[u]) }

// SuccI returns the i'th successor of u. This is an O(1) operation.
func (g *SparseWG[T]) SuccI(u, i int) int { return g.edges[u][i] }

// PredI returns the i'th predecessor of u (ordered by vertex index).
// Note that this is an O(|V|+|E|) operation.
func (g *SparseWG[T]) PredI(v, i int) int {
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
