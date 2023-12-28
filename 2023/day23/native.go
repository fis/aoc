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

package day23

// #include "native.h"
import "C"
import (
	"github.com/fis/aoc/util/graph"
)

func unsafeLongestPath(g *graph.SparseW, startV, endV int) (longest int) {
	sg := make([]C.struct_vertex, g.Len())
	for u := range sg {
		if u == startV || u == endV {
			continue
		}
		for it := g.Succ(u); it.Valid(); it = g.Next(it) {
			v, w := it.Head(), it.W()
			if u == startV || u == endV {
				continue
			}
			for _, e := range [2][2]int{{u, v}, {v, u}} {
				d := sg[e[0]].degree
				sg[e[0]].next[d].v, sg[e[0]].next[d].w = C.uint32_t(e[1]), C.uint32_t(w)
				sg[e[0]].degree = d + 1
			}
		}
	}
	firstV, lastV := g.SuccI(startV, 0), g.PredI(endV, 0)
	wS, wE := g.W(startV, firstV), g.W(lastV, endV)
	return wS + int(C.brute_force(&sg[0], C.uint32_t(firstV), C.uint32_t(lastV))) + wE
}
