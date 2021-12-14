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

package day14

import "testing"

var ex = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

var tests = []struct {
	steps int
	want  int
}{
	{steps: 10, want: 1588},
	{steps: 40, want: 2188189693529},
}

func TestPairRules(t *testing.T) {
	rb, initial, ends, err := parsePairRules(ex)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		pairs := rb.updateCountsN(initial, test.steps)
		quants := rb.countElements(pairs, ends)
		if got := maxDiff(quants); got != test.want {
			t.Errorf("%d steps -> %d, want %d", test.steps, got, test.want)
		}
	}
}

func TestElementCounter(t *testing.T) {
	rb, initial, err := parseRules(ex)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		ec := elementCounter{}
		quants := ec.countPolymer(initial, test.steps, rb)
		if got := maxDiff(quants); got != test.want {
			t.Errorf("%d steps -> %d, want %d", test.steps, got, test.want)
		}
	}
}

func TestExpand(t *testing.T) {
	rb, initial, err := parseRules(ex)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		steps int
		want  string
	}{
		{steps: 0, want: "NNCB"},
		{steps: 1, want: "NCNBCHB"},
		{steps: 2, want: "NBCCNBBBCBHCB"},
		{steps: 3, want: "NBBBCNCCNBBNBNBBCHBHHBCHB"},
		{steps: 4, want: "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"},
	}
	for _, test := range tests {
		if got := rb.asString(rb.expandN(initial, test.steps)); got != test.want {
			t.Errorf("expandN(..., %d) = %s, want %s", test.steps, got, test.want)
		}
	}
}
