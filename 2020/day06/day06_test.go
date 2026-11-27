package day06

import (
	"testing"
)

var example = [][]string{
	{"abc"},
	{"a", "b", "c"},
	{"ab", "ac"},
	{"a", "a", "a", "a"},
	{"b"},
}

func TestCountMerged(t *testing.T) {
	tests := []struct {
		name   string
		merger func([]answerSet) answerSet
		want   []int
	}{
		{"any", mergeAny, []int{3, 3, 3, 1, 1}},
		{"all", mergeAll, []int{3, 0, 1, 1, 1}},
	}
	for _, test := range tests {
		for i, group := range example {
			answers := parseGroup(group)
			got := countMerged(answers, test.merger)
			if got != test.want[i] {
				t.Errorf("countMerged(%v, %s) = %d, want %d", group, test.name, got, test.want[i])
			}
		}
	}
}
