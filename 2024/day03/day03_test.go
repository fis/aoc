package day03

import (
	"os"
	"testing"
)

var (
	ex1   = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	ex2   = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	algos = []struct {
		name string
		impl func(string) (uncondSum, condSum int)
	}{
		{name: "regexp", impl: findInstructionsRegexp},
		{name: "string", impl: findInstructionsString},
	}
)

func TestFindInstructions(t *testing.T) {
	want1, want2 := 161, 48
	for _, algo := range algos {
		got1, _ := algo.impl(ex1)
		_, got2 := algo.impl(ex2)
		if got1 != want1 {
			t.Errorf("findInstructions[%s](ex1) = %d, _, want %d", algo.name, got1, want1)
		}
		if got2 != want2 {
			t.Errorf("findInstructions[%s](ex2) = _, %d, want %d", algo.name, got2, want2)
		}

	}
}

func BenchmarkAlgos(b *testing.B) {
	data, err := os.ReadFile("../../testdata/2024/day03.txt")
	if err != nil {
		b.Fatal(err)
	}
	mem := string(data)
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for range b.N {
				if got1, got2 := algo.impl(mem); got1 != 190604937 || got2 != 82857512 {
					b.Errorf("%s = (%d, %d), want (190604937, 82857512)", algo.name, got1, got2)
				}
			}
		})
	}
}
