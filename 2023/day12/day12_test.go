package day12

import "testing"

var tests = []struct {
	row          string
	groups       []int
	ways         int
	unfoldedWays int
}{
	{"???.###", []int{1, 1, 3}, 1, 1},
	{".??..??...?##.", []int{1, 1, 3}, 4, 16384},
	{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1, 1},
	{"????.#...#...", []int{4, 1, 1}, 1, 16},
	{"????.######..#####.", []int{1, 6, 5}, 4, 2500},
	{"?###????????", []int{3, 2, 1}, 10, 506250},
}

func TestCountWays(t *testing.T) {
	for _, test := range tests {
		if got := countWays([]byte(test.row), test.groups); got != test.ways {
			t.Errorf("countWays(%q, %v) = %d, want %d", test.row, test.groups, got, test.ways)
		}
	}
}

func TestCountWaysUnfolded(t *testing.T) {
	for _, test := range tests {
		r := unfold([]record{{[]byte(test.row), test.groups}}, 5)[0]
		if got := countWays(r.row, r.groups); got != test.unfoldedWays {
			t.Errorf("countWays(%q*5, %v*5) = %d, want %d", test.row, test.groups, got, test.unfoldedWays)
		}
	}
}
