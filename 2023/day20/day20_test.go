package day20

import "testing"

var (
	ex1 = []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}
	ex2 = []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}
)

func TestStepN(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"ex1", ex1, 32000000},
		{"ex2", ex2, 11687500},
	}
	for _, test := range tests {
		g, _, err := parseGraph(test.input)
		if err != nil {
			t.Errorf("parseGraph(%s): %v", test.name, err)
		} else if got := stepN(g, 1000); got != test.want {
			t.Errorf("stepN(%s, 1000) = %d, want %d", test.name, got, test.want)
		}
	}
}
