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

// Package day21 solves AoC 2020 day 21.
package day21

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2020, 21, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	lines, err := util.ScanAll(r, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	labels, err := parseInput(lines)
	if err != nil {
		return nil, err
	}
	p1, p2 := analyze(labels)
	return []string{strconv.Itoa(p1), p2}, nil
}

type foodLabel struct {
	ingredients []string
	allergens   []string
}

func parseInput(lines []string) (out []foodLabel, err error) {
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("bad label: %s", line)
		}
		is := util.Words(parts[0])
		as := strings.Split(strings.TrimSuffix(parts[1], ")"), ", ")
		sort.Strings(is)
		out = append(out, foodLabel{ingredients: is, allergens: as})
	}
	return out, nil
}

func analyze(labels []foodLabel) (int, string) {
	possible := map[string][]string{}
	for _, l := range labels {
		for _, a := range l.allergens {
			if possible[a] == nil {
				possible[a] = l.ingredients
			} else {
				possible[a] = intersect(possible[a], l.ingredients)
			}
		}
	}

	type task struct {
		allergen string
		choices  []string
	}
	tasks := []task(nil)
	for a, is := range possible {
		tasks = append(tasks, task{a, is})
	}

	a2i, i2a := map[string]string{}, map[string]string{}
	for len(tasks) > 0 {
		sort.Slice(tasks, func(i, j int) bool { return len(tasks[i].choices) < len(tasks[j].choices) })
		ct := tasks[0]
		tasks = tasks[1:]
		if len(ct.choices) != 1 {
			panic(fmt.Sprintf("ambiguous/unsatisfiable: %s -> %v", ct.allergen, ct.choices))
		}
		a, i := ct.allergen, ct.choices[0]
		a2i[a], i2a[i] = i, a
		for nt := range tasks {
			tasks[nt].choices = prune(tasks[nt].choices, i)
		}
	}

	part1 := 0
	for _, l := range labels {
		for _, i := range l.ingredients {
			if i2a[i] == "" {
				part1++
			}
		}
	}

	order := []string(nil)
	for a := range a2i {
		order = append(order, a)
	}
	sort.Strings(order)

	part2 := []string(nil)
	for _, a := range order {
		part2 = append(part2, a2i[a])
	}
	return part1, strings.Join(part2, ",")
}

func intersect(a, b []string) (out []string) {
	for len(a) > 0 && len(b) > 0 {
		switch {
		case a[0] == b[0]:
			out = append(out, a[0])
			a, b = a[1:], b[1:]
		case a[0] < b[0]:
			a = a[1:]
		default:
			b = b[1:]
		}
	}
	return out
}

func prune(xs []string, x string) (out []string) {
	at := sort.Search(len(xs), func(i int) bool { return xs[i] >= x })
	if at >= len(xs) || xs[at] != x {
		return xs
	}
	out = xs[:len(xs)-1]
	copy(out[at:], xs[at+1:])
	return out
}
