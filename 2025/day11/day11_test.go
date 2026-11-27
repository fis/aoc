package day11

import "testing"

var (
	ex1 = []string{
		"you: bbb ccc",
		"bbb: ddd eee",
		"ccc: ddd eee fff",
		"ddd: ggg",
		"eee: out",
		"fff: out",
		"ggg: out",
		"hhh: ccc fff iii",
		"iii: out",
	}
	ex2 = []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}
)

func TestCountPaths(t *testing.T) {
	g := buildGraph(ex1)
	want := 5
	if got := newPathCounter(g, "you").countTo("out"); got != want {
		t.Errorf("pc(ex1, you).countTo(out) = %d, want %d", got, want)
	}
}

func TestCountValidPaths(t *testing.T) {
	g := buildGraph(ex2)
	want := 2
	if got := newValidPathCounter(g, "svr").countTo("out"); got != want {
		t.Errorf("vpc(ex2, svr).countTo(out) = %d, want %d", got, want)
	}
}
