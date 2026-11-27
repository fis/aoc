package day17

import (
	"strings"
	"testing"
)

var ex = strings.TrimPrefix(`
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`, "\n")

func TestInterp(t *testing.T) {
	regs, prog := parseInput(strings.NewReader(ex))
	want := "4,6,3,5,6,3,5,2,1,0"
	if got := interp(regs, prog); got != want {
		t.Errorf("interp(ex) = %q, want %q", got, want)
	}
}
