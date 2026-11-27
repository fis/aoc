package day05

import "testing"

var (
	rules = parseRules(`
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13
`)
	updates = [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
)

func TestCheckUpdates(t *testing.T) {
	want1, want2 := 143, 123
	if got1, got2 := checkUpdates(updates, rules); got1 != want1 || got2 != want2 {
		t.Errorf("checkUpdates(ex) = (%d,%d), want (%d,%d)", got1, got2, want1, want2)
	}
}
