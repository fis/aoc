package day13

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex1 = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`

var ex2 = `
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`

func TestSimulate(t *testing.T) {
	tests := []struct {
		name         string
		level        string
		keepGoing    bool
		wantX, wantY int
	}{
		{name: "ex1", level: ex1, keepGoing: false, wantX: 7, wantY: 3},
		{name: "ex2", level: ex2, keepGoing: true, wantX: 6, wantY: 4},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), ' ')
		gotX, gotY := simulate(level, test.keepGoing)
		if gotX != test.wantX || gotY != test.wantY {
			t.Errorf("simulate(%s, %v) = (%d,%d), want (%d,%d)", test.name, test.keepGoing, gotX, gotY, test.wantX, test.wantY)
		}
	}
}
