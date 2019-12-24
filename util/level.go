package util

import (
	"fmt"
	"io/ioutil"
)

// A Level models a two-dimensional map of ASCII character cells, similar to a roguelike level.
type Level struct {
	data     map[P]byte
	empty    byte
	min, max P
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
		min:   P{0, 0},
		max:   P{maxX, maxY},
	}
}

// ParseLevelString parses a string into a level. See ReadLevel.
func ParseLevelString(data string, empty byte) *Level {
	return ParseLevel([]byte(data), empty)
}

// At returns the byte at the given coordinates.
func (l *Level) At(x, y int) byte {
	if x < l.min.X || y < l.min.Y || x > l.max.X || y > l.max.Y {
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
		if x < l.min.X {
			l.min.X = x
		}
		if y < l.min.Y {
			l.min.Y = y
		}
		if x > l.max.X {
			l.max.X = x
		}
		if y > l.max.Y {
			l.max.Y = y
		}
	}
}

// Bounds returns the top-left and bottom-right corners of the level's bounding box. See InBounds
// for the definition.
func (l *Level) Bounds() (min P, max P) {
	min, max = l.min, l.max
	return
}

// InBounds returns true if the given coordinates are within the bounding box of the level. The
// bounds will grow to accommodate new non-empty characters, but will never shrink even if those
// characters are later overwritten to be empty.
func (l *Level) InBounds(x, y int) bool {
	return x >= l.min.X && y >= l.min.Y && x <= l.max.X && y <= l.max.Y
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
