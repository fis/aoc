// Copyright 2021 Google LLC
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

// Package day08 solves AoC 2017 day 8.
package day08

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 8, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2, err := maxRegs(lines)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func maxRegs(lines []string) (maxFinal, maxEver int, err error) {
	insts, err := parseInsts(lines)
	if err != nil {
		return 0, 0, err
	}
	maxFinal, maxEver = math.MinInt, math.MinInt
	regs := make(map[string]int)
	for _, inst := range insts {
		inst.eval(regs)
		if r := regs[inst.reg]; r > maxEver {
			maxEver = r
		}
	}
	for _, val := range regs {
		if val > maxFinal {
			maxFinal = val
		}
	}
	return maxFinal, maxEver, nil
}

type inst struct {
	reg     string
	inc     bool
	arg     int
	condReg string
	cond    condFunc
	condArg int
}

func parseInsts(lines []string) (insts []*inst, err error) {
	insts = make([]*inst, len(lines))
	for i, line := range lines {
		insts[i], err = parseInst(line)
		if err != nil {
			return nil, err
		}
	}
	return insts, nil
}

var instRe = regexp.MustCompile(`^(\w+) (inc|dec) (-?\d+) if (\w+) ([=!<>]=|[<>]) (-?\d+)`)

func parseInst(line string) (*inst, error) {
	parts := instRe.FindStringSubmatch(line)
	if len(parts) != 7 {
		return nil, fmt.Errorf("invalid instruction: %q", line)
	}
	arg, _ := strconv.Atoi(parts[3])
	condArg, _ := strconv.Atoi(parts[6])
	return &inst{
		reg:     parts[1],
		inc:     parts[2] == "inc",
		arg:     arg,
		condReg: parts[4],
		cond:    condFuncs[parts[5]],
		condArg: condArg,
	}, nil
}

func (i *inst) eval(regs map[string]int) {
	if i.cond(regs[i.condReg], i.condArg) {
		if i.inc {
			regs[i.reg] = regs[i.reg] + i.arg
		} else {
			regs[i.reg] = regs[i.reg] - i.arg
		}
	}
}

type condFunc func(a, b int) bool

var condFuncs = map[string]condFunc{
	"==": func(a, b int) bool { return a == b },
	"!=": func(a, b int) bool { return a != b },
	"<":  func(a, b int) bool { return a < b },
	"<=": func(a, b int) bool { return a <= b },
	">=": func(a, b int) bool { return a >= b },
	">":  func(a, b int) bool { return a > b },
}
