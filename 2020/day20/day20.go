// Copyright 2020 Google LLC
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

// Package day20 solves AoC 2020 day 20.
package day20

import (
	"fmt"
	"math"

	"github.com/fis/aoc-go/glue"
	"github.com/fis/aoc-go/util"
)

func init() {
	glue.RegisterSolver(2020, 20, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]int, error) {
	set, err := parseSet(chunks)
	if err != nil {
		return nil, err
	}
	img, size, part1, err := set.reassemble()
	if err != nil {
		return nil, err
	}
	part2 := roughness(img, size)
	return []int{part1, part2}, nil
}

/*
Orientation and edge numbering scheme:

        0    1    2    3      4    5    6    7

  123
0 456  123  741  987  369    321  147  789  963
  789   a    d    c    b      e    f    g    h

  369
1 258  369  123  741  987    963  321  147  789
  147   b    a    d    c      h    e    f    g

  987
2 654  987  369  123  741    789  963  321  147
  321   c    b    a    d      g    h    e    f

  741
3 852  741  987  369  123    147  789  963  321
  963   d    c    b    a      f    g    h    e


  321
4 654  321  963  789  147    123  369  987  741
  987   e    h    g    f      a    b    c    d

  147
5 258  147  321  963  789    741  123  369  987
  369   f    e    h    g      d    a    b    c

  789
6 456  789  147  321  963    987  741  123  369
  123   g    f    e    h      c    d    a    b

  963
7 852  963  789  147  321    369  987  741  123
	741   h    g    f    e      b    c    d    a
*/

type orientation int

const (
	oriR0 orientation = iota
	oriR90
	oriR180
	oriR270
	oriFR0
	oriFR90
	oriFR180
	oriFR270
)

func (o orientation) basis(size int) (zero, dx, dy util.P) {
	switch o {
	case oriR0:
		return util.P{0, 0}, util.P{1, 0}, util.P{0, 1}
	case oriR90:
		return util.P{0, size - 1}, util.P{0, -1}, util.P{1, 0}
	case oriR180:
		return util.P{size - 1, size - 1}, util.P{-1, 0}, util.P{0, -1}
	case oriR270:
		return util.P{size - 1, 0}, util.P{0, 1}, util.P{-1, 0}
	case oriFR0:
		return util.P{size - 1, 0}, util.P{-1, 0}, util.P{0, 1}
	case oriFR90:
		return util.P{0, 0}, util.P{0, 1}, util.P{1, 0}
	case oriFR180:
		return util.P{0, size - 1}, util.P{1, 0}, util.P{0, -1}
	case oriFR270:
		return util.P{size - 1, size - 1}, util.P{0, -1}, util.P{-1, 0}
	}
	return util.P{}, util.P{}, util.P{}
}

type edgeType int

const (
	edgeTop edgeType = iota
	edgeLeft
	edgeBottom
	edgeRight
	edgeFlipTop
	edgeFlipLeft
	edgeFlipBottom
	edgeFlipRight
)

func (e edgeType) reverse(o orientation) edgeType {
	es, ef, os, of := int(e%4), int(e/4), int(o%4), int(o/4)
	switch 2*ef + of {
	case 0b00:
		return edgeType((4 - os + es) % 4)
	case 0b01:
		return edgeType(4 + (4-es+os)%4)
	case 0b10:
		return edgeType(4 + (4-os+es)%4)
	case 0b11:
		return edgeType((4 - es + os) % 4)
	}
	return edgeTop
}

func (e edgeType) orientTo(t edgeType) orientation {
	es, ef, ts, tf := int(e%4), int(e/4), int(t%4), int(t/4)
	switch {
	case ef == tf:
		return orientation((4 - es + ts) % 4)
	default:
		return orientation(4 + (es+ts)%4)
	}
}

type imageTile struct {
	id        int
	size      int
	data      []byte
	edges     [8]uint
	links     [8]edgeLink
	linkCount int
}

type edgeLink struct {
	tile *imageTile
	edge edgeType
}

func parseTile(lines []string) (*imageTile, error) {
	tile := &imageTile{}

	if len(lines) < 2 || len(lines) != len(lines[1])+1 {
		return nil, fmt.Errorf("tile not square: %dx%d", len(lines[1]), len(lines)-1)
	}
	if _, err := fmt.Sscanf(lines[0], "Tile %d:", &tile.id); err != nil {
		return nil, fmt.Errorf("malformatted tile header: %s: %w", lines[0], err)
	}
	lines, tile.size = lines[1:], len(lines[1])

	tile.data = make([]byte, tile.size*tile.size)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				tile.data[y*tile.size+x] = 1
			}
		}
	}

	for typ := edgeTop; typ <= edgeFlipRight; typ++ {
		edge := uint(0)
		at, dx, _ := edgeTop.orientTo(typ).basis(tile.size)
		for i := 0; i < tile.size; i, at = i+1, at.Add(dx) {
			edge = (edge << 1) | uint(tile.data[at.Y*tile.size+at.X])
		}
		tile.edges[typ] = edge
	}

	return tile, nil
}

type imageSet struct {
	tiles []*imageTile
}

func parseSet(chunks []string) (*imageSet, error) {
	set := &imageSet{}
	for _, chunk := range chunks {
		t, err := parseTile(util.Lines(chunk))
		if err != nil {
			return nil, err
		}
		if err := set.addTile(t); err != nil {
			return nil, err
		}
	}
	return set, nil
}

