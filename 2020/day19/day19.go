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
	"bytes"
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

	part1 := 0
	rules.cache = nil
	for _, line := range util.Lines(chunks[1]) {
		if rules.matches(0, line, false) {
			part1++
		}
	}

	part2 := 0
	rules.cache = nil
	for _, line := range util.Lines(chunks[1]) {
		if rules.matches(0, line, true) {
			part2++
		}
	}

	return []int{part1, part2}, nil
}

type ruleSet struct {
	term   map[int][]byte
	unary  map[int][]int
	binary map[int][][2]int
	cache  ruleCache
}

func parseRules(lines []string) (rs ruleSet, err error) {
	rs.term = make(map[int][]byte)
	rs.unary = make(map[int][]int)
	rs.binary = make(map[int][][2]int)
	free := 1000

	for _, line := range lines {
		var r int
		if _, err := fmt.Sscanf(line, "%d:", &r); err != nil {
			return ruleSet{}, fmt.Errorf("missing rule number: %s: %w", line, err)
		}
		line = strings.TrimSpace(strings.TrimPrefix(line, fmt.Sprintf("%d:", r)))

		if len(line) == 3 && line[0] == '"' && line[2] == '"' {
			rs.term[r] = append(rs.term[r], line[1])
			continue
		}

		for _, alt := range strings.Split(line, " | ") {
			var nums []int
			for _, word := range util.Words(alt) {
				num, err := strconv.Atoi(word)
				if err != nil {
					return ruleSet{}, fmt.Errorf("expected number, got %q", num)
				}
				nums = append(nums, num)
			}
			switch len(nums) {
			case 1:
				rs.unary[r] = append(rs.unary[r], nums[0])
			case 2:
				rs.binary[r] = append(rs.binary[r], [2]int{nums[0], nums[1]})
			case 3:
				rs.binary[r] = append(rs.binary[r], [2]int{nums[0], free})
				rs.binary[free] = append(rs.binary[free], [2]int{nums[1], nums[2]})
				free++
			default:
				return ruleSet{}, fmt.Errorf("unimplemented: 0 or >3 symbols: %d: %s", r, line)
			}
		}
	}

	return rs, nil
}

func (rs *ruleSet) matches(r int, input string, magic bool) bool {
	if magic && r == 8 {
		return rs.magic8(input)
	} else if magic && r == 11 {
		return rs.magic11(input)
	}
	for _, alt := range rs.unary[r] {
		if rs.matches(alt, input, magic) {
			return rs.cache.put(r, input, true)
		}
	}
	if len(input) == 1 {
		return bytes.IndexByte(rs.term[r], input[0]) >= 0
	}
	if out, ok := rs.cache.get(r, input); ok {
		return out
	}
	for _, alt := range rs.binary[r] {
		for s := 1; s < len(input); s++ {
			if rs.matches(alt[0], input[:s], magic) && rs.matches(alt[1], input[s:], magic) {
				return rs.cache.put(r, input, true)
			}
		}
	}
	return rs.cache.put(r, input, false)
}

func (rs *ruleSet) magic8(input string) bool {
	if out, ok := rs.cache.get(8, input); ok {
		return out
	}
	for s := 1; s <= len(input); s++ {
		head, tail := input[:s], input[s:]
		if rs.matches(42, head, true) && (tail == "" || rs.magic8(tail)) {
			return rs.cache.put(8, input, true)
		}
	}
	return rs.cache.put(8, input, false)
}

func (rs *ruleSet) magic11(input string) bool {
	if out, ok := rs.cache.get(11, input); ok {
		return out
	}
	for s := 1; s < len(input); s++ {
		for t := len(input) - 1; t >= s; t-- {
			head, body, tail := input[:s], input[s:t], input[t:]
			if rs.matches(42, head, true) && rs.matches(31, tail, true) && (body == "" || rs.magic11(body)) {
				return rs.cache.put(11, input, true)
			}
		}
	}
	return rs.cache.put(11, input, false)
}

type ruleCache map[int]map[string]bool

func (c ruleCache) get(r int, input string) (out, ok bool) {
	out, ok = c[r][input]
	return out, ok
}

func (c *ruleCache) put(r int, input string, out bool) bool {
	if *c == nil {
		*c = make(ruleCache)
	}
	if (*c)[r] == nil {
		(*c)[r] = make(map[string]bool)
	}
	(*c)[r][input] = out
	return out
}
