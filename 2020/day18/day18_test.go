package day18

import "testing"

var examples = []struct {
	input            string
	simple, advanced int
}{
	{input: "1 + 2 * 3 + 4 * 5 + 6", simple: 71, advanced: 231},
	{input: "1 + (2 * 3) + (4 * (5 + 6))", simple: 51, advanced: 51},
	{input: "2 * 3 + (4 * 5)", simple: 26, advanced: 46},
	{input: "5 + (8 * 3 + 9 + 3 * 4 * 3)", simple: 437, advanced: 1445},
	{input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", simple: 12240, advanced: 669060},
	{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", simple: 13632, advanced: 23340},
}

func TestSimple(t *testing.T) {
	for _, test := range examples {
		if e, err := parseExpr(&tokenizer{input: test.input}); err != nil {
			t.Errorf("parseExpr(%s): %v", test.input, err)
		} else if got := e.value(); got != test.simple {
			t.Errorf("parseExpr(%s) = %d, want %d", test.input, got, test.simple)
		}
	}
}

func TestAdvanced(t *testing.T) {
	for _, test := range examples {
		if e, err := parseAdvanced(&tokenizer{input: test.input}); err != nil {
			t.Errorf("parseAdvanced(%s): %v", test.input, err)
		} else if got := e.value(); got != test.advanced {
			t.Errorf("parseAdvanced(%s) = %d, want %d", test.input, got, test.advanced)
		}
	}
}
