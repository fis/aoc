// Copyright 2020 Google LLC
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

// Package day19 solves AoC 2020 day 19.
package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2020, 19, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]int, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	rules, err := parseRules(util.Lines(chunks[0]))
	if err != nil {
		return nil, err
	}

	part1, part2 := 0, 0
	for _, line := range util.Lines(chunks[1]) {
		if rules.match(line) {
			part1++
		}
		if rules.matchSpecial(line) {
			part2++
		}
	}

	return []int{part1, part2}, nil
}

type rule struct {
	c  byte
	rs [][]int
}

type ruleSet map[int]*rule

func parseRules(lines []string) (rs ruleSet, err error) {
	rs = make(ruleSet, len(lines))

	for _, line := range lines {
		var i int
		if _, err := fmt.Sscanf(line, "%d:", &i); err != nil {
			return nil, fmt.Errorf("missing rule number: %s: %w", line, err)
		}
		rs[i] = &rule{}
		line = strings.TrimSpace(strings.TrimPrefix(line, fmt.Sprintf("%d:", i)))

		if len(line) == 3 && line[0] == '"' && line[2] == '"' {
			rs[i].c = line[1]
			continue
		}

		for _, alt := range strings.Split(line, " | ") {
			var r []int
			for _, num := range util.Words(alt) {
				n, err := strconv.Atoi(num)
				if err != nil {
					return nil, fmt.Errorf("expected number, got %q", num)
				}
				r = append(r, n)
			}
			rs[i].rs = append(rs[i].rs, r)
		}
	}

	return rs, nil
}

func (rs ruleSet) sort() (order []int) {
	type node struct {
		pred, succ map[int]struct{}
	}
	edges := make(map[int]*node, len(rs))
	avail := []int(nil)
	for i := range rs {
		edges[i] = &node{make(map[int]struct{}), make(map[int]struct{})}
	}
	for to, r := range rs {
		if r.c != 0 {
			avail = append(avail, to)
		}
		for _, alt := range r.rs {
			for _, from := range alt {
				edges[from].succ[to] = struct{}{}
				edges[to].pred[from] = struct{}{}
			}
		}
	}

	for len(avail) > 0 {
		n := avail[0]
		avail = avail[1:]
		order = append(order, n)
		for m := range edges[n].succ {
			delete(edges[m].pred, n)
			if len(edges[m].pred) == 0 {
				avail = append(avail, m)
			}
		}
		edges[n].succ = nil
	}
	if len(order) != len(rs) {
		panic("unsortable ruleset")
	}

	return order
}

func (rs ruleSet) match(input string) bool {
	valid := make([]map[int]map[int]struct{}, len(input))
	for at := 0; at < len(input); at++ {
		valid[at] = make(map[int]map[int]struct{})
	}

	for _, r := range rs.sort() {
		if c := rs[r].c; c != 0 {
			for i := 0; i < len(input); i++ {
				if input[i] == c {
					if valid[i][r] == nil {
						valid[i][r] = make(map[int]struct{})
					}
					valid[i][r][1] = struct{}{}
				}
			}
			continue
		}

		for at := 0; at < len(input); at++ {
			for _, alt := range rs[r].rs {
				pos := []int{at}
				for _, s := range alt {
					var next []int
					for _, p := range pos {
						if p >= len(input) {
							continue
						}
						for l := range valid[p][s] {
							next = append(next, p+l)
						}
					}
					pos = next
					if len(pos) == 0 {
						break
					}
				}
				for _, end := range pos {
					if valid[at][r] == nil {
						valid[at][r] = make(map[int]struct{})
					}
					valid[at][r][end-at] = struct{}{}
				}
			}
		}
	}

	_, ok := valid[0][0][len(input)]
	return ok
}

func (rs ruleSet) matchSpecial(input string) bool {
	valid := make([]map[int]map[int]struct{}, len(input))
	for at := 0; at < len(input); at++ {
		valid[at] = make(map[int]map[int]struct{})
		valid[at][8] = make(map[int]struct{})
		valid[at][11] = make(map[int]struct{})
	}

	for _, r := range rs.sort() {
		if c := rs[r].c; c != 0 {
			for i := 0; i < len(input); i++ {
				if input[i] == c {
					if valid[i][r] == nil {
						valid[i][r] = make(map[int]struct{})
					}
					valid[i][r][1] = struct{}{}
				}
			}
			continue
		}

		if r == 8 {
			for at := 0; at < len(input); at++ {
				pos := at
				for pos < len(input) && len(valid[pos][42]) == 1 {
					for l := range valid[pos][42] {
						pos += l
					}
					valid[at][8][pos-at] = struct{}{}
				}
			}
			continue
		} else if r == 11 {
			for at := 0; at < len(input); at++ {
				pos, c := at, 0
				for pos < len(input) && len(valid[pos][42]) == 1 {
					for l := range valid[pos][42] {
						pos += l
					}
					c++
					rp, rc := pos, 0
					for rc < c && rp < len(input) && len(valid[rp][31]) == 1 {
						for l := range valid[rp][31] {
							rp += l
						}
						rc++
					}
					if rc == c {
						valid[at][11][rp-at] = struct{}{}
					}
				}
			}
			continue
		}

		for at := 0; at < len(input); at++ {
			for _, alt := range rs[r].rs {
				pos := []int{at}
				for _, s := range alt {
					var next []int
					for _, p := range pos {
						if p >= len(input) {
							continue
						}
						for l := range valid[p][s] {
							next = append(next, p+l)
						}
					}
					pos = next
					if len(pos) == 0 {
						break
					}
				}
				for _, end := range pos {
					if valid[at][r] == nil {
						valid[at][r] = make(map[int]struct{})
					}
					valid[at][r][end-at] = struct{}{}
				}
			}
			if (r == 31 || r == 42) && len(valid[at][r]) > 1 {
				panic("TODO: rules 31/42 can match at multiple lengths")
			}
		}
	}

	_, ok := valid[0][0][len(input)]
	return ok
}
