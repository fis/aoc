package day16

import (
	"testing"
)

func TestPerform(t *testing.T) {
	dance := []string{"s1", "x3/4", "pe/b"}
	want := "baedc"
	if moves, err := parseMoves(dance); err != nil {
		t.Errorf("parseMoves(%v): %v", dance, err)
	} else if got := perform(5, moves); got != want {
		t.Errorf("perform = %s, want %s", got, want)
	}
}

func TestEncore(t *testing.T) {
	dance := []string{"s1", "x3/4", "pe/b"}
	moves, err := parseMoves(dance)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		times int
		want  string
	}{
		{times: 1, want: "baedc"},
		{times: 2, want: "ceadb"},
		{times: 1000000000, want: "abcde"},
	}
	for _, test := range tests {
		if got := encore(5, moves, test.times); got != test.want {
			t.Errorf("encore(..., %d) = %s, want %s", test.times, got, test.want)
		}
	}
}
