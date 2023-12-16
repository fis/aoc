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

// Package day15 solves AoC 2023 day 15.
package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 15, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line, got %d", len(lines))
	}
	steps := strings.Split(lines[0], ",")
	p1 := fn.SumF(steps, hash)
	p2 := initialize(fn.Map(steps, parseOperation))
	return glue.Ints(p1, p2), nil
}

func hash(s string) int {
	h := byte(0)
	for _, b := range []byte(s) {
		h += b
		h *= 17
	}
	return int(h)
}

func initialize(ops []lensLabel) (focusingPower int) {
	boxes := make([][]lensLabel, 256)
loop:
	for _, op := range ops {
		h := hash(op.label)
		box := boxes[h]
		if op.value == -1 {
			for i := range box {
				if box[i].label == op.label {
					copy(box[i:], box[i+1:])
					boxes[h] = box[:len(box)-1]
					break
				}
			}
		} else {
			for i := range box {
				if box[i].label == op.label {
					box[i].value = op.value
					continue loop
				}
			}
			boxes[h] = append(boxes[h], op)
		}
	}
	for i, box := range boxes {
		for j, lens := range box {
			focusingPower += (i + 1) * (j + 1) * lens.value
		}
	}
	return focusingPower
}

type lensLabel struct {
	label string
	value int
}

func parseOperation(text string) lensLabel {
	if text[len(text)-1] == '-' {
		return lensLabel{label: text[:len(text)-1], value: -1}
	}
	label, valueText, _ := strings.Cut(text, "=")
	value, _ := strconv.Atoi(valueText)
	return lensLabel{label: label, value: value}
}
