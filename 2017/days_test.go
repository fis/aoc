package y2017

import (
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	glue.RunTests(t, "../testdata", 2017)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, "../testdata", 2017)
}
