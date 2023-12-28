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
	"github.com/fis/aoc/util"
)

// A Builder is used to incrementally grow a graph.
//
// Note that there is no checking for duplicate edges by default. The caller should pick one of:
//   - Ensure that the input only lists each edge once.
//   - Accept that the resulting graph will in fact be a multigraph.
//   - Use one of the dense graphs, which use an adjacency matrix (inherently robust for this).
//   - TODO: add methods for deduplicating
type Builder struct {
	verts util.LabelMap
	edges [][3]int
}

// NewBuilder returns a new, empty Builder.
func NewBuilder() *Builder {
	return &Builder{
		verts: make(util.LabelMap),
	}
}

// Len returns the number of vertices added to the builder so far.
func (b *Builder) Len() int { return len(b.verts) }

// V returns the vertex number corresponding to the given label, creating it if necessary.
func (b *Builder) V(label string) int { return b.verts.Get(label) }

// AddEdgeW records a new weighted edge between two vertices in the graph.
func (b *Builder) AddEdgeW(from, to, w int) { b.edges = append(b.edges, [3]int{from, to, w}) }

// AddEdge records a new edge between two vertices in the graph.
func (b *Builder) AddEdge(from, to int) { b.AddEdgeW(from, to, 1) }

// AddEdgeWL records a new weighted edge between two vertices (denoted by labels, created if necessary).
func (b *Builder) AddEdgeWL(from, to string, w int) { b.AddEdgeW(b.V(from), b.V(to), w) }

// AddEdgeL records a new edge between two vertices (denoted by labels, created if necessary).
func (b *Builder) AddEdgeL(from, to string) { b.AddEdgeW(b.V(from), b.V(to), 1) }

// TODO: make adjacency list thing a function

// DenseDigraph returns the contents of the builder as an unweighted dense digraph.
func (b *Builder) DenseDigraph() (g *Dense) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]bool, N*N)
	g = &Dense{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] = true
	}
	return g
}

// DenseDigraphW returns the contents of the builder as a weighted dense digraph.
//
// If multiple edges have been recorded between two vertices, their values sum up.
func (b *Builder) DenseDigraphW() (g *DenseW) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]int, N*N)
	g = &DenseW{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] += e[2]
	}
	return g
}

// DenseGraph returns the contents of the builder as an unweighted dense undirected graph.
func (b *Builder) DenseGraph() (g *Dense) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]bool, N*N)
	g = &Dense{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] = true
		adj[g.e(e[1], e[2])] = true
	}
	return g
}

// DenseGraphW returns the contents of the builder as a weighted dense undirected graph.
//
// If multiple edges have been recorded between two vertices (in any order), their values sum up.
func (b *Builder) DenseGraphW() (g *DenseW) {
	labels := b.labels()
	N := labels.Len()
	adj := make([]int, N*N)
	g = &DenseW{graphLabels: labels, adj: adj}
	for _, e := range b.edges {
		adj[g.e(e[0], e[1])] += e[2]
		adj[g.e(e[1], e[0])] += e[2]
	}
	return g
}

// SparseDigraph returns the contents of the builder as an unweighted sparse digraph.
func (b *Builder) SparseDigraph() (g *Sparse) {
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
	for i := 0; i < N; i++ {
		edges[i] = allEdges[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
	}
	for _, e := range b.edges {
		allEdges[edgePos[e[0]]] = e[1]
		edgePos[e[0]]++
	}
	return &Sparse{graphLabels: labels, edges: edges}
}

// SparseDigraphW returns the contents of the builder as an unweighted sparse digraph.
func (b *Builder) SparseDigraphW() (g *SparseW) {
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
	for i := 0; i < N; i++ {
		edges[i] = allEdges[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
		edgeW[i] = allEdgeW[edgePos[i]:edgePos[i+1]:edgePos[i+1]]
	}
	for _, e := range b.edges {
		allEdges[edgePos[e[0]]] = e[1]
		allEdgeW[edgePos[e[0]]] = e[2]
		edgePos[e[0]]++
	}
	return &SparseW{Sparse: Sparse{graphLabels: labels, edges: edges}, edgeW: edgeW}
}

func (b *Builder) labels() graphLabels {
	return graphLabels{labels: b.verts.Slice(), labelMap: b.verts}
}
