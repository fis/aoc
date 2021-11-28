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

// Package day10 solves AoC 2017 day 10.
package day10

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 10, glue.GenericSolver(solve))
}

func solve(input io.Reader) ([]string, error) {
	lines, err := util.ScanAll(input, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	parts := strings.Split(lines[0], ",")
	lengths := make([]byte, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("not a number: %q: %w", part, err)
		}
		lengths[i] = byte(n)
	}
	p1 := part1(256, lengths)
	p2 := part2(256, 64, lines[0])
	return []string{strconv.Itoa(p1), p2}, nil
}

func part1(N int, lengths []byte) int {
	list := generateList(N)
	round(0, 0, list, lengths)
	return int(list[0]) * int(list[1])
}

func part2(N, rounds int, input string) string {
	list := generateList(N)
	lengths := append([]byte(input), []byte{17, 31, 73, 47, 23}...)
	pos, skip := 0, 0
	for i := 0; i < rounds; i++ {
		pos, skip = round(pos, skip, list, lengths)
	}
	hash := compact(list, 16)
	return fmt.Sprintf("%x", hash)
}

func generateList(N int) []byte {
	list := make([]byte, N)
	for i := 0; i < N; i++ {
		list[i] = byte(i)
	}
	return list
}

func round(pos, skip int, list []byte, lengths []byte) (newPos, newSkip int) {
	N := len(list)
	for _, length := range lengths {
		reverse(list, pos, int(length))
		pos = (pos + int(length) + skip) % N
		skip++
	}
	return pos, skip
}

func compact(list []byte, factor int) (hash []byte) {
	N := len(list) / factor
	hash = make([]byte, N)
	for i := 0; i < N; i++ {
		v := byte(0)
		for j := 0; j < factor; j++ {
			v ^= list[i*factor+j]
		}
		hash[i] = v
	}
	return hash
}

func reverse(list []byte, pos, length int) {
	N := len(list)
	for i := 0; i < length/2; i++ {
		pa, pb := (pos+i)%N, (pos+length-1-i)%N
		a, b := list[pa], list[pb]
		list[pa], list[pb] = b, a
	}
}
