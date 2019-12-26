// Package day14 solves AoC 2019 day 14.
package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc2019-go/util"
)

func Solve(path string) ([]string, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	reactions := parseReactions(lines)

	part1 := ore(1, reactions)
	part2 := maxFuel(1000000000000, reactions)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}, nil
}

type pile struct {
	name string
	q int
}

type reaction struct {
	out pile
	in []pile
}

func ore(wantFuel int, reactions map[string]reaction) int {
	order := reactionOrder(reactions)
	want := map[string]int{"FUEL": wantFuel}
	return oreFor(want, order, reactions)
}

func maxFuel(ore int, reactions map[string]reaction) int {
	order := reactionOrder(reactions)
	start, end := 1, ore+1
	for end - start >= 2 {
		mid := start + (end - start) / 2
		got := oreFor(map[string]int{"FUEL": mid}, order, reactions)
		if got > ore {
			end = mid
		} else {
			start = mid
		}
	}
	return start
}

func reactionOrder(reactions map[string]reaction) []string {
	var g util.Graph
	for out, r := range reactions {
		for _, in := range r.in {
			g.AddEdge(out, in.name)
		}
	}
	return g.TopoSort()
}

func oreFor(want map[string]int, order []string, reactions map[string]reaction) int {
	for _, ch := range order {
		n, ok := want[ch]
		if !ok {
			continue // not needed
		}
		if ch == "ORE" {
			return n
		}
		delete(want, ch)
		r := reactions[ch]
		k := (n + r.out.q - 1) / r.out.q
		for _, in := range r.in {
			want[in.name] += k * in.q
		}
	}
	panic("no ore required")
}

func parseReactions(lines []string) map[string]reaction {
	reactions := make(map[string]reaction)
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid reaction: %s", line))
		}
		r := reaction{out: parsePile(parts[1])}
		for _, spec := range strings.Split(parts[0], ", ") {
			r.in = append(r.in, parsePile(spec))
		}
		reactions[r.out.name] = r
	}
	return reactions
}

func parsePile(spec string) pile {
	parts := strings.Split(spec, " ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid pile: %s", spec))
	}
	q, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Sprintf("invalid pile (not a number): %s", spec))
	}
	return pile{name: parts[1], q: q}
}
