package day01

import (
	"testing"
)

func TestModuleFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 12, want: 2},
		{input: 14, want: 2},
		{input: 1969, want: 654},
		{input: 100756, want: 33583},
	}
	for _, test := range tests {
		got := moduleFuel(test.input)
		if got != test.want {
			t.Errorf("moduleFuel(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestModuleFuelComplete(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 14, want: 2},
		{input: 1969, want: 966},
		{input: 100756, want: 50346},
	}
	for _, test := range tests {
		got := moduleFuelComplete(test.input)
		if got != test.want {
			t.Errorf("moduleFuelComplete(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}
