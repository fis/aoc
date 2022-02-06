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

// Package day05 solves AoC 2016 day 5.
package day05

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 5, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	p1, p2 := search(lines[0], 8)
	return []string{p1, p2}, nil
}

var hexDigits = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func search(prefix string, count int) (string, string) {
	buf := append([]byte(prefix), make([]byte, 16)...)
	p1, p2 := make([]byte, 0, count), make([]byte, count)
	found := 0
	for i := int64(0); ; i++ {
		in := strconv.AppendInt(buf[:len(prefix)], i, 10)
		out := md5.Sum(in)
		if out[0] == 0 && out[1] == 0 && (out[2]&0xf0) == 0 {
			if len(p1) < count {
				p1 = append(p1, hexDigits[out[2]&0x0f])
			}
			if pos := out[2] & 0x0f; int(pos) < count && p2[pos] == 0 {
				p2[pos] = hexDigits[out[3]>>4]
				found++
				if found >= count {
					break
				}
			}
		}
	}
	return string(p1), string(p2)
}
