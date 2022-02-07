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

// Package day09 solves AoC 2016 day 9.
package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 9, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	in := strings.Join(lines, "")
	p1, err := decompressLen(in)
	if err != nil {
		return nil, err
	}
	p2, err := recursiveLen(in)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func decompressLen(in string) (out int, err error) {
	for len(in) > 0 {
		if in[0] == '(' {
			size, count, tail, err := parseMarker(in[1:])
			if err != nil {
				return 0, err
			} else if size > len(tail) {
				return 0, fmt.Errorf("repeat (%dx%d), |tail| = %d", size, count, len(tail))
			}
			out += size * count
			in = tail[size:]
			continue
		}
		out++
		in = in[1:]
	}
	return out, nil
}

func recursiveLen(in string) (out int, err error) {
	for len(in) > 0 {
		if in[0] == '(' {
			size, count, tail, err := parseMarker(in[1:])
			if err != nil {
				return 0, err
			} else if size > len(tail) {
				return 0, fmt.Errorf("repeat (%dx%d), |tail| = %d", size, count, len(tail))
			}
			rec, err := recursiveLen(tail[:size])
			if err != nil {
				return 0, err
			}
			out += rec * count
			in = tail[size:]
			continue
		}
		out++
		in = in[1:]
	}
	return out, nil
}

func parseMarker(in string) (size, count int, tail string, err error) {
	end := strings.IndexByte(in, ')')
	if end == -1 {
		return 0, 0, "", fmt.Errorf("missing terminating )")
	}
	in, tail = in[:end], in[end+1:]
	sep := strings.IndexByte(in, 'x')
	if sep == -1 {
		return 0, 0, "", fmt.Errorf("missing separating x in %q", in)
	}
	if size, err = strconv.Atoi(in[:sep]); err != nil {
		return 0, 0, "", fmt.Errorf("invalid size: %w", err)
	}
	if count, err = strconv.Atoi(in[sep+1:]); err != nil {
		return 0, 0, "", fmt.Errorf("invalid count: %w", err)
	}
	return size, count, tail, nil
}
