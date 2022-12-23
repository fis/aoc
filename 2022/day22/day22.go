// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package day22 solves AoC 2022 day 22.
package day22

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2022, 22, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	lvl, path, err := parseInput(lines)
	if err != nil {
		return nil, err
	}
	p1 := decode(lvl, path)
	c, err := fold(lvl, 50)
	if err != nil {
		return nil, err
	}
	p2 := decodeCube(c, path)
	return glue.Ints(p1, p2), nil
}

type level struct {
	data         []string
	w, h         int
	wrapX, wrapY [][2]int
}

type move struct {
	steps int
	turn  int
}

func decode(lvl *level, path []move) (pass int) {
	x, y, dx, dy := lvl.wrapX[0][0]+1, 0, 1, 0
	for _, m := range path {
		for i := 0; i < m.steps; i++ {
			nx := wrap(x+dx, lvl.wrapX[y][0], lvl.wrapX[y][1])
			ny := wrap(y+dy, lvl.wrapY[x][0], lvl.wrapY[x][1])
			if lvl.data[ny][nx] != '.' {
				break
			}
			x, y = nx, ny
		}
		if t := m.turn; t != 0 {
			dx, dy = -t*dy, t*dx
		}
	}
	dir := (1 - dx) + (dy & 2)
	return 1000*(y+1) + 4*(x+1) + dir
}

func wrap(i, lo, hi int) int {
	switch i {
	case lo:
		return hi - 1
	case hi:
		return lo + 1
	default:
		return i
	}
}

/*
# Cubical Conventions

Faces are numbered as they would be in this canonical net:

    4
    0123
    5

Directions are numbered following the scheme given in the question:

    0 > right
    1 v down
    2 < left
    3 ^ up

Face orientation is recorded in units of clockwise 90-degree rotations.
This means they can be added to or subtracted from directions (modulo 4).
The rotation indicates how the face data must be transformed in order to
match canonical orientations (those in the sample net).
*/

type cube struct {
	size  int
	faces [6]struct {
		data []string
		rot  rotation
		x, y int
	}
	edges [6][4]struct {
		face int
	}
}

func decodeCube(c *cube, path []move) (pass int) {
	f, x, y, dir, dx, dy := 0, 0, 0, dirRight, 1, 0
	for _, m := range path {
		for i := 0; i < m.steps; i++ {
			nf, nx, ny, ndir, ndx, ndy := f, x+dx, y+dy, dir, dx, dy
			if nx < 0 || nx >= c.size || ny < 0 || ny >= c.size {
				n := cubeLinks[f][c.faces[f].rot.rotD(dir)]
				nf = n.face
				ndir = c.faces[nf].rot.invert().rotD(n.dir)
				ndx, ndy = ndir.vec()
				cx, cy := c.faces[f].rot.rotXY(x, y, c.size-1)
				nx = n.mxx*cx + n.mxy*cy + n.mxs*(c.size-1)
				ny = n.myx*cx + n.myy*cy + n.mys*(c.size-1)
				nx, ny = c.faces[nf].rot.invert().rotXY(nx, ny, c.size-1)
			}
			if c.faces[nf].data[ny][nx] != '.' {
				break
			}
			f, x, y, dir, dx, dy = nf, nx, ny, ndir, ndx, ndy
		}
		if t := m.turn; t != 0 {
			dir, dx, dy = (dir+direction(t))&3, -t*dy, t*dx
		}
	}
	return 1000*(c.faces[f].y+y+1) + 4*(c.faces[f].x+x+1) + int(dir)
}

