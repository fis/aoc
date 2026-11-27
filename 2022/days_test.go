package y2022

import (
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	glue.RunTests(t, "../testdata", 2022)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, "../testdata", 2022)
}
