package day16

import (
	"testing"

	"github.com/fis/aoc/2018/cpu"
)

var example = []string{
	"Before: [3, 2, 1, 1]",
	"9 2 1 2",
	"After:  [3, 2, 2, 1]",
}

func TestValidOps(t *testing.T) {
	want := (uint(1) << cpu.AddI) | (uint(1) << cpu.MulR) | (uint(1) << cpu.SetI)
	if s, err := parseSample(example); err != nil {
		t.Errorf("parseSample: %v", err)
	} else if got := validOps(s); got != want {
		t.Errorf("validOps(%v) = %b, want %b", s, got, want)
	}
}
