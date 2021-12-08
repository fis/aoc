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

// Package day08 solves AoC 2021 day 8.
package day08

import (
	"bytes"
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 8, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^([a-g]+(?: [a-g]+)*) \| ([a-g]+(?: [a-g]+)*)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	data, err := parseInput(lines)
	if err != nil {
		return nil, err
	}
	p1, p2 := decodeAll(data)
	return glue.Ints(p1, p2), nil
}

type entry struct {
	samples [10][]byte
	code    [4][]byte
}

func decodeAll(data []entry) (sum1478, sumCodes int) {
	for i := range data {
		count1478, code := data[i].decode()
		sum1478 += count1478
		sumCodes += code
	}
	return sum1478, sumCodes
}

// Logic for decoding the segments:
//
//  aa    ..    aa    aa    ..    aa    aa    aa    aa    aa
// b  c  .  c  .  c  .  c  b  c  b  .  b  .  .  c  b  c  b  c
//  ..    ..    dd    dd    dd    dd    dd    ..    dd    dd
// e  f  .  f  e  .  .  f  .  f  .  f  e  f  .  f  e  f  .  f
//  gg    ..    gg    gg    ..    gg    gg    ..    gg    gg
//
//  00    ..    00    00    ..    00    00    00    00    00
// 1  2  .  2  .  2  .  2  1  2  1  .  1  .  .  2  1  2  1  2
//  ..    ..    33    33    33    33    33    ..    33    33
// 4  5  .  5  4  .  .  5  .  5  .  5  4  5  .  5  4  5  .  5
//  66    ..    66    66    ..    66    66    ..    66    66
//
// - Figure out which samples are 1/4/7/8 based on the unique segment counts (2/4/3/7).
// - Figure out what maps to a/0: it's the one that's in 7 but not in 1.
// - Disambiguate c/2 and f/5 (the two segments of '1') by frequency:
//   - c/2 appears 8 times (01234789)
//   - f/5 appears 9 times (all but 2)
// - Disambiguate b/1 and d/3 (the two segments of '4' not in '1') by frequency:
//   - b/1 appears 6 times (045689)
//   - d/3 appears 7 times (2345689)
// - Disambiguate e/4 and g/6 (two segments of '8' that are neither in '4' nor '7') by frequency:
//   - e/4 appears 4 times (0268)
//   - g/6 appears 7 times (0235689)

var digitMap = [...]byte{
	0b1110111: 0, // abcedf  / 012456
	0b0100100: 1, // cf      / 25
	0b1011101: 2, // acdeg   / 02346
	0b1101101: 3, // acdfg   / 02356
	0b0101110: 4, // bcdf    / 1235
	0b1101011: 5, // abdfg   / 01356
	0b1111011: 6, // abdefg  / 013456
	0b0100101: 7, // acf     / 025
	0b1111111: 8, // abcdefg / 0123456
	0b1101111: 9, // abcdfg  / 012356
}

func (e *entry) decode() (count1478, code int) {
	var (
		s1, s4, s7, s8 []byte
		freqs          [7]int
		remap          [7]byte
	)

	for _, s := range e.samples {
		switch len(s) {
		case 2:
			s1 = s
		case 3:
			s7 = s
		case 4:
			s4 = s
		case 7:
			s8 = s
		}
		for _, seg := range s {
			freqs[seg]++
		}
	}

	for i := range remap {
		remap[i] = 0xff
	}
	remap[diff1(s7, s1)] = 0
	if freqs[s1[0]] == 8 {
		remap[s1[0]] = 2
		remap[s1[1]] = 5
	} else {
		remap[s1[0]] = 5
		remap[s1[1]] = 2
	}
	if seg13 := diff(s4, s1); freqs[seg13[0]] == 6 {
		remap[seg13[0]] = 1
		remap[seg13[1]] = 3
	} else {
		remap[seg13[0]] = 3
		remap[seg13[1]] = 1
	}
	if seg46 := diff(diff(s8, s4), s7); freqs[seg46[0]] == 4 {
		remap[seg46[0]] = 4
		remap[seg46[1]] = 6
	} else {
		remap[seg46[0]] = 6
		remap[seg46[1]] = 4
	}

	for _, c := range e.code {
		bits := 0
		for _, b := range c {
			bits |= 1 << remap[b]
		}
		digit := digitMap[bits]
		if digit == 1 || digit == 4 || digit == 7 || digit == 8 {
			count1478++
		}
		code = code*10 + int(digit)
	}

	return count1478, code
}

func diff(as []byte, bs []byte) (d []byte) {
	for _, a := range as {
		if bytes.IndexByte(bs, a) == -1 {
			d = append(d, a)
		}
	}
	return d
}

func diff1(as []byte, bs []byte) byte {
	for _, a := range as {
		if bytes.IndexByte(bs, a) == -1 {
			return a
		}
	}
	return 0xff
}

func parseInput(lines [][]string) ([]entry, error) {
	data := make([]entry, len(lines))
	for i, line := range lines {
		samples := util.Words(line[0])
		code := util.Words(line[1])
		if len(samples) != 10 || len(code) != 4 {
			return nil, fmt.Errorf("unexpected number of words in %v", line)
		}
		parseSegments(data[i].samples[:], samples)
		parseSegments(data[i].code[:], code)
	}
	return data, nil
}

func parseSegments(out [][]byte, in []string) {
	for i, s := range in {
		out[i] = make([]byte, len(s))
		for j := range out[i] {
			out[i][j] = s[j] - 'a'
		}
	}
}
