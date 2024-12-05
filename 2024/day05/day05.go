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

// Package day05 solves AoC 2024 day 5.
package day05

import (
	"fmt"
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/iters"
)

func init() {
	glue.RegisterSolver(2024, 5, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	rules := parseRules(chunks[0])
	updates := parseUpdates(chunks[1])
	part1, part2 := checkUpdates(updates, rules)
	return glue.Ints(part1, part2), nil
}

func checkUpdates(updates, rules [][]int) (correct, incorrect int) {
	for _, update := range updates {
		if isOrdered(update, rules) {
			correct += update[len(update)/2]
		} else {
			fix(update, rules)
			incorrect += update[len(update)/2]
		}
	}
	return correct, incorrect
}

func isOrdered(update []int, rules [][]int) bool {
	for i := range len(update) - 1 {
		if rules[update[i]][update[i+1]] > 0 {
			return false
		}
	}
	return true
}

func fix(update []int, rules [][]int) {
	slices.SortFunc(update, func(a, b int) int {
		return rules[a][b]
	})
}

func parseRules(text string) (rules [][]int) {
	ruleData := util.Ints(text)
	N := fn.Max(ruleData) + 1
	rules = make([][]int, N)
	for i := range N {
		rules[i] = make([]int, N)
	}
	for i := 0; i < len(ruleData); i += 2 {
		a, b := ruleData[i], ruleData[i+1]
		rules[a][b] = -1
		rules[b][a] = 1
	}
	return rules
}

func parseUpdates(text string) (updates [][]int) {
	for update := range iters.Lines(text) {
		updates = append(updates, util.Ints(update))
	}
	return updates
}
