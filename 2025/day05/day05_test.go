package day05

import "testing"

var (
	exFresh       = []interval{{3, 5}, {10, 14}, {16, 20}, {12, 18}}
	exIngredients = []int{1, 5, 8, 11, 17, 32}
)

func TestCountFresh(t *testing.T) {
	want1, want2 := 3, 14
	if got1, got2 := countFresh(exFresh, exIngredients); got1 != want1 || got2 != want2 {
		t.Errorf("countFresh(%v, %v) = (%d, %d), want (%d, %d)", exFresh, exIngredients, got1, got2, want1, want2)
	}
}
