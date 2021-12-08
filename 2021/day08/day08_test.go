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

package day08

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{input: []string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab", "cdfeb fcadb cdfeb cdbaf"}, want: 5353},
		{input: []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb", "fdgacbe cefdb cefbgd gcbe"}, want: 8394},
		{input: []string{"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec", "fcgedb cgb dgebacf gc"}, want: 9781},
		{input: []string{"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef", "cg cg fdcagb cbg"}, want: 1197},
		{input: []string{"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega", "efabcd cedba gadfec cb"}, want: 9361},
		{input: []string{"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga", "gecf egdcabf bgf bfgea"}, want: 4873},
		{input: []string{"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf", "gebdcfa ecba ca fadegcb"}, want: 8418},
		{input: []string{"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf", "cefg dcbef fcge gbcadfe"}, want: 4548},
		{input: []string{"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd", "ed bcgafe cdgba cbgef"}, want: 1625},
		{input: []string{"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg", "gbdfcae bgc cg cgb"}, want: 8717},
		{input: []string{"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc", "fgae cfgab fg bagce"}, want: 4315},
	}
	for _, test := range tests {
		if data, err := parseInput([][]string{test.input}); err != nil {
			t.Errorf("parseInput(%v): %v", test.input, err)
		} else if _, got := data[0].decode(); got != test.want {
			t.Errorf("(%v).decode() = %d, want %d", data[0], got, test.want)
		}
	}
}
