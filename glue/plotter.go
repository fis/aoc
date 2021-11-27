// Copyright 2020 Google LLC
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
