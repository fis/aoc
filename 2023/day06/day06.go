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

// Package day06 solves AoC 2023 day 6.
package day06

import (
	"fmt"
	"math"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	times, dist, err := parseRecords(lines)
	if err != nil {
		return nil, fmt.Errorf("invalid race log: %w", err)
	}
	p1 := countAllWins(times, dist)
	p2 := countWins(joinDigits(times), joinDigits(dist))
	return glue.Ints(p1, p2), nil
}

func countAllWins(times, dist []int) (prod int) {
	prod = 1
	for i := range times {
		prod *= countWins(times[i], dist[i])
	}
	return prod
}

func countWins(time, dist int) int {
	t, d := float64(time), float64(dist)
	q := math.Sqrt(t*t - 4*d)
	hMin, hMax := int(math.Floor((t-q)/2)), int(math.Ceil((t+q)/2))
	return hMax - hMin - 1
}

// N.B. not used for the Go solution, but might be handy for the Z80 one.
func countWinsBinarySearch(time, dist int) int {
	// find low bound for winning solutions
	minL, minH := 0, time/2
	for minL+1 < minH {
		mid := minL + (minH-minL)/2
		if (time-mid)*mid > dist {
			minH = mid
		} else {
			minL = mid
		}
	}
	// find high bound for winning solutions
	maxL, maxH := time/2, time
	for maxL+1 < maxH {
		mid := maxL + (maxH-maxL)/2
		if (time-mid)*mid > dist {
			maxL = mid
		} else {
			maxH = mid
		}
	}
	return maxL - minL
}

func joinDigits(s []int) (joined int) {
	for _, n := range s {
		d := 10
		for d <= n {
			d *= 10
		}
		for d >= 10 {
			d /= 10
			joined = 10*joined + n/d%10
		}
	}
	return joined
}

func parseRecords(lines []string) (times, dist []int, err error) {
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("expected 2 lines, got %d", len(lines))
	}
	if !strings.HasPrefix(lines[0], "Time:") {
		return nil, nil, fmt.Errorf("missing \"Time:\" prefix in line: %s", lines[0])
	}
	if !strings.HasPrefix(lines[1], "Distance:") {
		return nil, nil, fmt.Errorf("missing \"Time:\" prefix in line: %s", lines[0])
	}
	times = util.Ints(lines[0][5:])
	dist = util.Ints(lines[1][9:])
	if len(times) != len(dist) {
		return nil, nil, fmt.Errorf("mismatched data: %d != %d", len(times), len(dist))
	}
	if len(times) == 0 {
		return nil, nil, fmt.Errorf("no data")
	}
	return times, dist, nil
}
