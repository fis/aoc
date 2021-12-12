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

package day12

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestCountAllPaths(t *testing.T) {
	var algos = []struct {
		name string
		f    func(edges [][]string, allowTwice bool) int
	}{
		{
			name: "intGraph",
			f: func(edges [][]string, allowTwice bool) int {
				g := makeIntGraph(edges)
				return g.countAllPaths(allowTwice)
			},
		},
		{
			name: "util.Graph",
			f: func(edges [][]string, allowTwice bool) int {
				g := makeGraph(edges)
				return countAllPaths(g, allowTwice)
			},
		},
	}
	tests := []struct {
		name       string
		edges      [][]string
		allowTwice bool
		want       int
	}{
		{name: "ex1", edges: ex1, allowTwice: false, want: 10},
		{name: "ex2", edges: ex2, allowTwice: false, want: 19},
		{name: "ex3", edges: ex3, allowTwice: false, want: 226},
		{name: "ex1", edges: ex1, allowTwice: true, want: 36},
		{name: "ex2", edges: ex2, allowTwice: true, want: 103},
		{name: "ex3", edges: ex3, allowTwice: true, want: 3509},
	}
	for _, test := range tests {
		for _, alg := range algos {
			if got := alg.f(test.edges, test.allowTwice); got != test.want {
				t.Errorf("countAllPaths[%s](%s, %t) = %d, want %d", alg.name, test.name, test.allowTwice, got, test.want)
			}
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	var algos = []struct {
		name string
		f    func(edges [][]string) (p1, p2 int)
	}{
		{
			name: "intGraph",
			f: func(edges [][]string) (p1, p2 int) {
				g := makeIntGraph(edges)
				p1 = g.countAllPaths(false)
				p2 = g.countAllPaths(true)
				return p1, p2
			},
		},
		{
			name: "util.Graph",
			f: func(edges [][]string) (p1, p2 int) {
				g := makeGraph(edges)
				p1 = countAllPaths(g, false)
				p2 = countAllPaths(g, true)
				return p1, p2
			},
		},
	}
	edges, err := util.ReadRegexp("../days/testdata/day12.txt", inputRegexp)
	if err != nil {
		b.Fatal(err)
	}
	for _, alg := range algos {
		b.Run(alg.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if got1, got2 := alg.f(edges); got1 != 5228 || got2 != 131228 {
					b.Errorf("%s = (%d, %d), want (5228, 131228)", alg.name, got1, got2)
				}
			}
		})
	}
}
