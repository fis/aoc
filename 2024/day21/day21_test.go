package day21

import (
	"testing"
)

func TestCodeComplexity(t *testing.T) {
	tests := []struct {
		code string
		want int
	}{
		{"029A", 68 * 29},
		{"980A", 60 * 980},
		{"179A", 68 * 179},
		{"456A", 64 * 456},
		{"379A", 64 * 379},
	}
	cc := makeCalculator()
	for _, test := range tests {
		code, _ := parseCode(test.code)
		if got := cc.codeComplexity(code, 2); got != test.want {
			t.Errorf("cc.codeComplexity(%s, 2) = %d, want %d", test.code, got, test.want)
		}
	}
}
