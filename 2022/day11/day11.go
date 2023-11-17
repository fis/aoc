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

// Package day11 solves AoC 2022 day 11.
package day11

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 11, glue.ChunkSolver(glue.WithParser(parseMonkey, solve)))
}

func solve(monkeys []monkey) ([]string, error) {
	p1 := monkeyBusiness(monkeys, 20)
	p2 := worryingMonkeyBusiness(monkeys, 10000)
	return glue.Ints(p1, p2), nil
}

func monkeyBusiness(monkeys []monkey, rounds int) int {
	items := fn.Map(monkeys, func(m monkey) []int { return slices.Clone(m.startItems) })
	act := make([]int, len(monkeys))
	for round := 0; round < rounds; round++ {
		for i, m := range monkeys {
			act[i] += len(items[i])
			for _, worry := range items[i] {
				worry = m.op.update(worry) / 3
				next := m.next[fn.If(worry%m.test == 0, 0, 1)]
				items[next] = append(items[next], worry)
			}
			items[i] = items[i][:0]
		}
	}
	slices.Sort(act)
	return act[len(act)-2] * act[len(act)-1]
}

func worryingMonkeyBusiness(monkeys []monkey, rounds int) int {
	items := fn.Map(monkeys, func(m monkey) []int { return slices.Clone(m.startItems) })
	act := make([]int, len(monkeys))
	modulo := fn.Prod(fn.Map(monkeys, func(m monkey) int { return m.test }))
	for round := 0; round < rounds; round++ {
		for i, m := range monkeys {
			act[i] += len(items[i])
			for _, worry := range items[i] {
				worry = m.op.update(worry) % modulo
				next := m.next[fn.If(worry%m.test == 0, 0, 1)]
				items[next] = append(items[next], worry)
			}
			items[i] = items[i][:0]
		}
	}
	slices.Sort(act)
	return act[len(act)-2] * act[len(act)-1]
}

type monkey struct {
	startItems []int
	op         monkeyOp
	test       int
	next       [2]int
}

type monkeyOp interface {
	update(worry int) int
}

type opAdd int

func (n opAdd) update(worry int) int {
	return worry + int(n)
}

type opMul int

func (n opMul) update(worry int) int {
	return worry * int(n)
}

type opSquare struct{}

func (opSquare) update(worry int) int {
	return worry * worry
}

func parseMonkey(chunk string) (m monkey, err error) {
	const (
		prefixID        = "Monkey "
		prefixItems     = "  Starting items: "
		prefixOp        = "  Operation: new = old "
		prefixTest      = "  Test: divisible by "
		prefixNextTrue  = "    If true: throw to monkey "
		prefixNextFalse = "    If false: throw to monkey "
	)
	lines := util.Lines(chunk)
	if len(lines) != 6 {
		return monkey{}, fmt.Errorf("bad monkey: unexpected line count: %#v", lines)
	}
	if !strings.HasPrefix(lines[0], prefixID) {
		return monkey{}, fmt.Errorf("bad monkey: expected ID, got: %q", lines[0])
	}
	if items, ok := util.CheckPrefix(lines[1], prefixItems); !ok {
		return monkey{}, fmt.Errorf("bad monkey: expected item list, got: %q", lines[1])
	} else if m.startItems, err = fn.MapE(strings.Split(items, ", "), strconv.Atoi); err != nil {
		return monkey{}, fmt.Errorf("bad monkey: expected item: %w", err)
	}
	if spec, ok := util.CheckPrefix(lines[2], prefixOp); !ok {
		return monkey{}, fmt.Errorf("bad monkey: expected op, got: %q", lines[2])
	} else if m.op, err = parseMonkeyOp(spec); err != nil {
		return monkey{}, fmt.Errorf("bad monkey: expected op: %w", err)
	}
	if test, ok := util.CheckPrefix(lines[3], prefixTest); !ok {
		return monkey{}, fmt.Errorf("bad monkey: expected test, got: %q", lines[3])
	} else if m.test, err = strconv.Atoi(test); err != nil {
		return monkey{}, fmt.Errorf("bad monkey: expected test: %w", err)
	}
	for i := 0; i < 2; i++ {
		if next, ok := util.CheckPrefix(lines[4+i], fn.If(i == 0, prefixNextTrue, prefixNextFalse)); !ok {
			return monkey{}, fmt.Errorf("bad monkey: expected next, got: %q", lines[4+i])
		} else if m.next[i], err = strconv.Atoi(next); err != nil {
			return monkey{}, fmt.Errorf("bad monkey: expected next: %w", err)
		}
	}
	return m, nil
}

func parseMonkeyOp(spec string) (monkeyOp, error) {
	if spec == "* old" {
		return opSquare{}, nil
	}
	if len(spec) < 3 || (spec[0] != '+' && spec[0] != '*') || spec[1] != ' ' {
		return nil, fmt.Errorf("misformatted spec: %q", spec)
	}
	n, err := strconv.Atoi(spec[2:])
	if err != nil {
		return nil, fmt.Errorf("bad operand: %w", err)
	}
	return fn.If[monkeyOp](spec[0] == '+', opAdd(n), opMul(n)), nil
}
