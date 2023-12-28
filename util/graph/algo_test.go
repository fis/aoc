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
	"testing"

	"github.com/fis/aoc/util/fn"
	"github.com/google/go-cmp/cmp"
)

func TestTopoSort(t *testing.T) {
	tests := []struct {
		edges [][2]string
		want  []string
	}{
		{
			edges: [][2]string{{"a", "b"}, {"b", "c"}, {"c", "d"}},
			want:  []string{"a", "b", "c", "d"},
		},
		{
			edges: [][2]string{{"a", "b"}, {"a", "c"}, {"b", "c"}, {"b", "d"}, {"c", "d"}},
			want:  []string{"a", "b", "c", "d"},
		},
	}
	type topoSortable interface {
		TopoSort(keepEdges bool) []int
		Label(v int) string
	}
	graphTypes := []struct {
		name    string
		builder func(*Builder) topoSortable
	}{
		{"Dense", func(b *Builder) topoSortable { return b.DenseDigraph() }},
		{"SparseW", func(b *Builder) topoSortable { return b.SparseDigraphW() }},
	}
	for _, gt := range graphTypes {
		for _, test := range tests {
			b := NewBuilder()
			for _, e := range test.edges {
				b.AddEdgeL(e[0], e[1])
			}
			g := gt.builder(b)
			got := fn.Map(g.TopoSort(false), g.Label)
			if !cmp.Equal(got, test.want) {
				t.Errorf("%s(%v).TopoSort() = %v, want %v", gt.name, test.edges, got, test.want)
			}
		}
	}
}
