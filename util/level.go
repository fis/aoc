// Copyright 2019 Google LLC
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

package util

import (
	"fmt"
	"io"
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
	return ParseLevelAt(data, empty, P{0, 0})
}

// ParseLevelAt parses a byte array into a level using a specified offset.
func ParseLevelAt(data []byte, empty byte, min P) *Level {
	level := make(map[P]byte)
	x, y, maxX, maxY := min.X, min.Y, 0, 0
	for _, b := range data {
		if b == '\n' {
			x = min.X
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
		min:   min,
		max:   P{maxX, maxY},
	}
}

// ParseLevelString parses a string into a level. See ReadLevel.
func ParseLevelString(data string, empty byte) *Level {
	return ParseLevel([]byte(data), empty)
}

// ParseLevelStringAt parses a string into a level using a specified offset.
func ParseLevelStringAt(data string, empty byte, min P) *Level {
	return ParseLevelAt([]byte(data), empty, min)
}

// Copy returns a deep copy of a level.
func (l *Level) Copy() *Level {
	c := &Level{data: make(map[P]byte), empty: l.empty, min: l.min, max: l.max}
	for k, v := range l.data {
		c.data[k] = v
	}
	return c
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

// Lines returns the contents of a part of the level as a list of strings.
func (l *Level) Lines(min, max P) []string {
	lines := make([]string, max.Y-min.Y+1)
	for y, yi := min.Y, 0; y <= max.Y; y, yi = y+1, yi+1 {
		line := make([]byte, max.X-min.X+1)
		for x, xi := min.X, 0; x <= max.X; x, xi = x+1, xi+1 {
			line[xi] = l.At(x, y)
		}
		lines[yi] = string(line)
	}
	return lines
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
func (l *Level) Bounds() (min, max P) {
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

// Write prints out the bytes in-bounds area of the level.
func (l *Level) Write(w io.Writer) error {
	min, max := l.Bounds()
	return l.WriteRect(w, min, max)
}

// WriteRect prints out the specified rectangle of the level.
func (l *Level) WriteRect(w io.Writer, min, max P) error {
	row := make([]byte, max.X-min.X+2)
	row[max.X-min.X+1] = '\n'
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			row[x-min.X] = l.At(x, y)
		}
		if _, err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}
