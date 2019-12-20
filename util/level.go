package util

import (
	"fmt"
	"io/ioutil"
)

// A Level models a two-dimensional map of ASCII character cells, similar to a roguelike level.
type Level struct {
	data       map[[2]int]byte
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
	level := make(map[[2]int]byte)
	x, y, maxX, maxY := 0, 0, 0, 0
	for _, b := range data {
		if b == '\n' {
			x = 0
			y++
			continue
		}
		if b != empty {
			level[[2]int{x, y}] = b
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
	if b, ok := l.data[[2]int{x, y}]; ok {
		return b
	}
	return l.empty
}

// InBounds returns true if the given coordinates are within the bounding box of the source file of
// the level.
func (l *Level) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x <= l.maxX && y <= l.maxY
}
