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

// Package day04 solves AoC 2016 day 4.
package day04

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 4, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := 0, -1
	for _, line := range lines {
		r, err := parseRoom(line)
		if err != nil {
			return nil, err
		}
		if r.validate() {
			p1 += r.sector
			if p2 < 0 && r.decode() == "northpole object storage" {
				p2 = r.sector
			}
		}
	}
	return glue.Ints(p1, p2), nil
}

type room struct {
	name     []string
	sector   int
	checksum [5]byte
}

func (r *room) validate() bool {
	var cs checksummer
	for _, part := range r.name {
		for _, b := range []byte(part) {
			cs.add(b)
		}
	}
	return r.checksum == cs.compute()
}

func (r *room) decode() string {
	var decoded strings.Builder
	for i, part := range r.name {
		if i > 0 {
			decoded.WriteByte(' ')
		}
		for _, b := range []byte(part) {
			decoded.WriteByte(byte('a' + (int(b)-'a'+r.sector)%26))
		}
	}
	return decoded.String()
}

func parseRoom(line string) (r room, err error) {
	parts := strings.Split(line, "-")
	if len(parts) < 2 {
		return room{}, fmt.Errorf("invalid room: %q: not enough components", line)
	}
	r.name = parts[:len(parts)-1]
	tail := parts[len(parts)-1]
	if len(tail) <= 7 || tail[len(tail)-7] != '[' || tail[len(tail)-1] != ']' {
		return room{}, fmt.Errorf("invalid room tail: %q: not ...[abcde]", tail)
	}
	r.sector, err = strconv.Atoi(tail[:len(tail)-7])
	if err != nil {
		return room{}, fmt.Errorf("invalid room sector ID: %w", err)
	}
	r.checksum = *(*[5]byte)([]byte(tail[len(tail)-6 : len(tail)-1]))
	return r, nil
}

type checksummer [26]int

func (cs *checksummer) add(b byte) {
	if b >= 'a' && b <= 'z' {
		cs[b-'a']++
	}
}

func (cs *checksummer) compute() (out [5]byte) {
	order := make([]int, 26)
	for i := 1; i < 26; i++ {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		a, b := order[i], order[j]
		if cs[a] > cs[b] {
			return true
		} else if cs[a] < cs[b] {
			return false
		}
		return a < b
	})
	for i := 0; i < 5; i++ {
		out[i] = byte('a' + order[i])
	}
	return out
}
