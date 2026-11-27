package day05

import (
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "aA", want: ""},
		{input: "abBA", want: ""},
		{input: "abAB", want: "abAB"},
		{input: "aabAAB", want: "aabAAB"},
		{input: "dabAcCaCBAcCcaDA", want: "dabCBAcaDA"},
	}
	for _, test := range tests {
		buf := []byte(test.input)
		buf = reduce(buf)
		got := string(buf)
		if got != test.want {
			t.Errorf("reduce(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestImprove(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	want := 4
	got := improve(input)
	if got != want {
		t.Errorf("improve(%q) = %d, want %d", input, got, want)
	}
}
