// Package day15 solves AoC 2024 day 15.
package day15

import (
	"bytes"
	"io"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 15, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	level, moves := parseInput(r)
	wideLevel := embiggen(level)
	part1 := sumBoxes(level, moves)
	part2 := sumWideBoxes(wideLevel, moves)
	return glue.Ints(part1, part2), nil
}

func sumBoxes(level *util.FixedLevel, moves []byte) (sum int) {
	applyMoves(level, moves)
	for y := 1; y < level.H-1; y++ {
		for x := 1; x < level.W-1; x++ {
			if level.At(x, y) == 'O' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func applyMoves(level *util.FixedLevel, moves []byte) {
	x, y, _ := level.Find('@')
	for _, move := range moves {
		dx, dy, ok := decodeMove(move)
		if !ok {
			continue
		}
		nx, ny, boxStack := x+dx, y+dy, 0
		for level.At(nx, ny) == 'O' {
			nx, ny, boxStack = nx+dx, ny+dy, boxStack+1
		}
		if level.At(nx, ny) != '.' {
			continue // blocked
		}
		level.Set(x, y, '.')
		level.Set(x+dx, y+dy, '@')
		if boxStack > 0 {
			level.Set(nx, ny, 'O')
		}
		x, y = x+dx, y+dy
	}
}

func sumWideBoxes(level *util.FixedLevel, moves []byte) (sum int) {
	applyWideMoves(level, moves)
	for y := 1; y < level.H-1; y++ {
		for x := 1; x < level.W-1; x++ {
			if level.At(x, y) == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func applyWideMoves(level *util.FixedLevel, moves []byte) {
	x, y, _ := level.Find('@')
	bc := newCollector(level)
	for _, move := range moves {
		dx, dy, ok := decodeMove(move)
		if !ok {
			continue
		}
		// easy cases first
		nx, ny := x+dx, y+dy
		switch level.At(nx, ny) {
		case '#':
			continue
		case '.':
			level.Set(x, y, '.')
			level.Set(nx, ny, '@')
			x, y = nx, ny
			continue
		}
		// horizontal moves are still just about pushing a stack
		if dy == 0 {
			boxStack := 0
			if dx == -1 {
				for level.At(nx, ny) == ']' {
					nx, boxStack = nx-2, boxStack+1
				}
			} else {
				for level.At(nx, ny) == '[' {
					nx, boxStack = nx+2, boxStack+1
				}
			}
			if level.At(nx, ny) != '.' {
				continue // still blocked
			}
			level.Set(x, y, '.')
			level.Set(x+dx, y, '@')
			var bx1, bx2 int
			if dx == -1 {
				bx1, bx2 = nx, x-2
			} else {
				bx1, bx2 = x+2, nx
			}
			for bx, bc := bx1, byte('['); bx <= bx2; bx, bc = bx+1, bc^('['^']') {
				level.Set(bx, y, bc)
			}
			x += dx
			continue
		}
		// vertical moves are the tricky case
		if level.At(nx, ny) == ']' {
			nx--
		}
		bc.reset()
		if bc.collect(nx, ny, dy) {
			for _, p := range bc.coords {
				level.Set(p.X, p.Y+dy, '[')
				level.Set(p.X+1, p.Y+dy, ']')
				level.Set(p.X, p.Y, '.')
				level.Set(p.X+1, p.Y, '.')
			}
			level.Set(x, y, '.')
			level.Set(x, y+dy, '@')
			y += dy
		}
	}
}

type boxCollector struct {
	level     *util.FixedLevel
	collected util.FixedBitmap2D
	coords    []util.P
}

func newCollector(level *util.FixedLevel) *boxCollector {
	return &boxCollector{
		level:     level,
		collected: util.MakeFixedBitmap2D(level.W, level.H),
		coords:    make([]util.P, 0, 64),
	}
}

func (bc *boxCollector) reset() {
	for _, p := range bc.coords {
		bc.collected.ClearWord(p.X, p.Y)
	}
	bc.coords = bc.coords[:0]
}

func (bc *boxCollector) collect(x, y, dy int) bool {
	t1, t2 := bc.level.At(x, y+dy), bc.level.At(x+1, y+dy)
	switch {
	case t1 == '.' && t2 == '.': // free space: move this box
		bc.record(x, y)
		return true
	case t1 == '[' && t2 == ']': // aligned box above/below; move this and that
		if bc.collect(x, y+dy, dy) {
			bc.record(x, y)
			return true
		}
	case t1 == ']' && t2 == '.': // single misaligned box to the left
		if bc.collected.Get(x-1, y+dy) || bc.collect(x-1, y+dy, dy) {
			bc.record(x, y)
			return true
		}
	case t1 == '.' && t2 == '[': // single misaligned box to the right
		if bc.collected.Get(x+1, y+dy) || bc.collect(x+1, y+dy, dy) {
			bc.record(x, y)
			return true
		}
	case t1 == ']' && t2 == '[': // branching case: move the two boxes above/below
		if (bc.collected.Get(x-1, y+dy) || bc.collect(x-1, y+dy, dy)) && (bc.collected.Get(x+1, y+dy) || bc.collect(x+1, y+dy, dy)) {
			bc.record(x, y)
			return true
		}
	}
	return false
}

func (bc *boxCollector) record(x, y int) {
	bc.collected.Set(x, y)
	bc.coords = append(bc.coords, util.P{x, y})
}

func decodeMove(move byte) (dx, dy int, ok bool) {
	switch move {
	case '<':
		dx, dy = -1, 0
	case '>':
		dx, dy = 1, 0
	case '^':
		dx, dy = 0, -1
	case 'v':
		dx, dy = 0, 1
	default:
		return 0, 0, false
	}
	return dx, dy, true
}

func parseInput(r io.Reader) (level *util.FixedLevel, moves []byte) {
	data, _ := io.ReadAll(r)
	levelData, moves, _ := bytes.Cut(data, []byte("\n\n"))
	levelData = append(levelData, '\n')
	level = util.ParseFixedLevel(levelData)
	return level, moves
}

func embiggen(level *util.FixedLevel) *util.FixedLevel {
	newLevel := &util.FixedLevel{W: level.W << 1, H: level.H, Data: make([]byte, len(level.Data)<<1)}
	for y := range level.H {
		for x := range level.W {
			nx := x << 1
			var t1, t2 byte
			switch level.At(x, y) {
			case '#':
				t1, t2 = '#', '#'
			case 'O':
				t1, t2 = '[', ']'
			case '.':
				t1, t2 = '.', '.'
			case '@':
				t1, t2 = '@', '.'
			}
			newLevel.Set(nx, y, t1)
			newLevel.Set(nx+1, y, t2)
		}
	}
	return newLevel
}
