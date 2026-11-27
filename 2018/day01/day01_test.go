package day01

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		changes []int
		want    int
	}{
		{changes: []int{1, -2, 3, 1}, want: 3},
		{changes: []int{1, 1, 1}, want: 3},
		{changes: []int{1, 1, -2}, want: 0},
		{changes: []int{-1, -2, -3}, want: -6},
	}
	for _, test := range tests {
		got := sum(test.changes)
		if got != test.want {
			t.Errorf("sum(%v) = %d, want %d", test.changes, got, test.want)
		}
	}
}

func TestFindRep(t *testing.T) {
	tests := []struct {
		changes []int
		want    int
	}{
		{changes: []int{1, -2, 3, 1}, want: 2},
		{changes: []int{1, -1}, want: 0},
		{changes: []int{3, 3, 4, -2, -4}, want: 10},
		{changes: []int{-6, 3, 8, 5, -6}, want: 5},
		{changes: []int{7, 7, -2, -7, 4}, want: 14},
	}
	for _, test := range tests {
		got := findRep(test.changes)
		if got != test.want {
			t.Errorf("findRep(%v) = %d, want %d", test.changes, got, test.want)
		}
	}
}
