package day06

import (
	"testing"
)

func TestExample1(t *testing.T) {
	data := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	om := parseOrbits(data)
	count := om.countOrbits()
	if count != 42 {
		t.Errorf("countOrbits(%v) = %d, want 42", data, count)
	}
}

func TestExample2(t *testing.T) {
	data := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	om := parseOrbits(data)
	dist := om.transfers("YOU", "SAN")
	if dist != 4 {
		t.Errorf("transfers(%v, YOU, SAN) = %d, want 4", data, dist)
	}
}
