// Package day10 solves AoC 2019 day 10.
package day10

import (
	"math"
	"sort"
	"strconv"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	level, err := util.ReadLevel(path, '.')
	if err != nil {
		return nil, err
	}

	at, vis := findBest(level)
	nth := findNth(at, 200, level)

	return []string{strconv.Itoa(vis), strconv.Itoa(nth.X*100 + nth.Y)}, nil
}

func findBest(level *util.Level) (util.P, int) {
	bestX, bestY, best := 0, 0, -1
	level.Range(func(fromX, fromY int, _ byte) {
		visible := findVisible(fromX, fromY, level)
		if len(visible) > best {
			bestX, bestY, best = fromX, fromY, len(visible)
		}
	})
	return util.P{bestX, bestY}, best
}

func findNth(from util.P, n int, level *util.Level) util.P {
	for {
		vis := findVisible(from.X, from.Y, level)
		if len(vis) == 0 {
			panic("out of asteroids :(")
		}
		if n <= len(vis) {
			sort.Slice(vis, func(i, j int) bool { return angle(from, vis[i]) < angle(from, vis[j]) })
			return vis[n-1]
		}
		for _, p := range vis {
			level.Set(p.X, p.Y, '.')
		}
		n -= len(vis)
	}
}

func findVisible(fromX, fromY int, level *util.Level) []util.P {
	var visible []util.P
	level.Range(func(toX, toY int, _ byte) {
		dx, dy := toX-fromX, toY-fromY
		if dx == 0 && dy == 0 {
			return
		}
		d := gcd(dx, dy)
		dx, dy = dx/d, dy/d
		tx, ty := fromX+dx, fromY+dy
		for (tx != toX || ty != toY) && level.At(tx, ty) != '#' {
			tx, ty = tx+dx, ty+dy
		}
		if tx == toX && ty == toY {
			visible = append(visible, util.P{toX, toY})
		}
	})
	return visible
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func angle(from, to util.P) float64 {
	dx, dy := float64(to.X-from.X), float64(to.Y-from.Y)
	return math.Mod(math.Atan2(dx, -dy)+2*math.Pi, 2*math.Pi)
}
