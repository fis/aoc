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
		got := g.TopoSort(false)
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v -> %v, want %v", test.edges, got, test.want)
		}
	}
}
