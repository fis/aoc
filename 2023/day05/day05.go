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

// Package day05 solves AoC 2023 day 5.
package day05

import (
	"cmp"
	"slices"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 5, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	al := parseAlmanac(chunks)
	p1 := lowestSingle(al)
	p2 := lowestRanged(al)
	return glue.Ints(p1, p2), nil
}

func lowestSingle(al almanac) int {
	locs := make([]int, len(al.seeds))
	for i, seed := range al.seeds {
		locs[i] = al.mapSeed(seed)
	}
	return fn.Min(locs)
}

func lowestRanged(al almanac) int {
	locs := make([]rng, len(al.seeds)/2)
	for i := range locs {
		locs[i] = rng{start: al.seeds[2*i], size: al.seeds[2*i+1]}
	}
	return al.mapRanges(locs)
}

type almanac struct {
	seeds []int
	maps  [][]mapping
}

type mapping struct {
	dst  int
	src  int
	size int
}

func (m mapping) sortKey() int {
	return m.src
}

type rng struct {
	start int
	size  int
}

func (r rng) sortKey() int {
	return r.start
}

func (al almanac) mapSeed(n int) int {
	for _, maps := range al.maps {
		mi, found := slices.BinarySearchFunc(maps, n, cmpIdWithMapping)
		if found {
			m := maps[mi]
			n = m.dst + (n - m.src)
		}
	}
	return n
}

func (al almanac) mapRanges(rs []rng) int {
	mapped := []rng(nil)
	for layer, maps := range al.maps {
		for _, r := range rs {
			for r.size > 0 {
				// N.B. since both the ranges and maps are in order, this could be done without the binary search.
				// Initial tests suggests any performance difference is negligible.
				mi, found := slices.BinarySearchFunc(maps, r.start, cmpIdWithMapping)
				var start, size int
				if found {
					m := maps[mi]
					offset := r.start - m.src
					start, size = m.dst+offset, min(r.size, m.size-offset)
				} else if mi < len(maps) {
					next := maps[mi]
					start, size = r.start, min(r.size, next.src-r.start)
				} else {
					start, size = r.start, r.size
				}
				mapped = append(mapped, rng{
					start: start,
					size:  size,
				})
				r.start, r.size = r.start+size, r.size-size
			}
		}
		rs, mapped = mapped, rs[:0]
		// N.B. this step does actually merge some ranges, though it doesn't seem to really affect performance.
		if layer != len(al.maps)-1 {
			rs = mergeRanges(rs)
		}
	}
	return fn.MinF(rs, rng.sortKey)
}

func mergeRanges(rs []rng) (merged []rng) {
	util.SortBy(rs, rng.sortKey)
	merged = rs[:0]
	for i := range rs {
		if i+1 < len(rs) && rs[i].start+rs[i].size == rs[i+1].start {
			rs[i+1].start, rs[i+1].size = rs[i].start, rs[i].size+rs[i+1].size
		} else {
			merged = append(merged, rs[i])
		}
	}
	return merged
}

func cmpIdWithMapping(m mapping, n int) int {
	if n >= m.src && n < m.src+m.size {
		return 0
	}
	return cmp.Compare(m.src, n)
}

func parseAlmanac(chunks []string) almanac {
	al := almanac{}

	_, seeds, _ := strings.Cut(chunks[0], ": ")
	al.seeds = util.Ints(seeds)
	chunks = chunks[1:]

	al.maps = make([][]mapping, len(chunks))
	for i, chunk := range chunks {
		_, chunk, _ = strings.Cut(chunk, "\n")
		for len(chunk) > 0 {
			r, tail, _ := strings.Cut(chunk, "\n")
			ri := util.Ints(r)
			al.maps[i] = append(al.maps[i], mapping{dst: ri[0], src: ri[1], size: ri[2]})
			chunk = tail
		}
		util.SortBy(al.maps[i], mapping.sortKey)
	}

	return al
}
