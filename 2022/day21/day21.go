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

// Package day21 solves AoC 2022 day 21.
package day21

import (
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2022, 21, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	root, human := parseJobs(lines)
	p1 := root.num()
	root.forceEqual()
	p2 := human.num()
	return glue.Ints(p1, p2), nil
}

type monkeyJob interface {
	num() int
	hasHuman() bool
	forceTo(v int)
}

type mathOp interface {
	eval(a, b int) int
	solveA(b, num int) int
	solveB(a, num int) int
}

type constJob int

type humanJob int

type mathJob struct {
	op    mathOp
	args  [2]monkeyJob
	human [2]bool
}

func (c constJob) num() int       { return int(c) }
func (c constJob) hasHuman() bool { return false }
func (c constJob) forceTo(v int)  { panic("constants are forever") }

func (h humanJob) num() int       { return int(h) }
func (h humanJob) hasHuman() bool { return true }
func (h *humanJob) forceTo(v int) { *h = humanJob(v) }

func (m *mathJob) num() int {
	return m.op.eval(m.args[0].num(), m.args[1].num())
}

func (m *mathJob) hasHuman() bool {
	m.human[0] = m.args[0].hasHuman()
	m.human[1] = m.args[1].hasHuman()
	return m.human[0] || m.human[1]
}

func (m *mathJob) forceTo(v int) {
	if m.human[0] {
		b := m.args[1].num()
		a := m.op.solveA(b, v)
		m.args[0].forceTo(a)
	} else if m.human[1] {
		a := m.args[0].num()
		b := m.op.solveB(a, v)
		m.args[1].forceTo(b)
	} else {
		panic("no humans to enforce")
	}
}

func (m *mathJob) forceEqual() {
	h0, h1 := m.args[0].hasHuman(), m.args[1].hasHuman()
	if h0 {
		m.args[0].forceTo(m.args[1].num())
	} else if h1 {
		m.args[1].forceTo(m.args[0].num())
	} else {
		panic("no humans to equalize")
	}
}

type mathAdd struct{}
type mathSub struct{}
type mathMul struct{}
type mathDiv struct{}

func (mathAdd) eval(a, b int) int     { return a + b }
func (mathAdd) solveA(b, num int) int { return num - b }
func (mathAdd) solveB(a, num int) int { return num - a }

func (mathSub) eval(a, b int) int     { return a - b }
func (mathSub) solveA(b, num int) int { return num + b } // num = a - b => a = num + b
func (mathSub) solveB(a, num int) int { return a - num } // num = a - b => b = a - num

func (mathMul) eval(a, b int) int     { return a * b }
func (mathMul) solveA(b, num int) int { return num / b }
func (mathMul) solveB(a, num int) int { return num / a }

func (mathDiv) eval(a, b int) int     { return a / b }
func (mathDiv) solveA(b, num int) int { return num * b } // num = a / b => a = num * b
func (mathDiv) solveB(a, num int) int { return a / num } // num = a / b => b = a / num

var mathOps = map[string]mathOp{
	"+": mathAdd{},
	"-": mathSub{},
	"*": mathMul{},
	"/": mathDiv{},
}

func parseJobs(lines []string) (root *mathJob, human *humanJob) {
	constMonkeys := make(map[string]constJob)
	mathMonkeys := make(map[string]*mathJob)
	human = new(humanJob)
	type ref struct {
		from string
		to   [2]string
	}
	var refs []ref
	for _, line := range lines {
		sep := strings.Index(line, ": ")
		name := line[:sep]
		if c, ok, _ := util.NextInt(line[sep+2:]); ok {
			if name == "humn" {
				*human = humanJob(c)
			} else {
				constMonkeys[name] = constJob(c)
			}
			continue
		}
		words := util.Words(line[sep+2:])
		mathMonkeys[name] = &mathJob{op: mathOps[words[1]]}
		refs = append(refs, ref{from: name, to: [2]string{words[0], words[2]}})
	}
	for _, r := range refs {
		var args [2]monkeyJob
		for i, to := range r.to {
			if to == "humn" {
				args[i] = human
			} else if c, ok := constMonkeys[to]; ok {
				args[i] = c
			} else {
				args[i] = mathMonkeys[to]
			}
		}
		mathMonkeys[r.from].args = args
	}
	return mathMonkeys["root"], human
}
