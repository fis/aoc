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

// Package day18 solves AoC 2021 day 18.
package day18

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 18, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) < 2 {
		return nil, fmt.Errorf("expected at least 2 lines, got %d", len(lines))
	}
	list, err := parseNumbers(lines)
	if err != nil {
		return nil, err
	}
	sum, t := &snailNumber{}, &snailNumber{}
	sum.add(&list[0], &list[1])
	for i := 2; i < len(list); i++ {
		t.add(sum, &list[i])
		sum, t = t, sum
	}
	p1, p2 := sum.magnitude(0), 0
	for i := range list {
		for j := range list {
			if i == j {
				continue
			}
			t.add(&list[i], &list[j])
			if m := t.magnitude(0); m > p2 {
				p2 = m
			}
		}
	}
	return glue.Ints(p1, p2), nil
}

func parseNumbers(lines []string) ([]snailNumber, error) {
	list := make([]snailNumber, len(lines))
	for i, line := range lines {
		if err := list[i].parse(line); err != nil {
			return nil, err
		}
	}
	return list, nil
}

type snailNumber struct {
	data [1 + 2 + 4 + 8 + 16 + 32]snailItem
}

type snailItem struct {
	value, split, depth snailValue
}

type snailValue = int32

const pairValue snailValue = 0xff

func (out *snailNumber) add(a, b *snailNumber) {
	out.data[0].value = pairValue
	out.data[0].split = 0 // never need a split at first
	out.data[0].depth = maxv(a.data[0].depth, b.data[0].depth) + 1
	out.insert(1, a, 0)
	out.insert(2, b, 0)
	out.reduce()
}

func (dst *snailNumber) insert(di int, src *snailNumber, si int) {
	dst.data[di] = src.data[si]
	if src.data[si].value == pairValue {
		dst.insert(2*di+1, src, 2*si+1)
		dst.insert(2*di+2, src, 2*si+2)
	}
}

func (sn *snailNumber) reduce() {
	for {
		switch {
		case sn.data[0].depth == 5:
			sn.explode(0)
		case sn.data[0].split != 0:
			sn.split(0)
		default:
			return
		}
	}
}

func (sn *snailNumber) explode(i int) (addLeft, addRight snailValue) {
	switch {
	case i >= 1+2+4+8:
		l, r := sn.data[2*i+1].value, sn.data[2*i+2].value
		sn.data[i] = snailItem{value: 0, split: 0, depth: 0}
		return l, r
	case sn.data[2*i+1].depth >= sn.data[2*i+2].depth:
		l, r := sn.explode(2*i + 1)
		sn.addLeft(2*i+2, r)
		sn.data[i].split = sn.data[2*i+1].split | sn.data[2*i+2].split
		sn.data[i].depth = maxv(sn.data[2*i+1].depth, sn.data[2*i+2].depth) + 1
		return l, 0
	default:
		l, r := sn.explode(2*i + 2)
		sn.addRight(2*i+1, l)
		sn.data[i].split = sn.data[2*i+1].split | sn.data[2*i+2].split
		sn.data[i].depth = maxv(sn.data[2*i+1].depth, sn.data[2*i+2].depth) + 1
		return 0, r
	}
}

func (sn *snailNumber) addLeft(i int, v snailValue) {
	if v == 0 {
		return
	}
	if sn.data[i].value != pairValue {
		nv := sn.data[i].value + v
		sn.data[i].value = nv
		if nv >= 10 {
			sn.data[i].split = 1
		} else {
			sn.data[i].split = 0
		}
		return
	}
	sn.addLeft(2*i+1, v)
	sn.data[i].split = sn.data[2*i+1].split | sn.data[2*i+2].split
}

func (sn *snailNumber) addRight(i int, v snailValue) {
	if v == 0 {
		return
	}
	if sn.data[i].value != pairValue {
		nv := sn.data[i].value + v
		sn.data[i].value = nv
		if nv >= 10 {
			sn.data[i].split = 1
		} else {
			sn.data[i].split = 0
		}
		return
	}
	sn.addRight(2*i+2, v)
	sn.data[i].split = sn.data[2*i+1].split | sn.data[2*i+2].split
}

func (sn *snailNumber) split(i int) {
	li, ri := 2*i+1, 2*i+2
	switch {
	case sn.data[i].value != pairValue:
		l, r := sn.data[i].value/2, (sn.data[i].value+1)/2
		sn.data[li] = snailItem{value: l, split: 0, depth: 0}
		sn.data[ri] = snailItem{value: r, split: 0, depth: 0}
		sn.data[i] = snailItem{value: pairValue, split: 0, depth: 1}
		if l >= 10 {
			sn.data[li].split = 1
			sn.data[i].split = 1
		}
		if r >= 10 {
			sn.data[ri].split = 1
			sn.data[i].split = 1
		}
	case sn.data[li].split != 0:
		sn.split(li)
		sn.data[i].split = sn.data[li].split | sn.data[ri].split
		sn.data[i].depth = maxv(sn.data[li].depth, sn.data[ri].depth) + 1
	default:
		sn.split(ri)
		sn.data[i].split = sn.data[li].split | sn.data[ri].split
		sn.data[i].depth = maxv(sn.data[li].depth, sn.data[ri].depth) + 1
	}
}

func (sn *snailNumber) magnitude(i int) int {
	if v := sn.data[i].value; v != pairValue {
		return int(v)
	}
	return 3*sn.magnitude(2*i+1) + 2*sn.magnitude(2*i+2)
}

func (sn *snailNumber) format() string {
	var sb strings.Builder
	sn.formatTo(&sb, 0)
	return sb.String()
}

func (sn *snailNumber) formatTo(sb *strings.Builder, i int) {
	if v := sn.data[i].value; v != pairValue {
		sb.WriteByte('0' + byte(v))
		return
	}
	sb.WriteByte('[')
	sn.formatTo(sb, 2*i+1)
	sb.WriteByte(',')
	sn.formatTo(sb, 2*i+2)
	sb.WriteByte(']')
}

func (sn *snailNumber) parse(s string) error {
	tail, err := sn.read(0, s)
	if err != nil {
		return err
	}
	if tail != "" {
		return fmt.Errorf("unexpected trailing content: %q", tail)
	}
	return nil
}

func (sn *snailNumber) read(i int, s string) (tail string, err error) {
	switch {
	case s == "":
		return "", fmt.Errorf("expected [ or digit, got %q", s)
	case s[0] >= '0' && s[0] <= '9':
		v := snailValue(s[0] - '0')
		sn.data[i] = snailItem{value: v, split: 0, depth: 0}
		return s[1:], nil
	case s[0] == '[':
		tail, err = sn.read(2*i+1, s[1:])
		if err != nil {
			return "", err
		}
		if tail == "" || tail[0] != ',' {
			return "", fmt.Errorf("expected comma, got %q", tail)
		}
		tail, err = sn.read(2*i+2, tail[1:])
		if err != nil {
			return "", err
		}
		if tail == "" || tail[0] != ']' {
			return "", fmt.Errorf("expected ], got %q", tail)
		}
		sn.data[i].value = pairValue
		sn.data[i].split = 0 // can't read overly large numbers
		sn.data[i].depth = maxv(sn.data[2*i+1].depth, sn.data[2*i+2].depth) + 1
		return tail[1:], nil
	default:
		return "", fmt.Errorf("expected [ or digit, got %q", s)
	}
}

func maxv(a, b snailValue) snailValue {
	if a > b {
		return a
	}
	return b
}
