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
	"regexp"
	"regexp/syntax"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2020, 19, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	rules, err := parseRules(util.Lines(chunks[0]))
	if err != nil {
		return nil, err
	}

	prog1, err := regexp.Compile(rules.toFullRegexp(false).String())
	if err != nil {
		return nil, err
	}
	prog2, err := regexp.Compile(rules.toFullRegexp(true).String())
	if err != nil {
		return nil, err
	}

	part1, part2 := 0, 0
	for _, line := range util.Lines(chunks[1]) {
		if prog1.MatchString(line) {
			part1++
		}
		if prog2.MatchString(line) {
			part2++
		}
	}

	return glue.Ints(part1, part2), nil
}

type ruleSet struct {
	term   map[int][]byte
	unary  map[int][]int
	binary map[int][][2]int
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

func (rs ruleSet) toFullRegexp(magic bool) *syntax.Regexp {
	return &syntax.Regexp{
		Op:    syntax.OpConcat,
		Flags: syntax.Simple,
		Sub: []*syntax.Regexp{
			{Op: syntax.OpBeginText, Flags: syntax.Simple},
			rs.toRegexp(0, magic),
			{Op: syntax.OpEndText, Flags: syntax.Simple},
		},
	}
}

func (rs ruleSet) toRegexp(r int, magic bool) *syntax.Regexp {
	if magic && r == 8 {
		ex := &syntax.Regexp{Op: syntax.OpPlus, Sub0: [1]*syntax.Regexp{rs.toRegexp(42, true)}}
		ex.Sub = ex.Sub0[:1]
		return ex
	}
	if magic && r == 11 {
		r42, r31 := rs.toRegexp(42, true), rs.toRegexp(31, true)
		subs := make([]*syntax.Regexp, 40)
		for i := 0; i < 20; i++ {
			subs[i] = r42
			subs[20+i] = r31
		}
		ex := &syntax.Regexp{Op: syntax.OpAlternate}
		for n := 1; n < 20; n++ {
			ex.Sub = append(ex.Sub, &syntax.Regexp{Op: syntax.OpConcat, Sub: subs[20-n : 20+n]})
		}
		return ex
	}

	var alts []*syntax.Regexp
	for _, c := range rs.term[r] {
		ex := &syntax.Regexp{Op: syntax.OpLiteral, Rune0: [2]rune{rune(c), 0}}
		ex.Rune = ex.Rune0[:1]
		alts = append(alts, ex)
	}
	for _, u := range rs.unary[r] {
		alts = append(alts, rs.toRegexp(u, magic))
	}
	for _, b := range rs.binary[r] {
		ex := &syntax.Regexp{
			Op: syntax.OpConcat,
			Sub: []*syntax.Regexp{
				rs.toRegexp(b[0], magic),
				rs.toRegexp(b[1], magic),
			},
		}
		alts = append(alts, ex)
	}
	if len(alts) == 0 {
		panic(fmt.Sprintf("no rules: %d", r))
	}
	if len(alts) == 1 {
		return alts[0]
	}
	return &syntax.Regexp{Op: syntax.OpAlternate, Sub: alts}
}
