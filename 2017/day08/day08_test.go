package day08

import (
	"testing"
)

func TestMaxRegs(t *testing.T) {
	code := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	wantFinal, wantEver := 1, 10
	if gotFinal, gotEver, err := maxRegs(code); err != nil {
		t.Errorf("part1: %v", err)
	} else if gotFinal != wantFinal || gotEver != wantEver {
		t.Errorf("part1 = (%d, %d), want (%d, %d)", gotFinal, gotEver, wantFinal, wantEver)
	}
}
