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

// Package day06 solves AoC 2022 day 6.
package day06

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2022, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if n := len(lines); n != 1 {
		return nil, fmt.Errorf("expected 1 line, got %d", n)
	}
	sig := lines[0]
	packet := findMarkerWindowed(sig, 4)
	msg := packet - 4 + findMarkerWindowed(sig[packet-4:], 14)
	return glue.Ints(packet, msg), nil
}

func findMarkerBitset(sig string, size int) int {
	for i := size; i <= len(sig); i++ {
		var b uint64
		for j := i - size; j < i; j++ {
			b |= 1 << (sig[j] & 0x3f)
		}
		if bits.OnesCount64(b) == size {
			return i
		}
	}
	return -1
}

func findMarkerWindowed(sig string, size int) int {
	var (
		counts [64]int
		unique int
	)
	for i := 0; i < size; i++ {
		b := sig[i] & 0x3f
		if counts[b] == 0 {
			unique++
		}
		counts[b]++
	}
	for i := size; i < len(sig); i++ {
		if unique == size {
			return i
		}
		out, in := sig[i-size]&0x3f, sig[i]&0x3f
		counts[out]--
		if counts[out] == 0 {
			unique--
		}
		if counts[in] == 0 {
			unique++
		}
		counts[in]++
	}
	if unique == size {
		return len(sig)
	}
	return -1
}
