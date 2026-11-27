package day08

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestLoopCheck(t *testing.T) {
	want := 5
	if prog, err := parseCode(util.Lines(example)); err != nil {
		t.Errorf("parseCode: %v", err)
	} else if loop, acc := loopCheck(prog); !loop || acc != want {
		t.Errorf("loopCheck = %v, %d, want true, %d", loop, acc, want)
	}
}

func TestRepair(t *testing.T) {
	want := 8
	if prog, err := parseCode(util.Lines(example)); err != nil {
		t.Errorf("parseCode: %v", err)
	} else if got := repair(prog); got != want {
		t.Errorf("repair = %d, want %d", got, want)
	}
}
