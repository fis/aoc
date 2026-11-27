package day23

import (
	"testing"
)

var ex = []string{
	"kh-tc",
	"qp-kh",
	"de-cg",
	"ka-co",
	"yn-aq",
	"qp-ub",
	"cg-tb",
	"vc-aq",
	"tb-ka",
	"wh-tc",
	"yn-cg",
	"kh-ub",
	"ta-co",
	"de-co",
	"tc-td",
	"tb-wq",
	"wh-td",
	"ta-ka",
	"td-qp",
	"aq-cg",
	"wq-ub",
	"ub-vc",
	"de-ta",
	"wq-aq",
	"wq-vc",
	"wh-yn",
	"ka-de",
	"kh-ta",
	"co-tc",
	"wh-qp",
	"tb-vc",
	"td-yn",
}

func TestCountCliqueT3(t *testing.T) {
	nodes, neigh := buildGraph(ex)
	want := 7
	if got := countCliqueT3(nodes, neigh); got != want {
		t.Errorf("findCandidateClusters(ex) = %d, want %d", got, want)
	}
}

func TestMaxClique(t *testing.T) {
	nodes, neigh := buildGraph(ex)
	want := "co,de,ka,ta"
	if got := maxClique(nodes, neigh); got != want {
		t.Errorf("maxClique(ex) = %q, want %q", got, want)
	}
}
