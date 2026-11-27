package day10

import (
	"testing"
)

var (
	ex1 = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	ex2 = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
)

func TestDeltas(t *testing.T) {
	tests := []struct {
		name         string
		adapters     []int
		want1, want3 int
	}{
		{name: "ex1", adapters: ex1, want1: 7, want3: 5},
		{name: "ex2", adapters: ex2, want1: 22, want3: 10},
	}
	for _, test := range tests {
		got1, got3 := deltas(test.adapters)
		if got1 != test.want1 || got3 != test.want3 {
			t.Errorf("deltas(%s) = (%d, %d), want (%d, %d)", test.name, got1, got3, test.want1, test.want3)
		}
	}
}

func TestArrangements(t *testing.T) {
	tests := []struct {
		name     string
		adapters []int
		want     int
	}{
		{name: "ex1", adapters: ex1, want: 8},
		{name: "ex2", adapters: ex2, want: 19208},
	}
	for _, test := range tests {
		got := arrangements(test.adapters)
		if got != test.want {
			t.Errorf("arrangements(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}
