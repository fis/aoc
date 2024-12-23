// Copyright 2024 Google LLC
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

// Package day23 solves AoC 2024 day 23.
package day23

import (
	"slices"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	nodes, neigh := buildGraph(lines)
	part1 := countCliqueT3(nodes, neigh)
	part2 := maxClique(nodes, neigh)
	return []string{
		strconv.Itoa(part1),
		part2,
	}, nil
}

func countCliqueT3(nodes []nodeID, neigh *neighMap) (count int) {
	for _, a := range nodes {
		na := neigh[a]
		for bi, _ := na.find(a); bi < len(na)-1; bi++ {
			b := na[bi]
			for ci := bi + 1; ci < len(na); ci++ {
				c := na[ci]
				if !a.isT() && !b.isT() && !c.isT() {
					continue
				}
				if _, ok := neigh[b].find(c); !ok {
					continue
				}
				count++
			}
		}
	}
	return count
}

func maxClique(nodes []nodeID, neigh *neighMap) string {
	cliq, best := make(nodeList, 1, maxID), nodeList(nil)
	for _, n := range nodes {
		cliq[0] = n
		best = findMaxClique(cliq, best, neigh)
	}
	ids := make([]byte, 3*len(best)-1)
	for i, n := range best {
		if i > 0 {
			ids[3*i-1] = ','
		}
		ids[3*i] = 'a' + byte(n/26)
		ids[3*i+1] = 'a' + byte(n%26)
	}
	return string(ids)
}

func findMaxClique(cliq, best nodeList, neigh *neighMap) nodeList {
	last := cliq[len(cliq)-1]
	for _, c := range neigh[last].from(last) {
		if !neigh[c].hasAll(cliq[:len(cliq)-1]) {
			continue
		}
		newCliq := append(cliq, c)
		if len(newCliq) > len(best) {
			best = slices.Clone(newCliq)
		}
		best = findMaxClique(newCliq, best, neigh)
	}
	return best
}

type nodeID uint16

const maxID = 26 * 26

func (id nodeID) isT() bool {
	return id >= ('t'-'a')*26 && id < ('u'-'a')*26
}

type nodeList []nodeID
type neighMap [maxID]nodeList

func (nl nodeList) find(node nodeID) (pos int, ok bool) {
	n := len(nl)
	a, b := 0, n
	for a < b {
		h := (a + b) >> 1
		if nl[h] < node {
			a = h + 1
		} else {
			b = h
		}
	}
	return a, a < len(nl) && nl[a] == node
}

func (nl nodeList) from(node nodeID) nodeList {
	pos, _ := nl.find(node)
	return nl[pos:]
}

func (nl nodeList) hasAll(nodes nodeList) bool {
	for _, n := range nodes {
		pos, ok := nl.find(n)
		if !ok {
			return false
		}
		nl = nl[pos+1:]
	}
	return true
}

func buildGraph(lines []string) (nodes nodeList, neigh *neighMap) {
	neigh = new(neighMap)
	seen := util.MakeFixedBitmap1D(maxID)
	for _, line := range lines {
		a := 26*int(line[0]-'a') + int(line[1]-'a')
		b := 26*int(line[3]-'a') + int(line[4]-'a')
		if !seen.Get(a) {
			seen.Set(a)
			nodes = append(nodes, nodeID(a))
		}
		if !seen.Get(b) {
			seen.Set(b)
			nodes = append(nodes, nodeID(b))
		}
		neigh[a] = append(neigh[a], nodeID(b))
		neigh[b] = append(neigh[b], nodeID(a))
	}
	slices.Sort(nodes)
	for i := range maxID {
		slices.Sort(neigh[i])
	}
	return nodes, neigh
}
