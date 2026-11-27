package day07

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindRoot(t *testing.T) {
	lines := [][]string{
		{"pbga", "66", ""},
		{"xhth", "57", ""},
		{"ebii", "61", ""},
		{"havc", "66", ""},
		{"ktlj", "57", ""},
		{"fwft", "72", "ktlj, cntj, xhth"},
		{"qoyq", "66", ""},
		{"padx", "45", "pbga, havc, qoyq"},
		{"tknk", "41", "ugml, padx, fwft"},
		{"jptl", "61", ""},
		{"ugml", "68", "gyxo, ebii, jptl"},
		{"gyxo", "61", ""},
		{"cntj", "57", ""},
	}
	wantRoot, wantWeight := "tknk", 60
	progs := parseLines(lines)
	g := buildGraph(progs)
	root := findRoot(g)
	if root != wantRoot {
		t.Errorf("findRoot = %s, want %s", root, wantRoot)
	} else if _, w := fixWeight(root, progs); w != wantWeight {
		t.Errorf("fixWeight = %d, want %d", w, wantWeight)
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		line []string
		want program
	}{
		{
			line: []string{"pbga", "66", ""},
			want: program{name: "pbga", weight: 66},
		},
		{
			line: []string{"fwft", "72", "ktlj, cntj, xhth"},
			want: program{name: "fwft", weight: 72, subNames: []string{"ktlj", "cntj", "xhth"}},
		},
	}
	for _, test := range tests {
		if got := parseLine(test.line); !cmp.Equal(got, test.want, cmp.AllowUnexported(program{})) {
			t.Errorf("parseLine(%s) = %#v, want %#v", test.line, got, test.want)
		}
	}
}
