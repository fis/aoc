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
	"bytes"
	"fmt"
	"io"
	"os"
)

// A Level models a two-dimensional map of ASCII character cells, similar to a roguelike level.
// The origin of the coordinate system is the top-left cell; X grows right, Y grows down.
type Level struct {
	min, max         P
	empty            byte
	coreMin, coreMax P
	core             [][]byte
	spill            map[P]byte
}

// EmptyLevel returns a level filled with empty space matching the provided symbol.
// The given bounds determine the densely allocated space and the initial bounds of the level.
func EmptyLevel(min, max P, empty byte) *Level {
	w, h := max.X-min.X+1, max.Y-min.Y+1
	l := &Level{
		min: min, max: max,
		empty:   empty,
		coreMin: min, coreMax: max,
		core:  make([][]byte, h),
		spill: make(map[P]byte),
	}
	for i := range l.core {
		row := make([]byte, w)
		for j := range row {
			row[j] = empty
		}
		l.core[i] = row
	}
	return l
}

// SparseLevel returns a level with no densely allocated space at all.
// The level bounds are set to contain only the origin point.
func SparseLevel(origin P, empty byte) *Level {
	return &Level{
		min: origin, max: origin,
		empty:   empty,
		coreMin: origin, coreMax: P{origin.X - 1, origin.Y - 1}, // empty region
		core:  nil,
		spill: make(map[P]byte),
	}
}

