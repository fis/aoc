package day06

import (
	"testing"
)

func TestCountFish(t *testing.T) {
	initial := []int{3, 4, 3, 1, 2}
	tests := []struct {
		days int
		want int
	}{
		{days: 1, want: 5},
		{days: 2, want: 6},
		{days: 3, want: 7},
		{days: 4, want: 9},
		{days: 5, want: 10},
		{days: 18, want: 26},
		{days: 80, want: 5934},
		{days: 256, want: 26984457539},
	}
	for _, test := range tests {
		if got := countFish(initial, test.days); got != test.want {
			t.Errorf("countFish(%v, %d) = %d, want %d", initial, test.days, got, test.want)
		}
	}
}
