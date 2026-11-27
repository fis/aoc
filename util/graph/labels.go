package graph

import "github.com/fis/aoc/util"

type graphLabels[T comparable] struct {
	labels   []T
	labelMap util.LabelMap[T]
}

// Len returns the number of vertices in the graph.
func (l graphLabels[T]) Len() int { return len(l.labels) }

// V returns the vertex index of the vertex with the given name.
func (l graphLabels[T]) V(name T) (idx int, ok bool) {
	idx, ok = l.labelMap[name]
	return
}

// Label returns the label corresponding to a vertex index.
func (l graphLabels[T]) Label(v int) T { return l.labels[v] }
