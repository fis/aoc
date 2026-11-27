package day03

import (
	"testing"
)

var ex = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestGammaEpsilon(t *testing.T) {
	want1, want2 := 22, 9
	if got1, got2 := gammaEpsilon(ex); got1 != want1 || got2 != want2 {
		t.Errorf("gammaEpsilon = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}

func TestFilterBits(t *testing.T) {
	tests := []struct {
		keepLCB int
		want    int
	}{
		{0, 23},
		{1, 10},
	}
	for _, test := range tests {
		if got := filterBits(ex, test.keepLCB); got != test.want {
			t.Errorf("filterBits(..., %d) = %d, want %d", test.keepLCB, got, test.want)
		}
	}
}
