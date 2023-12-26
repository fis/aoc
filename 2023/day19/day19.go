// Copyright 2023 Google LLC
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

// Package day19 solves AoC 2023 day 19.
package day19

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 19, glue.LineSolver(solve))
	glue.RegisterPlotter(2023, 19, "", glue.LinePlotter(plot), map[string]string{"ex": ex}) // TODO: ex
}

func solve(lines []string) ([]string, error) {
	workflows, parts, err := parseInput(lines)
	if err != nil {
		return nil, err
	}
	p1 := workflows.sumAccepted(parts)
	p2 := workflows.countAccepted()
	return glue.Ints(p1, p2), nil
}

const numCategories = 4

var catLabels = map[byte]int{'x': 0, 'm': 1, 'a': 2, 's': 3}

type part [numCategories]uint32

type interval struct{ lo, hi uint32 }

const (
	actAccept = -1
	actReject = -2
)

type workflowSet []workflow

func (wfs workflowSet) sumAccepted(parts []part) (sum int) {
	for _, p := range parts {
		if wfs.accepts(p) {
			sum += int(fn.Sum(p[:]))
		}
	}
	return sum
}

func (wfs workflowSet) accepts(p part) bool {
	wf := 0
	for wf >= 0 {
		wf = wfs[wf].evaluate(p)
	}
	return wf == actAccept
}

func (wfs workflowSet) countAccepted() int {
	iv := interval{1, 4000}
	return wfs.countAcceptedIn(0, [...]interval{iv, iv, iv, iv})
}

func (wfs workflowSet) countAcceptedIn(wf int, space [numCategories]interval) int {
	switch wf {
	case actAccept:
		p := 1
		for _, s := range space {
			p *= max(0, int(s.hi-s.lo+1))
		}
		return p
	case actReject:
		return 0
	}
	c := 0
	for _, r := range wfs[wf].rules {
		subSpace := space
		old := space[r.cat]
		switch r.op {
		case '<':
			subSpace[r.cat] = interval{old.lo, r.val - 1}
			space[r.cat] = interval{r.val, old.hi}
		case '>':
			subSpace[r.cat] = interval{r.val + 1, old.hi}
			space[r.cat] = interval{old.lo, r.val}
		}
		c += wfs.countAcceptedIn(r.dst, subSpace)
	}
	c += wfs.countAcceptedIn(wfs[wf].def, space)
	return c
}

type workflow struct {
	label string
	rules []rule
	def   int
}

func (wf workflow) evaluate(p part) int {
	for _, rule := range wf.rules {
		if rule.matches(p) {
			return rule.dst
		}
	}
	return wf.def
}

type rule struct {
	op  byte
	cat int
	val uint32
	dst int
}

func (r rule) matches(p part) bool {
	switch r.op {
	case '<':
		return p[r.cat] < r.val
	case '>':
		return p[r.cat] > r.val
	}
	return false
}

func parseInput(lines []string) (workflowSet, []part, error) {
	gap := slices.Index(lines, "")
	if gap < 0 {
		return nil, nil, fmt.Errorf("missing blank line between workflows and parts")
	}

	workflows, err := parseWorkflows(lines[:gap])
	if err != nil {
		return nil, nil, err
	}

	partLines := lines[gap+1:]
	parts := make([]part, len(partLines))
	for i, line := range partLines {
		parts[i], err = parsePart(line)
		if err != nil {
			return nil, nil, err
		}
	}

	return workflows, parts, nil
}

func parseWorkflows(lines []string) (wfs workflowSet, err error) {
	const entryWorkflow = "in"
	wfId := make(workflowMap)
	wfId.get(entryWorkflow) // ensure it gets ID 0

	workflows := make(workflowSet, len(lines))
	for _, line := range lines {
		wfName, spec, ok := strings.Cut(line, "{")
		if !ok {
			return nil, fmt.Errorf("missing '{' between workflow name and content: %q", line)
		}
		spec, ok = strings.CutSuffix(spec, "}")
		if !ok {
			return nil, fmt.Errorf("missing '}' at the end of workflow: %q", line)
		}
		numRules := strings.Count(spec, ",")
		rules := make([]rule, numRules)
		for i := range rules {
			ruleText, tail, _ := strings.Cut(spec, ",")
			rules[i], err = parseRule(ruleText, wfId)
			if err != nil {
				return nil, err
			}
			spec = tail
		}
		workflows[wfId.get(wfName)] = workflow{
			label: wfName,
			rules: rules,
			def:   wfId.get(spec),
		}
	}
	return workflows, nil
}

