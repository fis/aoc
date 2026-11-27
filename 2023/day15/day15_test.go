package day15

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"HASH", 52},
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"cm=2", 47},
		{"qp-", 14},
		{"pc=4", 180},
		{"ot=9", 9},
		{"ab=5", 197},
		{"pc-", 48},
		{"pc=6", 214},
		{"ot=7", 231},
	}
	for _, test := range tests {
		if got := hash(test.input); got != test.want {
			t.Errorf("hash(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestInitialize(t *testing.T) {
	steps := []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}
	ops := fn.Map(steps, parseOperation)
	want := 145
	if got := initialize(ops); got != want {
		t.Errorf("initialize(ex) = %d, want %d", got, want)
	}
}
