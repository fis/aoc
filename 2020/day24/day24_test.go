package day24

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var example = `
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
`

func TestFlip(t *testing.T) {
	paths, err := parsePaths(util.Lines(strings.TrimPrefix(example, "\n")))
	if err != nil {
		t.Fatalf("parsePaths: %v", err)
	}

	m := newTileMap()
	m.flipAll(paths)

	want := 10
	got := m.countBlack()
	if want != got {
		t.Errorf("countBlack() = %d, want %d", got, want)
	}
}

func TestEvolve(t *testing.T) {
	paths, err := parsePaths(util.Lines(strings.TrimPrefix(example, "\n")))
	if err != nil {
		t.Fatalf("parsePaths: %v", err)
	}

	m := newTileMap()
	m.flipAll(paths)
	for range 100 {
		m = m.evolve()
	}

	want := 2208
	got := m.countBlack()
	if want != got {
		t.Errorf("evolve x100 -> countBlack() = %d, want %d", got, want)
	}
}
