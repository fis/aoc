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

// Package day04 solves AoC 2018 day 4.
package day04

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(4, util.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	log := parseLog(lines)
	g1, m1 := strategy1(log)
	g2, m2 := strategy2(log)
	return []int{g1 * m1, g2 * m2}, nil
}

type sleepMask [60]bool

func parseLog(lines []string) map[int][]sleepMask {
	sort.Strings(lines)
	var (
		masks = make(map[int][]sleepMask)
		mask  *sleepMask
		start int
	)
	for _, line := range lines {
		action := line[19:]
		if action == "falls asleep" {
			start, _ = strconv.Atoi(line[15:17])
			continue
		}
		if action == "wakes up" {
			end, _ := strconv.Atoi(line[15:17])
			for m := start; m < end; m++ {
				mask[m] = true
			}
			continue
		}
		var id int
		if _, err := fmt.Sscanf(action, "Guard #%d begins shift", &id); err == nil {
			masks[id] = append(masks[id], sleepMask{})
			mask = &masks[id][len(masks[id])-1]
			continue
		}
	}
	return masks
}

func strategy1(log map[int][]sleepMask) (guard, minute int) {
	var max int
	for g, masks := range log {
		var (
			total int
			freq  [60]int
		)
		for _, mask := range masks {
			for m, s := range mask {
				if s {
					freq[m]++
					total++
				}
			}
		}
		if total > max {
			max = total
			var maxm, maxf int
			for m, f := range freq {
				if f > maxf {
					maxm, maxf = m, f
				}
			}
			guard, minute = g, maxm
		}
	}
	return guard, minute
}

func strategy2(log map[int][]sleepMask) (guard, minute int) {
	var bestM, bestG, bestF int
	for m := 0; m < 60; m++ {
		var maxG, maxF int
		for g, masks := range log {
			var f int
			for _, mask := range masks {
				if mask[m] {
					f++
				}
			}
			if f > maxF {
				maxG, maxF = g, f
			}
		}
		if maxF > bestF {
			bestM, bestG, bestF = m, maxG, maxF
		}
	}
	return bestG, bestM
}
