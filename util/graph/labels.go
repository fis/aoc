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

import "github.com/fis/aoc/util"

type graphLabels struct {
	labels   []string
	labelMap util.LabelMap
}

// Len returns the number of vertices in the graph.
func (l graphLabels) Len() int { return len(l.labels) }

// V returns the vertex index of the vertex with the given name.
func (l graphLabels) V(name string) (idx int, ok bool) {
	idx, ok = l.labelMap[name]
	return
}

// Label returns the label corresponding to a vertex index.
func (l graphLabels) Label(v int) string { return l.labels[v] }
