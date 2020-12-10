// Copyright 2019 Google LLC
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

// Package day25 solves AoC 2019 day 25.
package day25

import (
	"fmt"
	"strings"

	"github.com/fis/aoc-go/intcode"
	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(25, intcode.SolverS(solve))
}

// TODO: Write a generalized solver for this puzzle.

var motions = []string{
	"east", "take antenna",
	"north", "north", "take asterisk",
	"south", "west", "south", "take hologram",
	"north", "west", "take astronaut ice cream",
	"east", "east", "south", "east", "take ornament",
	"north", "west", "take fixed point",
	"east", "south", "west", "west", "south", "south", "south", "take dark matter",
	"north", "west", "north", "take monolith",
	"north", "north", // at checkpoint
}

func solve(prog []int64) ([]string, error) {
	term := terminal{}
	term.vm.Load(prog)

	var items []string
	term.read()
	for _, cmd := range motions {
		if strings.HasPrefix(cmd, "take ") {
			items = append(items, cmd[5:])
		}
		term.write(cmd)
		term.read()
	}

	dropped := 0
	for _, attempt := range grayCodes(len(items)) {
		for i, item := range items {
			if dropped&(1<<i) != attempt&(1<<i) {
				if dropped&(1<<i) != 0 {
					term.write(fmt.Sprintf("take %s", item))
				} else {
					term.write(fmt.Sprintf("drop %s", item))
				}
				term.read()
			}
		}
		dropped = attempt
		term.write("east")
		out := term.read()
		if !strings.Contains(out, "== Security Checkpoint ==") {
			return extract(out), nil
		}
	}
	panic("access denied")
}

func extract(out string) []string {
	var lines []string
	keep := false
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		if strings.Contains(line, "Santa") {
			keep = true
		}
		if keep {
			lines = append(lines, line)
		}
	}
	return lines
}

func grayCodes(items int) []int {
	n := 1 << items
	out := make([]int, n)
	out[0] = 0
	out[1] = 1
	for k := 2; k < n; k *= 2 {
		for i := 0; i < k; i++ {
			out[2*k-1-i] = out[i] | k
		}
	}
	return out
}

type terminal struct {
	vm    intcode.VM
	token intcode.WalkToken
}

func (t *terminal) read() string {
	out := strings.Builder{}
	for t.vm.Walk(&t.token) && t.token.IsOutput() {
		out.WriteByte(byte(t.token.ReadOutput()))
	}
	return out.String()
}

func (t *terminal) write(line string) {
	for _, r := range line {
		t.token.ProvideInput(int64(r))
		t.vm.Walk(&t.token)
	}
	t.token.ProvideInput(10)
}
