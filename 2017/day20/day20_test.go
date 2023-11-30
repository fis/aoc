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

package day20

import (
	"bufio"
	"os"
	"regexp"
	"testing"

	"github.com/fis/aoc/util"
)

func TestClosest(t *testing.T) {
	ps := []particle{
		{p3{3, 0, 0}, p3{2, 0, 0}, p3{-1, 0, 0}},
		{p3{4, 0, 0}, p3{0, 0, 0}, p3{-2, 0, 0}},
	}
	want := 0
	if got := closest(ps); got != want {
		t.Errorf("closest = %d, want %d", got, want)
	}
}

func TestCollide(t *testing.T) {
	tests := []struct {
		name string
		f    func([]particle) int
	}{
		{name: "sim", f: func(ps []particle) int { return collideSim(ps, 5) }},
		{name: "calc", f: collideCalc},
	}
	for _, test := range tests {
		ps := []particle{
			{p3{-6, 0, 0}, p3{3, 0, 0}, p3{0, 0, 0}},
			{p3{-4, 0, 0}, p3{2, 0, 0}, p3{0, 0, 0}},
			{p3{-2, 0, 0}, p3{1, 0, 0}, p3{0, 0, 0}},
			{p3{3, 0, 0}, p3{-1, 0, 0}, p3{0, 0, 0}},
		}
		want := 1
		if got := test.f(ps); got != want {
			t.Errorf("collide(%s) = %d, want %d", test.name, got, want)
		}

	}
}

func TestCollide1D(t *testing.T) {
	tests := []struct {
		p1, v1, a1, p2, v2, a2 int
		want1, want2           int
	}{
		{0, 1, 0, 6, -1, 0, 3, -1},
		{0, 1, 0, 5, -1, 0, -1, -1},
		{0, 0, 1, -3, 3, 0, 2, 3},
	}
	for _, test := range tests {
		got1, got2, _ := collide1D(test.p1, test.v1, test.a1, test.p2, test.v2, test.a2)
		if got1 != test.want1 || got2 != test.want2 {
			t.Errorf(
				"collide1D(%d, %d, %d, %d, %d, %d) = (%d, %d), want (%d, %d)",
				test.p1, test.v1, test.a1, test.p2, test.v2, test.a2,
				got1, got2, test.want1, test.want2,
			)
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	f, err := os.Open("../../testdata/2017/day20.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	re, err := regexp.Compile(inputRegexp)
	if err != nil {
		b.Fatal(err)
	}
	lines, err := util.ScanAll(f, bufio.ScanLines)
	if err != nil {
		b.Fatal(err)
	}
	parsed := make([][]string, len(lines))
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		if parts == nil {
			b.Fatalf("line %q does not match pattern %s", line, inputRegexp)
		}
		parsed[i] = parts[1:]
	}
	ps := parseInput(parsed)
	algos := []struct {
		name string
		f    func([]particle) int
	}{
		{name: "sim", f: func(ps []particle) int { return collideSim(ps, 40) }},
		{name: "calc", f: collideCalc},
	}
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ps2 := append([]particle(nil), ps...)
				want := 574
				if got := algo.f(ps2); got != want {
					b.Errorf("got %d, want %d", got, want)
				}
			}
		})
	}
}
