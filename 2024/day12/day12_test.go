// Copyright 2024 Google LLC
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
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var (
	ex1 = strings.TrimPrefix(`
AAAA
BBCD
BBCC
EEEC
	`, "\n")
	ex2 = strings.TrimPrefix(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
	`, "\n")
	ex3 = strings.TrimPrefix(`
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
	`, "\n")
	ex4 = strings.TrimPrefix(`
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
	`, "\n")
	ex5 = strings.TrimPrefix(`
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
	`, "\n")
)

func TestFencingPrice(t *testing.T) {
	tests := []struct {
		name         string
		garden       string
		want1, want2 int
	}{
		{"ex1", ex1, 140, 80},
		{"ex2", ex2, 772, 436},
		{"ex3", ex3, 1930, 1206},
		{"ex4", ex4, 692, 236},
		{"ex5", ex5, 1184, 368},
	}
	for _, test := range tests {
		level := util.ParseFixedLevel([]byte(test.garden))
		if got1, got2 := fencingPrice(level); got1 != test.want1 || got2 != test.want2 {
			t.Errorf("fencingPrice(%s) = (%d,%d), want (%d,%d)", test.name, got1, got2, test.want1, test.want2)
		}
	}
}
