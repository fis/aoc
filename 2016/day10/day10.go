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

// Package day10 solves AoC 2016 day 10.
package day10

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2016, 10, glue.RegexpSolver{Solver: solve, Regexp: inputRegexp})
	glue.RegisterPlotter(2016, 10, plotter{}, map[string]string{"ex": ex})
}

const inputRegexp = `value (\d+) goes to bot (\d+)|bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`

var ex = strings.TrimPrefix(`
value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
`, "\n")

func solve(input [][]string) ([]string, error) {
	g := buildGraph(input)
	//g.writeDOT(os.Stdout)
	p1 := g.findBot(61, 17)
	p2 := g.outputs[0] * g.outputs[1] * g.outputs[2]
	return glue.Ints(p1, p2), nil
}

type graph struct {
	inputs  []link
	bots    map[int]*bot
	outputs map[int]int
}

type bot struct {
	outL, outH link
	in1, in2   int
}

type link struct {
	bot bool
	id  int
	val int
}

func buildGraph(input [][]string) *graph {
	g := &graph{bots: make(map[int]*bot), outputs: make(map[int]int)}
	type sig struct{ to, val int }
	var q []sig
	for _, inst := range input {
		if inst[0] != "" {
			val, _ := strconv.Atoi(inst[0])
			bot, _ := strconv.Atoi(inst[1])
			g.inputs = append(g.inputs, link{bot: true, id: bot, val: val})
			q = append(q, sig{to: bot, val: val})
		} else {
			from, _ := strconv.Atoi(inst[2])
			lo, _ := strconv.Atoi(inst[4])
			hi, _ := strconv.Atoi(inst[6])
			g.bots[from] = &bot{
				outL: link{bot: inst[3] == "bot", id: lo},
				outH: link{bot: inst[5] == "bot", id: hi},
			}
		}
	}
	for len(q) > 0 {
		to, val := g.bots[q[0].to], q[0].val
		q = q[1:]
		switch {
		case to.in1 == 0:
			to.in1 = val
		case to.in2 == 0:
			to.in2 = val
			lo, hi := inOrder(to.in1, to.in2)
			to.outL.val = lo
			if to.outL.bot {
				q = append(q, sig{to: to.outL.id, val: lo})
			} else {
				g.outputs[to.outL.id] = lo
			}
			to.outH.val = hi
			if to.outH.bot {
				q = append(q, sig{to: to.outH.id, val: hi})
			} else {
				g.outputs[to.outH.id] = hi
			}
		}
	}
	return g
}

func (g *graph) findBot(a, b int) int {
	for id, bot := range g.bots {
		if bot.in1 == a && bot.in2 == b || bot.in1 == b && bot.in2 == a {
			return id
		}
	}
	return -1
}

type plotter struct{}

func (plotter) Plot(r io.Reader, w io.Writer) error {
	input, err := util.ScanAllRegexp(r, inputRegexp)
	if err != nil {
		return err
	}
	g := buildGraph(input)
	g.writeDOT(w)
	return nil
}

func (g *graph) writeDOT(w io.Writer) {
	fmt.Fprintln(w, "digraph bots {")
	for _, in := range g.inputs {
		fmt.Fprintf(w, "  i%d [label=\"%d\",style=\"filled\",fillcolor=\"#0f9d58\",fontcolor=\"white\"];\n", in.val, in.val)
		fmt.Fprintf(w, "  i%d -> b%d;\n", in.val, in.id)
	}
	linkChar := map[bool]rune{true: 'b', false: 'o'}
	for id, bot := range g.bots {
		fmt.Fprintf(w, "  b%d [label=\"bot %d\"];\n", id, id)
		fmt.Fprintf(w, "  b%d -> %c%d [label=\"%d\"];\n", id, linkChar[bot.outL.bot], bot.outL.id, bot.outL.val)
		fmt.Fprintf(w, "  b%d -> %c%d [label=\"%d\"];\n", id, linkChar[bot.outH.bot], bot.outH.id, bot.outH.val)
	}
	for id := range g.outputs {
		fmt.Fprintf(w, "  o%d [label=\"out %d\",style=\"filled\",fillcolor=\"#db4437\",fontcolor=\"white\"];\n", id, id)
	}
	fmt.Fprintln(w, "}")
}

func inOrder(a, b int) (lo, hi int) {
	if a < b {
		return a, b
	}
	return b, a
}
