// Copyright 2023 Google LLC
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

package day04

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/google/go-cmp/cmp"
)

var (
	ex = []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	exCards = []card{
		{winners: []byte{41, 48, 83, 86, 17}, numbers: []byte{83, 86, 6, 31, 17, 9, 48, 53}},
		{winners: []byte{13, 32, 20, 16, 61}, numbers: []byte{61, 30, 68, 82, 17, 32, 24, 19}},
		{winners: []byte{1, 21, 53, 59, 44}, numbers: []byte{69, 82, 63, 72, 16, 21, 14, 1}},
		{winners: []byte{41, 92, 73, 84, 69}, numbers: []byte{59, 84, 76, 51, 58, 5, 54, 83}},
		{winners: []byte{87, 83, 26, 28, 32}, numbers: []byte{88, 30, 70, 12, 93, 22, 82, 36}},
		{winners: []byte{31, 18, 13, 56, 72}, numbers: []byte{74, 77, 10, 23, 35, 67, 36, 11}},
	}
)

func TestCountPoints(t *testing.T) {
	cards, err := fn.MapE(ex, parseCardFast)
	if err != nil {
		t.Fatal(err)
	}
	want := 13
	if got := countPoints(cards); got != want {
		t.Errorf("countPoints(ex) = %d, want %d", got, want)
	}
}

func TestCountCards(t *testing.T) {
	cards, err := fn.MapE(ex, parseCardFast)
	if err != nil {
		t.Fatal(err)
	}
	want := 30
	if got := countCards(cards); got != want {
		t.Errorf("countCards(ex) = %d, want %d", got, want)
	}
}

var parsers = []struct {
	name string
	f    func(string) (card, error)
}{
	{"simple", parseCardSimple},
	{"fast", parseCardFast},
}

func TestParseCard(t *testing.T) {
	for _, p := range parsers {
		for i, line := range ex {
			want := exCards[i]
			if got, err := p.f(line); err != nil {
				t.Errorf("%s(%s): %v", p.name, line, err)
			} else if !cmp.Equal(got.winners, want.winners) || !cmp.Equal(got.numbers, want.numbers) {
				t.Errorf("%s(%s) = %v, want %v", p.name, line, got, want)
			}
		}
	}
}

func BenchmarkParseCard(b *testing.B) {
	lines, err := util.ReadLines("../../testdata/2023/day04.txt")
	if err != nil {
		b.Fatal(err)
	}
	for _, p := range parsers {
		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, line := range lines {
					if _, err := p.f(line); err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}
