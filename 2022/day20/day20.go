// Copyright 2022 Google LLC
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

// Package day20 solves AoC 2022 day 20.
package day20

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2022, 20, glue.IntSolver(solve))
}

func solve(file []int) ([]string, error) {
	p1 := decrypt(file, 1, 1, 50)
	p2 := decrypt(file, 811589153, 10, 50)
	return glue.Ints(p1, p2), nil
}

func decrypt(file []int, key, rounds, skipSize int) (plain int) {
	N := len(file)
	if N >= skipSize && N%skipSize != 0 {
		panic(fmt.Sprintf("%d %% %d != 0", N, skipSize))
	}

	nodes := make([]node, N)
	var zero *node
	for i, v := range file {
		if v == 0 {
			zero = &nodes[i]
		}
		p, n := (i+N-1)%N, (i+1)%N
		nodes[i] = node{value: key * v, prev: &nodes[p], next: &nodes[n]}
		if i%skipSize == 0 {
			ps, ns := (i+N-skipSize)%N, (i+skipSize)%N
			nodes[i].pskip, nodes[i].nskip = &nodes[ps], &nodes[ns]
		}
	}
	for round := 0; round < rounds; round++ {
		for i := range nodes {
			from := &nodes[i]
			v := from.value % (N - 1)
			if v >= N/2 {
				v -= N - 1
			} else if v <= -N/2 {
				v += N - 1
			}
			to := from
			var skipMoved *node
			if v > 0 {
				for v > 0 {
					if to.nskip != nil && to != skipMoved {
						toSkip := to.nskip
						to.pskip.nskip, to.nskip.pskip = to.next, to.next
						to.next.pskip, to.next.nskip = to.pskip, to.nskip
						to.pskip, to.nskip = nil, nil
						skipMoved = to.next
						if v >= skipSize {
							to = toSkip
							v -= skipSize
							continue
						}
					}
					to = to.next
					v--
				}
				if to.nskip != nil && to != skipMoved {
					to.pskip.nskip, to.nskip.pskip = from, from
					from.pskip, from.nskip = to.pskip, to.nskip
					to.pskip, to.nskip = nil, nil
				}
				from.prev.next, from.next.prev = from.next, from.prev                 // pick up
				to.next, from.prev, from.next, to.next.prev = from, to, to.next, from // put down
			} else if v < 0 {
				for v < 0 {
					if to.pskip != nil && to != skipMoved {
						toSkip := to.pskip
						to.pskip.nskip, to.nskip.pskip = to.prev, to.prev
						to.prev.pskip, to.prev.nskip = to.pskip, to.nskip
						to.pskip, to.nskip = nil, nil
						skipMoved = to.prev
						if v <= -skipSize {
							to = toSkip
							v += skipSize
							continue
						}
					}
					to = to.prev
					v++
				}
				if to.pskip != nil && to != skipMoved {
					to.pskip.nskip, to.nskip.pskip = from, from
					from.pskip, from.nskip = to.pskip, to.nskip
					to.pskip, to.nskip = nil, nil
				}
				from.prev.next, from.next.prev = from.next, from.prev                 // pick up
				to.prev.next, from.prev, from.next, to.prev = from, to.prev, to, from // put down
			} else {
				continue
			}
		}
	}
	for i, p := 0, zero; i < 3; i++ {
		for j := 1000; j > 0; {
			if p.nskip != nil && j >= skipSize {
				p = p.nskip
				j -= skipSize
			} else {
				p = p.next
				j--
			}
		}
		plain += p.value
	}
	return plain
}

type node struct {
	value        int
	prev, next   *node
	pskip, nskip *node
}

func decryptPlain(file []int, key int, rounds int) (plain int) {
	N := len(file)
	nodes := make([]nodePlain, N)
	var zero *nodePlain
	for i, v := range file {
		if v == 0 {
			zero = &nodes[i]
		}
		p, n := (i+N-1)%N, (i+1)%N
		nodes[i] = nodePlain{value: key * v, prev: &nodes[p], next: &nodes[n]}
	}
	for round := 0; round < rounds; round++ {
		for i := range nodes {
			from := &nodes[i]
			v := from.value % (N - 1)
			if v >= N/2 {
				v -= N - 1
			} else if v <= -N/2 {
				v += N - 1
			}
			to := from
			if v > 0 {
				for i := 0; i < v; i++ {
					to = to.next
				}
			} else if v < 0 {
				for i := 0; i >= v; i-- {
					to = to.prev
				}
			} else {
				continue
			}
			from.prev.next, from.next.prev = from.next, from.prev                 // pick up
			to.next, from.prev, from.next, to.next.prev = from, to, to.next, from // put down
		}
	}
	for i, p := 0, zero; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			p = p.next
		}
		plain += p.value
	}
	return plain
}

type nodePlain struct {
	value      int
	prev, next *nodePlain
}
