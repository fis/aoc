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

// Package day07 solves AoC 2022 day 7.
package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 7, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if first := fn.Head(lines, "(EOF)"); first != "$ cd /" {
		return nil, fmt.Errorf("expecting '$ cd /', got %q", first)
	}
	sizes := parseListing(lines[1:])
	p1 := part1(sizes)
	p2 := part2(sizes)
	return glue.Ints(p1, p2), nil
}

func part1(sizes []int) int {
	return fn.Sum(fn.Filter(sizes, func(s int) bool { return s <= 100000 }))
}

func part2(sizes []int) int {
	rootSize := sizes[len(sizes)-1]
	free := 70000000 - rootSize
	needed := 30000000 - free
	return fn.Min(fn.Filter(sizes, func(s int) bool { return s >= needed }))
}

func parseListing(lines []string) (sizes []int) {
	stack, depth := []int{0}, 0
	for len(lines) > 0 || depth > 0 {
		line := "$ cd .."
		if len(lines) > 0 {
			line, lines = lines[0], lines[1:]
		}
		switch {
		case line == "$ cd ..":
			size := stack[depth]
			sizes = append(sizes, size)
			stack, depth = stack[:depth], depth-1
			stack[depth] += size
		case strings.HasPrefix(line, "$ cd "):
			stack, depth = append(stack, 0), depth+1
		case line[0] == '$':
			// ignore other commands
		default:
			sep := strings.IndexByte(line, ' ')
			if word := line[:sep]; word != "dir" {
				size, _ := strconv.Atoi(line[:sep])
				stack[depth] += size
			}
		}
	}
	sizes = append(sizes, stack[0])
	return sizes
}
