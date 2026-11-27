package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		square int
		want   int
	}{
		{square: 1, want: 0},
		{square: 12, want: 3},
		{square: 23, want: 2},
		{square: 1024, want: 31},
	}
	for _, test := range tests {
		if got := part1(test.square); got != test.want {
			t.Errorf("part1(%d) = %d, want %d", test.square, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{limit: 50, want: 54},
		{limit: 55, want: 57},
		{limit: 100, want: 122},
		{limit: 300, want: 304},
		{limit: 500, want: 747},
		{limit: 800, want: 806},
	}
	for _, test := range tests {
		if got := part2(test.limit); got != test.want {
			t.Errorf("part2(%d) = %d, want %d", test.limit, got, test.want)
		}
	}
}
