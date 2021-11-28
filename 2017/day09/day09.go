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

// Package day09 solves AoC 2017 day 9.
package day09

import (
	"fmt"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 9, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	groups, err := parseStream([]byte(strings.Join(lines, "")))
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}
	p1, p2 := score(groups, 1)
	return []int{p1, p2}, nil
}

type group struct {
	data    []group
	garbage int
}

func score(groups []group, depth int) (sumScore, sumGarbage int) {
	for _, g := range groups {
		sumScore += depth
		sumGarbage += g.garbage
		gScore, gGarbage := score(g.data, depth+1)
		sumScore += gScore
		sumGarbage += gGarbage
	}
	return sumScore, sumGarbage
}

func parseStream(input []byte) ([]group, error) {
	tail, groups, _, err := parseGroups(input)
	if err != nil {
		return nil, err
	}
	if len(tail) > 0 {
		return nil, fmt.Errorf("leftover data")
	}
	return groups, nil
}

func parseGroups(input []byte) (tail []byte, groups []group, garbage int, err error) {
loop:
	for len(input) > 0 {
		switch input[0] {
		case '{':
			tail, g, err := parseGroup(input)
			if err != nil {
				return nil, nil, 0, err
			}
			groups = append(groups, g)
			input = tail
		case '<':
			tail, g, err := parseGarbage(input)
			if err != nil {
				return nil, nil, 0, err
			}
			garbage += g
			input = tail
		case '}':
			break loop
		default:
			return nil, nil, 0, fmt.Errorf("expected { or <, got %c", input[0])
		}
		if len(input) == 0 || input[0] == '}' {
			break
		}
		if input[0] != ',' {
			return nil, nil, 0, fmt.Errorf("expected , or } or EOF, got %c", input[0])
		}
		input = input[1:]
	}
	return input, groups, garbage, nil
}

func parseGroup(input []byte) ([]byte, group, error) {
	if len(input) == 0 || input[0] != '{' {
		return nil, group{}, fmt.Errorf("expected {")
	}
	tail, data, garbage, err := parseGroups(input[1:])
	if err != nil {
		return nil, group{}, err
	}
	if len(tail) == 0 || tail[0] != '}' {
		return nil, group{}, fmt.Errorf("expected }")
	}
	return tail[1:], group{data: data, garbage: garbage}, nil
}

func parseGarbage(input []byte) (tail []byte, size int, err error) {
	if len(input) == 0 || input[0] != '<' {
		return nil, 0, fmt.Errorf("expected <")
	}
	input = input[1:]
	for {
		if len(input) == 0 {
			return nil, 0, fmt.Errorf("unterminated garbage")
		}
		if input[0] == '>' {
			return input[1:], size, nil
		}
		if input[0] == '!' {
			if len(input) == 1 {
				return nil, 0, fmt.Errorf("! at EOF")
			}
			input = input[2:]
		} else {
			size++
			input = input[1:]
		}
	}
}
