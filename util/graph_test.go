package util

import (
	"testing"

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
	for _, test := range tests {
		g := Graph{}
		for _, e := range test.edges {
			g.AddEdge(e[0], e[1])
		}
		got := g.TopoSort()
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v -> %v, want %v", test.edges, got, test.want)
		}
	}
}
