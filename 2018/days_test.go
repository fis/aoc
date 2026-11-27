package y2018

import (
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	glue.RunTests(t, "../testdata", 2018)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, "../testdata", 2018)
}
