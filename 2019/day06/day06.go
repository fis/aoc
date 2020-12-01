// Package day06 solves AoC 2019 day 6.
package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	orbits := parseOrbits(lines)

	p1 := orbits.countOrbits()
	p2 := orbits.transfers("YOU", "SAN")

	return []string{strconv.Itoa(p1), strconv.Itoa(p2)}, nil
}

func parseOrbits(lines []string) *orbitMap {
	om := &orbitMap{}
	for _, line := range lines {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid orbit: %q", line))
		}
		om.add(parts[0], parts[1])
	}
	return om
}

type orbitMap struct {
	orbits map[string]*orbit
}

type orbit struct {
	name   string
	parent *orbit
	child  []*orbit
}

func (om *orbitMap) get(name string) *orbit {
	if om.orbits == nil {
		om.orbits = make(map[string]*orbit)
	}
	o, ok := om.orbits[name]
	if !ok {
		o = &orbit{name: name}
		om.orbits[name] = o
	}
	return o
}

func (om *orbitMap) add(parent, child string) {
	p, c := om.get(parent), om.get(child)
	if c.parent != nil {
		panic(fmt.Sprintf("conflicting orbits: %s)%s, %s)%s", c.parent.name, child, parent, child))
	}
	p.child = append(p.child, c)
	c.parent = p
}

func (om *orbitMap) countOrbits() int {
	return om.get("COM").countOrbits(1)
}

func (om *orbitMap) transfers(from, to string) int {
	fromP, toP := om.get(from).path(), om.get(to).path()
	for len(fromP) > 0 && len(toP) > 0 && fromP[0] == toP[0] {
		fromP, toP = fromP[1:], toP[1:]
	}
	return len(fromP) + len(toP) - 2
}

func (o *orbit) countOrbits(depth int) int {
	s := 0
	for _, c := range o.child {
		s += depth
		s += c.countOrbits(depth + 1)
	}
	return s
}

func (o *orbit) path() []string {
	n := 0
	for p := o; p != nil; p = p.parent {
		n++
	}
	path := make([]string, n)
	for p, i := o, n-1; p != nil; p, i = p.parent, i-1 {
		path[i] = p.name
	}
	return path
}
