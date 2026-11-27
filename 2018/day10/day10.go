// Package day10 solves AoC 2018 day 10.
package day10

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 10, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	stars := parseInput(lines)

	t := findConjunction(stars) // not perfect, but close enough
	msg := drawAt(stars, t-1)
	return append(msg, strconv.Itoa(t-1)), nil
}

type star struct{ pos, vel util.P }

func parseInput(lines []string) (stars []star) {
	for _, line := range lines {
		var s star
		if _, err := fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &s.pos.X, &s.pos.Y, &s.vel.X, &s.vel.Y); err == nil {
			stars = append(stars, s)
		}
	}
	return stars
}

func findConjunction(stars []star) int {
	minDX, maxDX := stars[0].vel.X, stars[0].vel.X
	for _, s := range stars {
		if s.vel.X < minDX {
			minDX = s.vel.X
		}
		if s.vel.X > maxDX {
			maxDX = s.vel.X
		}
	}

	minCX, nMinCX, maxCX, nMaxCX := 0, 0, 0, 0
	for _, s := range stars {
		if s.vel.X == minDX {
			minCX, nMinCX = minCX+s.pos.X, nMinCX+1
		}
		if s.vel.X == maxDX {
			maxCX, nMaxCX = maxCX+s.pos.X, nMaxCX+1
		}
	}
	minCX, maxCX = minCX/nMinCX, maxCX/nMaxCX

	dist, speed := minCX-maxCX, maxDX-minDX
	return (dist + speed/2) / speed
}

func drawAt(stars []star, t int) (lines []string) {
	var points []util.P
	for _, s := range stars {
		points = append(points, s.pos.Add(s.vel.Scale(t)))
	}
	min, max := util.Bounds(points)
	w, h := max.X-min.X+1, max.Y-min.Y+1
	pixels := make([]byte, w*h)
	for i := range pixels {
		pixels[i] = ' '
	}
	for _, p := range points {
		x, y := p.X-min.X, p.Y-min.Y
		pixels[y*w+x] = '#'
	}
	for y := range h {
		lines = append(lines, string(pixels[y*w:(y+1)*w]))
	}
	return lines
}
