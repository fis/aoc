package y2024

import (
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	glue.RunTests(t, "../testdata", 2024)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, "../testdata", 2024)
}
