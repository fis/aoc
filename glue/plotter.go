package glue

import (
	"bufio"
	"io"

	"github.com/fis/aoc/util"
)

// LinePlotter wraps a plotter that wants the lines of the input as strings.
type LinePlotter func([]string, io.Writer) error

// Plot implements the Plotter interface.
func (p LinePlotter) Plot(r io.Reader, w io.Writer) error {
	data, err := util.ScanAll(r, bufio.ScanLines)
	if err != nil {
		return err
	}
	return p(data, w)
}

type ChunkPlotter func([]string, io.Writer) error

// Plot implements the Plotter interface.
func (p ChunkPlotter) Plot(r io.Reader, w io.Writer) error {
	data, err := util.ScanAll(r, util.ScanChunks)
	if err != nil {
		return err
	}
	return p(data, w)
}
