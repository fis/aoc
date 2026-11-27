package day10

import (
	"testing"

	"github.com/fis/aoc/util/fn"
	"github.com/google/go-cmp/cmp"
)

var ex = []string{
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

func TestConfigLights(t *testing.T) {
	machines, err := fn.MapE(ex, parseMachine)
	if err != nil {
		t.Fatalf("parseMachine(ex): %v", err)
	}
	want := []int{2, 3, 2}
	if got := fn.Map(machines, machine.configLights); !cmp.Equal(got, want) {
		t.Errorf("(ex).configLights = %v, want %v", got, want)
	}
}

func TestConfigJolts(t *testing.T) {
	machines, err := fn.MapE(ex, parseMachine)
	if err != nil {
		t.Fatalf("parseMachine(ex): %v", err)
	}
	want := []int{10, 12, 11}
	if got := fn.Map(machines, machine.configJolts); !cmp.Equal(got, want) {
		t.Errorf("(ex).configJolts = %v, want %v", got, want)
	}
}
