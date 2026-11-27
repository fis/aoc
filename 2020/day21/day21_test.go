package day21

import "testing"

var example = []string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
}

func TestAnalyze(t *testing.T) {
	want1, want2 := 5, "mxmxvkd,sqjhc,fvjkl"
	if labels, err := parseInput(example); err != nil {
		t.Errorf("parseInput: %v", err)
	} else if got1, got2 := analyze(labels); got1 != want1 || got2 != want2 {
		t.Errorf("part1 = (%d,%q), want (%d,%q)", got1, got2, want1, want2)
	}
}
