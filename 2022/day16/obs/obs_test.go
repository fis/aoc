package obs

import (
	"testing"

	"github.com/fis/aoc/2022/day16"
	"github.com/fis/aoc/util/fn"
)

func TestReleasePressure(t *testing.T) {
	scan, err := fn.MapE(day16.ExampleScan, day16.ParseValveScan)
	if err != nil {
		t.Fatal(err)
	}
	sum := day16.Preprocess(scan)
	want := 1651
	if got := releasePressure(sum, 30); got != want {
		t.Errorf("releasePressure(ex, 30) = %d, want %d", got, want)
	}
}

func TestReleasePressure2(t *testing.T) {
	scan, err := fn.MapE(day16.ExampleScan, day16.ParseValveScan)
	if err != nil {
		t.Fatal(err)
	}
	sum := day16.Preprocess(scan)
	want := 1707
	if got := releasePressure2(sum, 26); got != want {
		t.Errorf("releasePressure2(ex, 26) = %d, want %d", got, want)
	}
}
