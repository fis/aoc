package day19

import (
	"testing"
)

var (
	ex1 = blueprint{4, 2, 3, 14, 2, 7}
	ex2 = blueprint{2, 3, 3, 8, 3, 12}
)

func TestMaxGeodes(t *testing.T) {
	tests := []struct {
		bp      blueprint
		maxTime int
		want    int
	}{
		{ex1, 24, 9},
		{ex2, 24, 12},
		{ex1, 32, 56},
		{ex2, 32, 62},
	}
	for _, test := range tests {
		if got := maxGeodes(test.bp, test.maxTime); got != test.want {
			t.Errorf("maxGeodes(%v, %d) = %d, want %d", test.bp, test.maxTime, got, test.want)
		}
	}
}
