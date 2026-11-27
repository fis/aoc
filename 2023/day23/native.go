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
		for v, w := range g.SuccW(u) {
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
