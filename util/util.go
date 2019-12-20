// Package util contains shared functions for several AoC days.
package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines returns the contents of a text file as a slice of strings representing the lines. The
// newline separators are not kept. The last line need not have a newline character at the end.
func ReadLines(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading lines: %v", err)
	}
	lines := strings.Split(string(data), "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines, nil
}

// ReadIntRows parses a text file formatted as one integer per line.
func ReadIntRows(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	var ints []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parsing ints from %s: %v", path, err)
		}
		ints = append(ints, i)
	}
	return ints, nil
}