// cubeLinks[face][direction]
var cubeLinks = [6][4]struct {
	face          int
	dir           direction
	mxx, mxy, mxs int
	myx, myy, mys int
}{
	{{1, 0, 0, 0, 0, 0, 1, 0}, {5, 1, 1, 0, 0, 0, 0, 0}, {3, 2, 0, 0, 1, 0, 1, 0}, {4, 3, 1, 0, 0, 0, 0, 1}},   // face 0
	{{2, 0, 0, 0, 0, 0, 1, 0}, {5, 2, 0, 0, 1, 1, 0, 0}, {0, 2, 0, 0, 1, 0, 1, 0}, {4, 2, 0, 0, 1, -1, 0, 1}},  // face 1
	{{3, 0, 0, 0, 0, 0, 1, 0}, {5, 3, -1, 0, 1, 0, 0, 1}, {1, 2, 0, 0, 1, 0, 1, 0}, {4, 1, -1, 0, 1, 0, 0, 0}}, // face 2
	{{0, 0, 0, 0, 0, 0, 1, 0}, {5, 0, 0, 0, 0, -1, 0, 1}, {2, 2, 0, 0, 1, 0, 1, 0}, {4, 0, 0, 0, 0, 1, 0, 0}},  // face 3
	{{1, 1, 0, -1, 1, 0, 0, 0}, {0, 1, 1, 0, 0, 0, 0, 0}, {3, 1, 0, 1, 0, 0, 0, 0}, {2, 1, -1, 0, 1, 0, 0, 0}}, // face 4
	{{1, 3, 0, 1, 0, 0, 0, 1}, {2, 3, -1, 0, 1, 0, 0, 1}, {3, 3, 0, -1, 1, 0, 0, 1}, {0, 3, 1, 0, 0, 0, 0, 1}}, // face 5
}

func fold(lvl *level, size int) (c *cube, err error) {
	if lvl.w%size != 0 || lvl.h%size != 0 {
		return nil, fmt.Errorf("level size (%d x %d) not a multiple of edge length (%d)", lvl.w, lvl.h, size)
	}
	c = &cube{size: size}

	fw, fh := lvl.w/size, lvl.h/size
	fmap, fcount := util.MakeFixedBitmap2D(fw, fh), 0
	for fy := 0; fy < fh; fy++ {
		for fx := 0; fx < fw; fx++ {
			x, y := fx*size, fy*size
			if x < len(lvl.data[y]) && lvl.data[y][x] != ' ' {
				fmap.Set(fx, fy)
				fcount++
			}
		}
	}
	if fcount != 6 {
		return nil, fmt.Errorf("expected 6 faces, got %d (total in file)", fcount)
	}

	fcount = 0
	for fx := 0; fx < fw; fx++ {
		if lvl.data[0][fx*size] != ' ' {
			fcount = findFaces(c, lvl, fw, fh, fx, 0, 0, 0, fmap)
			break
		}
	}
	if fcount != 6 {
		return nil, fmt.Errorf("expected 6 faces, got %d (connected)", fcount)
	}

	return c, nil
}

func findFaces(c *cube, lvl *level, fw, fh, fx, fy, face int, rot rotation, fmap util.FixedBitmap2D) (foundFaces int) {
	foundFaces = 1
	fmap.Clear(fx, fy)

	c.faces[face].data = make([]string, c.size)
	for i := 0; i < c.size; i++ {
		c.faces[face].data[i] = lvl.data[fy*c.size+i][fx*c.size : (fx+1)*c.size]
	}
	c.faces[face].rot = rot
	c.faces[face].x, c.faces[face].y = fx*c.size, fy*c.size

	if fx+1 < fw && fmap.Get(fx+1, fy) {
		n := cubeLinks[face][rot.rotD(0)]
		foundFaces += findFaces(c, lvl, fw, fh, fx+1, fy, n.face, angleOf(0, n.dir), fmap)
	}
	if fy+1 < fh && fmap.Get(fx, fy+1) {
		n := cubeLinks[face][rot.rotD(1)]
		foundFaces += findFaces(c, lvl, fw, fh, fx, fy+1, n.face, angleOf(1, n.dir), fmap)
	}
	if fx-1 >= 0 && fmap.Get(fx-1, fy) {
		n := cubeLinks[face][rot.rotD(2)]
		foundFaces += findFaces(c, lvl, fw, fh, fx-1, fy, n.face, angleOf(2, n.dir), fmap)
	}
	if fy-1 >= 0 && fmap.Get(fx, fy-1) {
		n := cubeLinks[face][rot.rotD(3)]
		foundFaces += findFaces(c, lvl, fw, fh, fx, fy-1, n.face, angleOf(3, n.dir), fmap)
	}

	return foundFaces
}