func (s *imageSet) addTile(t *imageTile) error {
	if len(s.tiles) > 0 && len(s.tiles[0].data) != len(t.data) {
		return fmt.Errorf("mismatching image sizes: %d, %d", len(s.tiles[0].data), len(t.data))
	}
	s.tiles = append(s.tiles, t)
	return nil
}

func (s *imageSet) reassemble() (merged []byte, msize, checksum int, err error) {
	size := int(math.Sqrt(float64(len(s.tiles))))
	if len(s.tiles) != size*size {
		return nil, 0, 0, fmt.Errorf("not a square: %d", len(s.tiles))
	}

	matches := map[uint][]edgeLink{}
	for _, t := range s.tiles {
		for typ, val := range t.edges {
			matches[val] = append(matches[val], edgeLink{t, edgeType(typ)})
		}
	}

	unmatched, matched, ambiguous := 0, 0, 0
	for _, edges := range matches {
		switch len(edges) {
		case 0:
			panic("impossible: matching 0 edges")
		case 1:
			unmatched++
		case 2:
			matched++
			l, r := edges[0], edges[1]
			l.tile.links[l.edge] = edgeLink{tile: r.tile, edge: r.edge}
			r.tile.links[r.edge] = edgeLink{tile: l.tile, edge: l.edge}
			l.tile.linkCount++
			r.tile.linkCount++
		default:
			ambiguous++
		}
	}
	if unmatched != 2*4*size || matched != 4*size*(size-1) || ambiguous > 0 {
		return nil, 0, 0, fmt.Errorf("unexpected match counts: %d, %d, %d", unmatched, matched, ambiguous)
	}

	cornerTiles, edgeTiles, innerTiles := 0, 0, 0
	checksum = 1
	for _, t := range s.tiles {
		switch t.linkCount {
		case 4:
			cornerTiles++
			checksum *= t.id
		case 6:
			edgeTiles++
		case 8:
			innerTiles++
		default:
			return nil, 0, 0, fmt.Errorf("tile %d has %d links", t.id, t.linkCount)
		}
	}
	if cornerTiles != 4 || edgeTiles != 4*(size-2) || innerTiles != (size-2)*(size-2) {
		return nil, 0, 0, fmt.Errorf("unexpected link counts: %d, %d, %d", cornerTiles, edgeTiles, innerTiles)
	}

	type orientedTile struct {
		tile *imageTile
		ori  orientation
	}
	arrangement := make([]orientedTile, size*size)
	for y := 0; y < size; y++ {
		if y == 0 {
			var (
				corner *imageTile
				rot    orientation
			)
			for _, t := range s.tiles {
				if t.linkCount == 4 {
					corner = t
					break
				}
			}
			for rot = oriR0; rot <= oriR270; rot++ {
				if corner.links[edgeTop.reverse(rot)].tile == nil && corner.links[edgeLeft.reverse(rot)].tile == nil {
					break
				}
			}
			if rot > oriR270 {
				return nil, 0, 0, fmt.Errorf("tile %d failed to work as top-left corner", corner.id)
			}
			arrangement[0] = orientedTile{tile: corner, ori: rot}
		} else {
			top := arrangement[(y-1)*size]
			link := top.tile.links[edgeBottom.reverse(top.ori)]
			arrangement[y*size] = orientedTile{tile: link.tile, ori: link.edge.orientTo(edgeFlipTop)}
		}

		for x := 1; x < size; x++ {
			left := arrangement[y*size+x-1]
			link := left.tile.links[edgeRight.reverse(left.ori)]
			arrangement[y*size+x] = orientedTile{tile: link.tile, ori: link.edge.orientTo(edgeFlipLeft)}
		}
	}

	tsize := s.tiles[0].size - 2
	msize = size * tsize
	merged = make([]byte, msize*msize)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			t := arrangement[y*size+x]
			mx, my := x*tsize, y*tsize
			tz, tdx, tdy := t.ori.basis(tsize)
			for ty := 0; ty < tsize; ty++ {
				for tx := 0; tx < tsize; tx++ {
					tp := tz.Add(tdy.Scale(ty)).Add(tdx.Scale(tx))
					merged[(my+tp.Y)*msize+mx+tp.X] = t.tile.data[(ty+1)*t.tile.size+tx+1]
				}
			}
		}
	}

	return merged, msize, checksum, nil
}

func roughness(img []byte, size int) (sum int) {
	// Reference image, 20x3
	// ..................#.
	// #....##....##....###
	// .#..#..#..#..#..#...
	refW, refH := 20, 3
	ref := []util.P{
		{0, 1}, {1, 2}, {4, 2}, {5, 1}, {6, 1},
		{7, 2}, {10, 2}, {11, 1}, {12, 1}, {13, 2},
		{16, 2}, {17, 1}, {18, 1}, {19, 1}, {18, 0},
	}
	for ori := oriR0; ori <= oriFR270; ori++ {
		var monsters []util.P
		z, dx, dy := ori.basis(size)
		for by := 0; by+refH <= size; by++ {
		next:
			for bx := 0; bx+refW <= size; bx++ {
				for _, r := range ref {
					p := z.Add(dy.Scale(by + r.Y)).Add(dx.Scale(bx + r.X))
					if img[p.Y*size+p.X] == 0 {
						continue next
					}
				}
				monsters = append(monsters, util.P{bx, by})
			}
		}
		if len(monsters) > 0 {
			for _, b := range monsters {
				for _, r := range ref {
					p := z.Add(dy.Scale(b.Y + r.Y)).Add(dx.Scale(b.X + r.X))
					img[p.Y*size+p.X] = 0
				}
			}
			for _, v := range img {
				sum += int(v)
			}
			return sum
		}
	}
	return -1
}
