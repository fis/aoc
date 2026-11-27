package day14

import (
	"testing"
)

func TestHash(t *testing.T) {
	key := "flqrgnkx"
	wantUsed, wantRegions := 8108, 1242

	hashes, gotUsed := hash(key)
	lvl := buildLevel(hashes)
	gotRegions := countRegions(lvl)

	if gotUsed != wantUsed {
		t.Errorf("part1(%q) = %d, want %d", key, gotUsed, wantUsed)
	}
	if gotRegions != wantRegions {
		t.Errorf("part2(%q) = %d, want %d", key, gotRegions, wantRegions)
	}
}
