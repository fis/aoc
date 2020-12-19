// Copyright 2020 Google LLC
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

package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
)

var (
	rules1 = `
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
`
	input1 = `
ababbb
bababa
abbbab
aaabbb
aaaabbb
`
	rules2 = `
42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1
`
	input2 = `
abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
`
)

func TestMatches(t *testing.T) {
	tests := []struct {
		name   string
		rules  []string
		inputs []string
		want   map[string]struct{}
	}{
		{
			name:   "ex1",
			rules:  util.Lines(strings.TrimSpace(rules1)),
			inputs: util.Lines(strings.TrimSpace(input1)),
			want: map[string]struct{}{
				"ababbb": {},
				"abbbab": {},
			},
		},
		{
			name:   "ex2",
			rules:  util.Lines(strings.TrimSpace(rules2)),
			inputs: util.Lines(strings.TrimSpace(input2)),
			want: map[string]struct{}{
				"bbabbbbaabaabba": {},
				"ababaaaaaabaaab": {},
				"ababaaaaabbbaba": {},
			},
		},
	}
	for _, test := range tests {
		rs, err := parseRules(test.rules)
		if err != nil {
			t.Errorf("parseRules(%s): %v", test.name, err)
			continue
		}
		for _, line := range test.inputs {
			_, want := test.want[line]
			got := rs.matches(0, line, false)
			if got != want {
				t.Errorf("%s.matches(0, %s, false) = %v, want %v", test.name, line, got, want)
			}
		}
	}
}

func TestMatchesMagic(t *testing.T) {
	wants := map[string]struct{}{
		"bbabbbbaabaabba":                               {},
		"babbbbaabbbbbabbbbbbaabaaabaaa":                {},
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa": {},
		"bbbbbbbaaaabbbbaaabbabaaa":                     {},
		"bbbababbbbaaaaaaaabbababaaababaabab":           {},
		"ababaaaaaabaaab":                               {},
		"ababaaaaabbbaba":                               {},
		"baabbaaaabbaaaababbaababb":                     {},
		"abbbbabbbbaaaababbbbbbaaaababb":                {},
		"aaaaabbaabaaaaababaa":                          {},
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa":           {},
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba":      {},
	}
	rs, err := parseRules(util.Lines(strings.TrimSpace(rules2)))
	if err != nil {
		t.Errorf("parseRules: %v", err)
		return
	}
	for _, line := range util.Lines(strings.TrimSpace(input2)) {
		_, want := wants[line]
		got := rs.matches(0, line, true)
		if got != want {
			t.Errorf("ex2.matches(0, %s, true) = %v, want %v", line, got, want)
		}
	}
}
