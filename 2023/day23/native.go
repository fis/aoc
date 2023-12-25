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
import "github.com/fis/aoc/util"

func unsafeLongestPath(g *util.Graph, startV, endV int) (longest int) {
	sg := make([]C.struct_vertex, g.Len())
	for u := range sg {
		g.RangeSuccV(u, func(v int) bool {
			if v != startV && v != endV {
				d := sg[u].degree
				sg[u].next[d].v, sg[u].next[d].w = C.uint32_t(v), C.uint32_t(g.W(u, v))
				sg[u].degree = d + 1
			}
			return true
		})
	}
	firstV, lastV := g.SuccV(startV, 0), g.PredV(endV, 0)
	wS, wE := g.W(startV, firstV), g.W(lastV, endV)
	return wS + int(C.brute_force(&sg[0], C.uint32_t(firstV), 0, C.uint32_t(lastV))) + wE
}