// ReadLevel reads the contents of a text file into a level. Character cells outside the contents of
// the file are considered to be the specified empty byte.
func ReadLevel(path string, empty byte) (*Level, error) {
	data, err := os.ReadFile(path)
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
	x, y, maxX, maxY := min.X, min.Y, 0, 0
	for _, b := range data {
		if b == '\n' {
			x = min.X
			y++
			continue
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		x++
	}
	w, h := maxX-min.X+1, maxY-min.Y+1
	core := make([][]byte, h)
	for i := range core {
		row := make([]byte, w)
		for j := range row {
			row[j] = empty
		}
		core[i] = row
	}
	i, j := 0, 0
	for _, b := range data {
		if b == '\n' {
			j = 0
			i++
			continue
		}
		if b != empty {
			core[i][j] = b
		}
		j++
	}
	max := P{maxX, maxY}
	return &Level{
		min:     min,
		max:     max,
		empty:   empty,
		coreMin: min,
		coreMax: max,
		core:    core,
		spill:   make(map[P]byte),
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
	w, h := l.coreMax.X-l.coreMin.X+1, l.coreMax.Y-l.coreMin.Y+1
	c := &Level{
		min:     l.min,
		max:     l.max,
		empty:   l.empty,
		coreMin: l.coreMin,
		coreMax: l.coreMax,
		core:    make([][]byte, h),
		spill:   make(map[P]byte),
	}
	for i := range c.core {
		c.core[i] = make([]byte, w)
		copy(c.core[i], l.core[i])
	}
	for k, v := range l.spill {
		c.spill[k] = v
	}
	return c
}

// At returns the byte at the given coordinates.
func (l *Level) At(x, y int) byte {
	if x < l.min.X || y < l.min.Y || x > l.max.X || y > l.max.Y {
		return l.empty
	}
	if x >= l.coreMin.X && y >= l.coreMin.Y && x <= l.coreMax.X && y <= l.coreMax.Y {
		return l.core[y-l.coreMin.Y][x-l.coreMin.X]
	} else if b, ok := l.spill[P{x, y}]; ok {
		return b
	}
	return l.empty
}

// CoreAt returns the byte at the given coordinates, which must be within the densely stored region (see CoreBounds).
func (l *Level) CoreAt(x, y int) byte {
	return l.core[y-l.coreMin.Y][x-l.coreMin.X]
}

// Row returns a []byte with the contents of a single contiguous row from the level.
// If this is within the core area, the returned slice will share storage with the level.
func (l *Level) Row(x1, x2, y int) []byte {
	if x1 >= l.coreMin.X && y >= l.coreMin.Y && x2 <= l.coreMax.X && y <= l.coreMax.Y {
		return l.core[y-l.coreMin.Y][x1-l.coreMin.X : x2-l.coreMin.X+1]
	}
	row := make([]byte, x2-x1+1)
	for x := x1; x <= x2; x++ {
		row[x-x1] = l.At(x, y)
	}
	return row
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
	if x >= l.coreMin.X && y >= l.coreMin.Y && x <= l.coreMax.X && y <= l.coreMax.Y {
		l.core[y-l.coreMin.Y][x-l.coreMin.X] = b
		return
	}
	if b == l.empty {
		delete(l.spill, P{x, y})
		return
	}
	l.spill[P{x, y}] = b
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

// CoreSet sets the byte at the given coordinates, which must be within the densely stored region (see CoreBounds).
func (l *Level) CoreSet(x, y int, b byte) {
	l.core[y-l.coreMin.Y][x-l.coreMin.X] = b
}

// Bounds returns the top-left and bottom-right corners of the level's bounding box. See InBounds
// for the definition.
func (l *Level) Bounds() (min, max P) {
	return l.min, l.max
}

// Size returns the width and height of the bounding box of the level (see InBounds).
func (l *Level) Size() (w, h int) {
	return l.max.X - l.min.X + 1, l.max.Y - l.min.Y + 1
}

// InBounds returns true if the given coordinates are within the bounding box of the level. The
// bounds will grow to accommodate new non-empty characters, but will never shrink even if those
// characters are later overwritten to be empty.
func (l *Level) InBounds(x, y int) bool {
	return x >= l.min.X && y >= l.min.Y && x <= l.max.X && y <= l.max.Y
}

// CoreBounds returns the top-left and bottom-right corners of the core (densely stored) area.
func (l *Level) CoreBounds() (min, max P) {
	return l.coreMin, l.coreMax
}

// Range calls the callback function for all non-empty cells in the level.
func (l *Level) Range(cb func(x, y int, b byte)) {
	for i, row := range l.core {
		for j, b := range row {
			if b != l.empty {
				cb(l.coreMin.X+j, l.coreMin.Y+i, b)
			}
		}
	}
	for p, b := range l.spill {
		cb(p.X, p.Y, b)
	}
}

// Find locates the coordinates of a byte, which must be unique on the level.
func (l *Level) Find(key byte) (x, y int, found bool) {
	for i, row := range l.core {
		for j, b := range row {
			if b == key {
				if found {
					return 0, 0, false
				}
				x, y, found = l.coreMin.X+j, l.coreMin.Y+i, true
			}
		}
	}
	for p, b := range l.spill {
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

// FixedLevel is a restricted subset of Level, with less overhead for some operations.
//
// Limitations:
//   - The bounding box of the level has a predefined size based on the initial contents.
//   - The northwest corner is fixed at (0, 0).
type FixedLevel struct {
	// W and H are the level size.
	W, H int
	// Data is a row-major representation of the level data in a W*H length array.
	Data []byte
}

var singleNL = []byte{'\n'}

// ParseFixedLevel converts the input to a level structure.
// All lines are expected to be equally long.
func ParseFixedLevel(allData []byte) *FixedLevel {
	w, h := bytes.IndexByte(allData, '\n'), bytes.Count(allData, singleNL)
	data := make([]byte, w*h)
	for src, dst := 0, 0; dst < len(data); {
		eol := bytes.IndexByte(allData[src:], '\n')
		copy(data[dst:], allData[src:src+eol])
		src += eol + 1
		dst += w
	}
	return &FixedLevel{W: w, H: h, Data: data}
}

// At returns the value at the given coordinates.
func (l *FixedLevel) At(x, y int) byte { return l.Data[y*l.W+x] }

// Row returns the contents of an entire single row of the level.
func (l *FixedLevel) Row(y int) []byte { return l.Data[y*l.W : (y+1)*l.W] }

// Set assigns a new value at the given coordinates.
func (l *FixedLevel) Set(x, y int, b byte) { l.Data[y*l.W+x] = b }
