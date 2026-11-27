package day12

import "testing"

var (
	sampleRules = []string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}
	sampleState = "#..#.#..##......###...###"
)

func TestEvolve(t *testing.T) {
	rules := parseRules(sampleRules)
	state := parseState(sampleState)
	state.evolve(20, rules)
	want := 325
	got := state.checksum()
	if got != want {
		t.Errorf("evolve x20 -> %d, want %d", got, want)
	}
}
