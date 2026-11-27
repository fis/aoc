package day05

import (
	"testing"
)

func TestSearch(t *testing.T) {
	want1, want2 := "18f47a30", "05ace8e3"
	got1, got2 := search("abc", 8)
	if got1 != want1 || got2 != want2 {
		t.Errorf("search(abc, 8) = (%s, %s), want (%s, %s)", got1, got2, want1, want2)
	}
}
