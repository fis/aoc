package day22

import "testing"

var (
	deck1 = []string{"Player 1:", "9", "2", "6", "3", "1"}
	deck2 = []string{"Player 2:", "5", "8", "4", "7", "10"}
)

func TestGame(t *testing.T) {
	p1, _ := parseDeck(1, deck1)
	p2, _ := parseDeck(2, deck2)
	want := 306
	got := game(p1, p2)
	if got != want {
		t.Errorf("game = %d, want %d", got, want)
	}
}

func TestRecursive(t *testing.T) {
	p1, _ := parseDeck(1, deck1)
	p2, _ := parseDeck(2, deck2)
	want := -291
	got := recursive(p1, p2)
	if got != want {
		t.Errorf("recursive = %d, want %d", got, want)
	}
}
