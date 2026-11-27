package graph

import (
	"github.com/fis/aoc/util"
)

// Builder is a compatibility alias for BuilderG[string].
type Builder = BuilderG[string]

// BuilderG is used to incrementally grow a graph.
//
// Note that there is no checking for duplicate edges by default. The caller should pick one of:
//   - Ensure that the input only lists each edge once.
//   - Accept that the resulting graph will in fact be a multigraph.
//   - Use one of the dense graphs, which use an adjacency matrix (inherently robust for this).
//   - TODO: add methods for deduplicating
type BuilderG[T comparable] struct {
	verts util.LabelMap[T]
	edges [][3]int
}

// NewBuilder returns a new, empty Builder.
func NewBuilder() *Builder {
	return NewBuilderG[string]()
}

// NewBuilder returns a new, empty Builder.
func NewBuilderG[T comparable]() *BuilderG[T] {
	return &BuilderG[T]{
		verts: make(util.LabelMap[T]),
	}
}

// Len returns the number of vertices added to the builder so far.
func (b *BuilderG[T]) Len() int { return len(b.verts) }

// V returns the vertex number corresponding to the given label, creating it if necessary.
func (b *BuilderG[T]) V(label T) int { return b.verts.Get(label) }

// AddEdgeW records a new weighted edge between two vertices in the graph.
func (b *BuilderG[T]) AddEdgeW(from, to, w int) { b.edges = append(b.edges, [3]int{from, to, w}) }

// AddEdge records a new edge between two vertices in the graph.
func (b *BuilderG[T]) AddEdge(from, to int) { b.AddEdgeW(from, to, 1) }

// AddEdgeWL records a new weighted edge between two vertices (denoted by labels, created if necessary).
func (b *BuilderG[T]) AddEdgeWL(from, to T, w int) { b.AddEdgeW(b.V(from), b.V(to), w) }

// AddEdgeL records a new edge between two vertices (denoted by labels, created if necessary).
func (b *BuilderG[T]) AddEdgeL(from, to T) { b.AddEdgeW(b.V(from), b.V(to), 1) }

// TODO: make adjacency list thing a function

// DenseDigraph returns the contents of the builder as an unweighted dense digraph.
func (b *BuilderG[T]) DenseDigraph() (g *DenseG[T]) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]bool, N*N)
	g = &DenseG[T]{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] = true
	}
	return g
}

// DenseDigraphW returns the contents of the builder as a weighted dense digraph.
//
// If multiple edges have been recorded between two vertices, their values sum up.
func (b *BuilderG[T]) DenseDigraphW() (g *DenseWG[T]) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]int, N*N)
	g = &DenseWG[T]{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] += e[2]
	}
	return g
}

// DenseGraph returns the contents of the builder as an unweighted dense undirected graph.
func (b *BuilderG[T]) DenseGraph() (g *DenseG[T]) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]bool, N*N)
	g = &DenseG[T]{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] = true
		adj[g.e(e[1], e[2])] = true
	}
	return g
}

// DenseGraphW returns the contents of the builder as a weighted dense undirected graph.
//
// If multiple edges have been recorded between two vertices (in any order), their values sum up.
func (b *BuilderG[T]) DenseGraphW() (g *DenseWG[T]) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]int, N*N)
	g = &DenseWG[T]{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] += e[2]
		adj[g.e(e[1], e[0])] += e[2]
	}
	return g
}

// SparseDigraph returns the contents of the builder as an unweighted sparse digraph.
func (b *BuilderG[T]) SparseDigraph() (g *SparseG[T]) {
	labels := b.labels()
	N := labels.Len()
	edgePos := make([]int, N+1)
	for _, e := range b.edges {
		edgePos[e[0]]++
	}
	for i, at := 0, 0; i < N+1; i++ {
		edgePos[i], at = at, at+edgePos[i]
	}
	allEdges := make([]int, len(b.edges))
	edges := make([][]int, N)
	for i := range N {
		edges[i] = allEdges[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
	}
	for _, e := range b.edges {
		allEdges[edgePos[e[0]]] = e[1]
		edgePos[e[0]]++
	}
	return &SparseG[T]{graphLabels: labels, edges: edges}
}

// SparseDigraphW returns the contents of the builder as a weighted sparse digraph.
func (b *BuilderG[T]) SparseDigraphW() (g *SparseWG[T]) {
	labels := b.labels()
	N := labels.Len()
	edgePos := make([]int, N+1)
	for _, e := range b.edges {
		edgePos[e[0]]++
	}
	for i, at := 0, 0; i < N+1; i++ {
		edgePos[i], at = at, at+edgePos[i]
	}
	allEdges := make([]int, len(b.edges))
	allEdgeW := make([]int, len(b.edges))
	edges := make([][]int, N)
	edgeW := make([][]int, N)
	for i := range N {
		edges[i] = allEdges[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
		edgeW[i] = allEdgeW[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
	}
	for _, e := range b.edges {
		allEdges[edgePos[e[0]]] = e[1]
		allEdgeW[edgePos[e[0]]] = e[2]
		edgePos[e[0]]++
	}
	return &SparseWG[T]{SparseG: SparseG[T]{graphLabels: labels, edges: edges}, edgeW: edgeW}
}

func (b *BuilderG[T]) labels() graphLabels[T] {
	return graphLabels[T]{labels: b.verts.Slice(), labelMap: b.verts}
}
