package day15

import (
	"strings"
	"testing"
)

var (
	ex1 = strings.TrimPrefix(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`, "\n")
	ex2 = strings.TrimPrefix(`
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`, "\n")
)

func TestSumBoxes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ex1", ex1, 2028},
		{"ex2", ex2, 10092},
	}
	for _, test := range tests {
		level, moves := parseInput(strings.NewReader(test.input))
		if got := sumBoxes(level, moves); got != test.want {
			t.Errorf("sumBoxes(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}

func TestSumWideBoxes(t *testing.T) {
	level, moves := parseInput(strings.NewReader(ex2))
	wideLevel := embiggen(level)
	want := 9021
	if got := sumWideBoxes(wideLevel, moves); got != want {
		t.Errorf("sumWideBoxes(ex2) = %d, want %d", got, want)
	}
}
