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

// Package day18 solves AoC 2022 day 18.
package day18

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2022, 18, glue.LineSolver(glue.WithParser(parseCube, solve)))
}

func parseCube(line string) (p3, error) {
	x, okx, tail := util.NextInt(line)
	tail, okc1 := util.CheckPrefix(tail, ",")
	y, oky, tail := util.NextInt(tail)
	tail, okc2 := util.CheckPrefix(tail, ",")
	z, okz, _ := util.NextInt(tail)
	if okx && okc1 && oky && okc2 && okz {
		return p3{x + 1, y + 1, z + 1}, nil
	}
	return p3{}, fmt.Errorf("bad point: expecting \"X,Y,Z\", got %q", line)
}

func solve(cubes []p3) ([]string, error) {
	p1, p2 := surfaceAreas(cubes)
	return glue.Ints(p1, p2), nil
}

func surfaceAreas(cubes []p3) (sa1, sa2 int) {
	cubeMap := makeMap(cubes)
	findOutside(cubes, cubeMap)

	for _, c := range cubes {
		for _, n := range c.neigh() {
			m := cubeMap[n.z][n.y][n.x]
			if m != cubeStuff {
				sa1++
			}
			if m == outsideAir {
				sa2++
			}
		}
	}

	return sa1, sa2
}

func makeMap(cubes []p3) (cubeMap [][][]material) {
	max := max3(cubes)
	maxX, maxY, maxZ := max.x+2, max.y+2, max.z+2

	data := make([]material, maxX*maxY*maxZ)
	rows := make([][]material, maxY*maxZ)
	for zy := range rows {
		rows[zy] = data[zy*maxX : (zy+1)*maxX]
	}
	cubeMap = make([][][]material, maxZ)
	for z := range cubeMap {
		cubeMap[z] = rows[z*maxY : (z+1)*maxY]
	}

	for _, c := range cubes {
		cubeMap[c.z][c.y][c.x] = cubeStuff
	}

	return cubeMap
}

func findOutside(cubes []p3, cubeMap [][][]material) {
	maxX, maxY, maxZ := len(cubeMap[0][0]), len(cubeMap[0]), len(cubeMap)
	cubeMap[0][0][0] = outsideAir
	q := []p3{{0, 0, 0}}
	for len(q) > 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]
		for _, n := range p.neigh() {
			if n.x < 0 || n.x >= maxX || n.y < 0 || n.y >= maxY || n.z < 0 || n.z >= maxZ {
				continue
			}
			if cubeMap[n.z][n.y][n.x] != insideAir {
				continue
			}
			cubeMap[n.z][n.y][n.x] = outsideAir
			q = append(q, n)
		}
	}
}

type material byte

const (
	insideAir material = iota
	cubeStuff
	outsideAir
)

type p3 struct {
	x, y, z int
}

func (p p3) neigh() [6]p3 {
	return [6]p3{
		{p.x - 1, p.y, p.z},
		{p.x + 1, p.y, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y, p.z - 1},
		{p.x, p.y, p.z + 1},
	}
}

func max3(cubes []p3) (maxP p3) {
	maxP = cubes[0]
	for _, c := range cubes[1:] {
		maxP.x = max(maxP.x, c.x)
		maxP.y = max(maxP.y, c.y)
		maxP.z = max(maxP.z, c.z)
	}
	return maxP
}
