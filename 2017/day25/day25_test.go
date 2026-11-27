package day25

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
`

func TestRun(t *testing.T) {
	blocks, err := util.ScanAll(strings.NewReader(strings.TrimPrefix(ex, "\n")), util.ScanChunks)
	if err != nil {
		t.Fatal(err)
	}
	tm, steps, err := parseTM(blocks)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	tm.run(steps)
	if got := tm.tape.popCount(); got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}
