// Package day08 solves AoC 2019 day 8.
package day08

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/fis/aoc2019-go/util"
)

const (
	imgW = 25
	imgH = 6
)

func Solve(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 && data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	layers := decode(data, imgW, imgH)

	leastZeros, checksum := imgW*imgH+1, 0
	for _, layer := range layers {
		counts := countPixels(layer)
		if counts[0] < leastZeros {
			leastZeros = counts[0]
			checksum = counts[1] * counts[2]
		}
	}

	level := util.ParseLevelString("", ' ')
	for i, p := range render(layers) {
		if p == 1 {
			level.Set(i%imgW, i/imgW, '#')
		}
	}
	lines := level.Lines(util.P{0, 0}, util.P{imgW - 1, imgH - 1})

	return append([]string{strconv.Itoa(checksum)}, lines...), nil
}

func countPixels(layer []byte) [3]int {
	var counts [3]int
	for _, p := range layer {
		if p <= 2 {
			counts[p]++
		}
	}
	return counts
}

func decode(data []byte, w, h int) [][]byte {
	ls := w * h
	if len(data)%ls != 0 {
		panic(fmt.Sprintf("invalid image size: %d vs. %dx%d", len(data), w, h))
	}
	ln := len(data) / ls
	layers := make([][]byte, ln)
	for i, off := 0, 0; i < ln; i++ {
		layers[i] = make([]byte, ls)
		for j := 0; j < ls; j, off = j+1, off+1 {
			layers[i][j] = data[off] - '0'
		}
	}
	return layers
}

func render(layers [][]byte) []byte {
	img := make([]byte, len(layers[0]))
	for l := len(layers) - 1; l >= 0; l-- {
		layer := layers[l]
		for i, p := range layer {
			if p != 2 {
				img[i] = p
			}
		}
	}
	return img
}
