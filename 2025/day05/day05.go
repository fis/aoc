// Package day05 solves AoC 2025 day 5.
package day05

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2025, 5, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	freshIntervals, err := fn.MapE(util.Lines(chunks[0]), parseInterval)
	if err != nil {
		return nil, fmt.Errorf("failed to parse fresh intervals: %w", err)
	}
	ingredients, err := fn.MapE(util.Lines(chunks[1]), strconv.Atoi)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ingredient list: %w", err)
	}
	p1, p2 := countFresh(freshIntervals, ingredients)
	return glue.Ints(p1, p2), nil
}

func countFresh(intervals []interval, ingredients []int) (available, allFresh int) {
	if len(intervals) == 0 {
		return 0, 0
	}

	slices.SortFunc(intervals, func(a, b interval) int { return cmp.Compare(a.start, b.start) })
	merged := make(ivList, 1, len(intervals))
	merged[0] = intervals[0]
	for _, iv := range intervals[1:] {
		last := &merged[len(merged)-1]
		switch {
		case iv.start > last.end+1:
			merged = append(merged, iv) // gap between
		case iv.end > last.end:
			last.end = iv.end // overlapping, but extends further
		default:
			// completely inside, ignore
		}
	}

	for _, ing := range ingredients {
		if merged.contains(ing) {
			available++
		}
	}

	for _, iv := range merged {
		allFresh += iv.len()
	}

	return available, allFresh
}

type interval struct{ start, end int }

func (iv interval) len() int            { return iv.end - iv.start + 1 }
func (iv interval) contains(i int) bool { return i >= iv.start && i <= iv.end }

type ivList []interval

func (ivs ivList) contains(i int) bool {
	// TODO
	for len(ivs) >= 2 {
		mi := len(ivs) / 2
		switch m := ivs[mi]; {
		case i < m.start:
			ivs = ivs[:mi]
		case i > m.end:
			ivs = ivs[mi+1:]
		default:
			return true
		}
	}
	return len(ivs) == 1 && ivs[0].contains(i)
}

func parseInterval(line string) (interval, error) {
	startS, endS, ok := strings.Cut(line, "-")
	if !ok {
		return interval{}, fmt.Errorf("invalid interval: %s", line)
	}
	start, err1 := strconv.Atoi(startS)
	end, err2 := strconv.Atoi(endS)
	if err1 != nil || err2 != nil {
		return interval{}, fmt.Errorf("invalid endpoints: %s/%s: %w", startS, endS, errors.Join(err1, err2))
	}
	return interval{start: start, end: end}, nil
}