func parseRule(ruleText string, wfId workflowMap) (rule, error) {
	cond, dst, ok := strings.Cut(ruleText, ":")
	if !ok {
		return rule{}, fmt.Errorf("missing ':' between condition and destination: %q", ruleText)
	}
	if len(cond) < 3 {
		return rule{}, fmt.Errorf("condition too short: %q", cond)
	}
	cat, ok := catLabels[cond[0]]
	if !ok {
		return rule{}, fmt.Errorf("unknown category: %q", cond[0])
	}
	op := cond[1]
	if op != '<' && op != '>' {
		return rule{}, fmt.Errorf("unknown operator: %q", op)
	}
	val, err := strconv.ParseUint(cond[2:], 10, 32)
	if err != nil {
		return rule{}, fmt.Errorf("bad condition value: %w", err)
	}
	return rule{op: op, cat: cat, val: uint32(val), dst: wfId.get(dst)}, nil
}

func parsePart(line string) (p part, err error) {
	line = strings.TrimPrefix(line, "{")
	line = strings.TrimSuffix(line, "}")
	for len(line) > 0 {
		spec, tail, ok := strings.Cut(line, ",")
		if ok {
			line = tail
		} else {
			spec, line = line, ""
		}
		if len(spec) < 3 || spec[1] != '=' {
			return part{}, fmt.Errorf("bad part attribute: %q", spec)
		}
		cat, ok := catLabels[spec[0]]
		if !ok {
			return part{}, fmt.Errorf("bad attribute category: %q", spec[0])
		}
		val, err := strconv.ParseUint(spec[2:], 10, 32)
		if err != nil {
			return part{}, fmt.Errorf("bad attribute value: %w", err)
		}
		p[cat] = uint32(val)
	}
	return p, nil
}

type workflowMap util.LabelMap

func (m workflowMap) get(label string) int {
	switch label {
	case "A":
		return actAccept
	case "R":
		return actReject
	}
	return util.LabelMap(m).Get(label)
}

// plotting

var ex = strings.TrimPrefix(`
px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
`, "\n")

func plot(lines []string, w io.Writer) error {
	workflows, _, err := parseInput(lines)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "digraph G {")
	fmt.Fprintln(w, "  graph [rankdir=\"LR\",dpi=60,size=600];")
	fmt.Fprintln(w, "  start [shape=\"point\"];")
	for i, wf := range workflows {
		var label strings.Builder
		fmt.Fprintf(&label, "<l> %s", wf.label)
		for j, r := range wf.rules {
			fmt.Fprintf(&label, "|<r%d> %c%s%d", j, "xmas"[r.cat], fn.If(r.op == '<', "&lt;", "&gt;"), r.val)
		}
		fmt.Fprintf(&label, "|<r%d> else", len(wf.rules))
		fmt.Fprintf(w, "  w%d [label=\"%s\",shape=record];\n", i, label.String())
		for j, r := range append(wf.rules, rule{dst: wf.def}) {
			if r.dst == actAccept || r.dst == actReject {
				fmt.Fprintf(w, "  w%dr%da [label=\"%c\",shape=plaintext,width=0,height=0];\n", i, j, fn.If(r.dst == actAccept, 'A', 'R'))
			}
		}
	}
	fmt.Fprintln(w, "  start -> w0:l;")
	for i, wf := range workflows {
		for j, r := range append(wf.rules, rule{dst: wf.def}) {
			if r.dst == actAccept || r.dst == actReject {
				fmt.Fprintf(w, "  w%d:r%d -> w%dr%da;\n", i, j, i, j)
			} else {
				fmt.Fprintf(w, "  w%d:r%d -> w%d:l;\n", i, j, r.dst)
			}
		}
	}
	fmt.Fprintln(w, "}")

	return nil
}
