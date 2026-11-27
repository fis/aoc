package day22

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
..#
#..
...
`

func TestSimulateSimple(t *testing.T) {
	tests := []struct {
		rounds int
		want   int
	}{
		{rounds: 7, want: 5},
		{rounds: 70, want: 41},
		{rounds: 10000, want: 5587},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
		if got := simulateSimple(level, test.rounds); got != test.want {
			t.Errorf("simulateSimple(..., %d) = %d, want %d", test.rounds, got, test.want)
		}
	}
}

func TestSimulateEvolved(t *testing.T) {
	tests := []struct {
		rounds int
		want   int
	}{
		{rounds: 7, want: 1},
		{rounds: 100, want: 26},
		{rounds: 10000000, want: 2511944},
	}
	for _, test := range tests {
		for i, f := range []func(*util.Level, int) int{simulateEvolvedLevel, simulateEvolvedArray} {
			level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
			if got := f(level, test.rounds); got != test.want {
				t.Errorf("simulateEvolved[%d](..., %d) = %d, want %d", i, test.rounds, got, test.want)
			}
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	algos := []struct {
		name string
		f    func(*util.Level, int) int
	}{
		{name: "level", f: simulateEvolvedLevel},
		{name: "array", f: simulateEvolvedArray},
	}
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for range b.N {
				level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
				want := 2511944
				if got := algo.f(level, 10000000); got != want {
					b.Errorf("got %d, want %d", got, want)
				}
			}
		})
	}
}
