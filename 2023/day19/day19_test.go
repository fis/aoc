// Package day19 solves AoC 2023 day 19.
package day19

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestAccepts(t *testing.T) {
	workflows, parts, err := parseInput(util.Lines(ex))
	if err != nil {
		t.Fatal(err)
	}
	wants := []bool{true, false, true, false, true} // parts in example
	for i, want := range wants {
		if got := workflows.accepts(parts[i]); got != want {
			t.Errorf("(ex).accepts(%v) = %t, want %t", parts[i], got, want)
		}
	}
}

func TestCountAccepted(t *testing.T) {
	workflows, _, err := parseInput(util.Lines(ex))
	if err != nil {
		t.Fatal(err)
	}
	want := 167409079868000
	if got := workflows.countAccepted(); got != want {
		t.Errorf("(ex).countAccepted() = %d, want %d", got, want)
	}
}
