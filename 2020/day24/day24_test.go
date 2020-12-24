// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package day24

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
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
	for i := 0; i < 100; i++ {
		m = m.evolve()
	}

	want := 2208
	got := m.countBlack()
	if want != got {
		t.Errorf("evolve x100 -> countBlack() = %d, want %d", got, want)
	}
}
