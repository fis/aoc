// Package day05 solves AoC 2020 day 5.
package day05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 5, glue.LineSolver(solve))
}

func solve(passes []string) ([]string, error) {
	decoded, err := decodePasses(passes)
	if err != nil {
		return nil, err
	}

	var (
		max, gap int
		found    [1024]bool
	)
	for _, rc := range decoded {
		id := seatID(rc[0], rc[1])
		found[id] = true
		if id > max {
			max = id
		}
	}
	for id := range found {
		if id > 0 && id < 1023 && !found[id] && found[id-1] && found[id+1] {
			gap = id
			break
		}
	}

	return glue.Ints(max, gap), nil
}

var passPattern = regexp.MustCompile(`^[FB]{7}[LR]{3}$`)

func decodePasses(passes []string) (out [][2]int, err error) {
	for _, pass := range passes {
		row, col, err := decodePass(pass)
		if err != nil {
			return nil, err
		}
		out = append(out, [2]int{row, col})
	}
	return out, nil
}

func decodePass(pass string) (row, col int, err error) {
	if !passPattern.MatchString(pass) {
		return 0, 0, fmt.Errorf("invalid pass: %q", pass)
	}
	bits := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(pass)
	r, _ := strconv.ParseInt(bits[0:7], 2, 64)
	c, _ := strconv.ParseInt(bits[7:10], 2, 64)
	return int(r), int(c), nil
}

func seatID(row, col int) int {
	return row<<3 | col
}
