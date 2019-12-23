package util

import (
	"fmt"
	"io/ioutil"
)

// A Level models a two-dimensional map of ASCII character cells, similar to a roguelike level.
type Level struct {
	data       map[P]byte
	empty      byte
	maxX, maxY int
}

// ReadLevel reads the contents of a text file into a level. Character cells outside the contents of
// the file are considered to be the specified empty byte.
func ReadLevel(path string, empty byte) (*Level, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading level: %v", err)
	}
	return ParseLevel(data, empty), nil
}

// ParseLevel parses a byte array into a level. See ReadLevel.
func ParseLevel(data []byte, empty byte) *Level {
	level := make(map[P]byte)
	x, y, maxX, maxY := 0, 0, 0, 0
	for _, b := range data {
		if b == '\n' {
			x = 0
			y++
			continue
		}
		if b != empty {
			level[P{x, y}] = b
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		x++
	}
	return &Level{
		data:  level,
		empty: empty,
		maxX:  maxX,
		maxY:  maxY,
	}
}

// ParseLevelString parses a string into a level. See ReadLevel.
func ParseLevelString(data string, empty byte) *Level {
	return ParseLevel([]byte(data), empty)
}

// At returns the byte at the given coordinates.
func (l *Level) At(x, y int) byte {
	if x < 0 || y < 0 || x > l.maxX || y > l.maxY {
		return l.empty
	}
	if b, ok := l.data[P{x, y}]; ok {
		return b
	}
	return l.empty
}

// Set sets the byte at the given coordinates.
func (l *Level) Set(x, y int, b byte) {
	if b == l.empty {
		delete(l.data, P{x, y})
	} else {
		l.data[P{x, y}] = b
	}
}

// InBounds returns true if the given coordinates are within the bounding box of the source file of
// the level.
func (l *Level) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x <= l.maxX && y <= l.maxY
}

// Range calls the callback function for all non-empty cells in the level.
func (l *Level) Range(cb func(x, y int, b byte)) {
	for p, b := range l.data {
		cb(p.X, p.Y, b)
	}
}

// Find locates the coordinates of a byte, which must be unique on the level.
func (l *Level) Find(key byte) (x, y int, found bool) {
	for p, b := range l.data {
		if b == key {
			if found {
				return 0, 0, false
			}
			x, y, found = p.X, p.Y, true
		}
	}
	return
}
