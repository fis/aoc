package day10

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		line          string
		wantCorrupted bool
		wantScore     int
	}{
		{line: "[({(<(())[]>[[{[]{<()<>>", wantCorrupted: false, wantScore: 288957},
		{line: "[(()[<>])]({[<{<<[]>>(", wantCorrupted: false, wantScore: 5566},
		{line: "{([(<{}[<>[]}>{[]{[(<()>", wantCorrupted: true, wantScore: 1197},
		{line: "(((({<>}<{<{<>}{[]{[]{}", wantCorrupted: false, wantScore: 1480781},
		{line: "[[<[([]))<([[{}[[()]]]", wantCorrupted: true, wantScore: 3},
		{line: "[{[{({}]{}}([{[{{{}}([]", wantCorrupted: true, wantScore: 57},
		{line: "{<[[]]>}<{[{[{[]{()[[[]", wantCorrupted: false, wantScore: 995444},
		{line: "[<(<(<(<{}))><([]([]()", wantCorrupted: true, wantScore: 3},
		{line: "<{([([[(<>()){}]>(<<{{", wantCorrupted: true, wantScore: 25137},
		{line: "<{([{{}}[<[[[<>{}]]]>[]]", wantCorrupted: false, wantScore: 294},
	}
	for _, test := range tests {
		corrupted, score := check(test.line)
		if corrupted != test.wantCorrupted || score != test.wantScore {
			t.Errorf("check(%q) = (%t, %d), want (%t, %d)", test.line, corrupted, score, test.wantCorrupted, test.wantScore)
		}
	}
}