type direction int

const (
	dirRight direction = iota
	dirDown
	dirLeft
	dirUp
)

var dirVecs = [4]struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func (dir direction) vec() (dx, dy int) {
	return dirVecs[dir].x, dirVecs[dir].y
}

type rotation int

const (
	rot0 rotation = iota
	rot90
	rot180
	rot270
)

var rotOrigins = [4]struct{ mx, my int }{
	rot0: {0, 0}, rot90: {1, 0}, rot180: {1, 1}, rot270: {0, 1},
}

var rotVecs = [4]struct{ mxx, mxy, mxs, myx, myy, mys int }{
	rot0:   {1, 0, 0, 0, 1, 0},
	rot90:  {0, -1, 1, 1, 0, 0},
	rot180: {-1, 0, 1, 0, -1, 1},
	rot270: {0, 1, 0, -1, 0, 1},
}

func (rot rotation) invert() rotation {
	return (-rot) & 3
}

func (rot rotation) origin(size int) (x, y int) {
	return (size - 1) * rotOrigins[rot].mx, (size - 1) * rotOrigins[rot].my
}

func (rot rotation) rotD(dir direction) direction {
	return (dir + direction(rot)) & 3
}

func (rot rotation) rotXY(x, y, size int) (nx, ny int) {
	v := rotVecs[rot]
	return v.mxx*x + v.mxy*y + v.mxs*size, v.myx*x + v.myy*y + v.mys*size
}

func angleOf(fileDir, cubeDir direction) rotation {
	return rotation((cubeDir - fileDir) & 3)
}

func parseInput(lines []string) (lvl *level, path []move, err error) {
	if len(lines) < 3 || lines[len(lines)-2] != "" {
		return nil, nil, fmt.Errorf("expected input: level, blank line, path")
	}

	w, h := 0, len(lines)-2
	wrapX := make([][2]int, h)
	for y := 0; y < h; y++ {
		line := lines[y]
		if len(line) > w {
			w = len(line)
		}
		x0, x1 := -1, len(line)
		for x0+1 < len(line) && line[x0+1] == ' ' {
			x0++
		}
		for x1-1 >= 0 && line[x1-1] == ' ' {
			x1--
		}
		if x0+1 > x1-1 {
			return nil, nil, fmt.Errorf("no content on line %d", y)
		}
		wrapX[y] = [2]int{x0, x1}
	}
	wrapY := make([][2]int, w)
	for x := 0; x < w; x++ {
		y0, y1 := -1, h
		for y0+1 < h && (x >= len(lines[y0+1]) || lines[y0+1][x] == ' ') {
			y0++
		}
		for y1-1 >= 0 && (x >= len(lines[y1-1]) || lines[y1-1][x] == ' ') {
			y1--
		}
		if y0+1 > y1-1 {
			return nil, nil, fmt.Errorf("no content on column %d", x)
		}
		wrapY[x] = [2]int{y0, y1}
	}
	lvl = &level{data: lines[:h], w: w, h: h, wrapX: wrapX, wrapY: wrapY}

	pathSpec := lines[len(lines)-1]
	for len(pathSpec) > 0 {
		steps, ok, tail := util.NextInt(pathSpec)
		if !ok {
			return nil, nil, fmt.Errorf("expected an integer")
		}
		m := move{steps: steps}
		pathSpec = tail
		if len(pathSpec) > 0 {
			turn := pathSpec[0]
			switch turn {
			case 'L':
				m.turn = -1
			case 'R':
				m.turn = +1
			default:
				return nil, nil, fmt.Errorf("expected L or R")
			}
			pathSpec = pathSpec[1:]
		}
		path = append(path, m)
	}

	return lvl, path, nil
}
