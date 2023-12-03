// Copyright 2023 Google LLC
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

package z80

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/z80/z80ex"
	"github.com/google/go-cmp/cmp"
)

func TestSolutions(t *testing.T) {
	tests, err := glue.FindAllTests("../testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		var progs [][]byte
		for part := 1; part <= 2; part++ {
			binFile := fmt.Sprintf("%04d/day%02d-%d.bin", test.Year, test.Day, part)
			prog, err := os.ReadFile(binFile)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					break
				}
				t.Fatalf("error loading solution: %v", err)
			}
			progs = append(progs, prog)
		}
		if len(progs) == 0 {
			continue
		}
		name := fmt.Sprintf("year=%04d/day=%02d", test.Year, test.Day)
		t.Run(name, func(t *testing.T) {
			testSolution(t, test, progs)
		})
	}
}

func testSolution(t *testing.T, test glue.TestCase, progs [][]byte) {
	input, err := os.Open(test.InputFile)
	if err != nil {
		t.Fatalf("error opening test input: %v", err)
	}
	defer input.Close()
	var output strings.Builder

	cpu := z80ex.NewCPU()
	defer cpu.Destroy()
	for _, prog := range progs {
		if _, err := input.Seek(0, 0); err != nil {
			t.Fatalf("error rewinding test input: %v", err)
		}
		cpu.Reset(true)
		cpu.WriteMem(prog, 0)
		cpu.Run(bufio.NewReader(input), &output)
	}

	got, want := util.Lines(output.String()), test.Want
	if len(got) > 0 && len(got) < len(want) {
		// Accept non-empty but truncated output in case only first part of puzzle has been solved.
		want = want[:len(got)]
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("output mismatch (-want +got):\n%s", diff)
	}
}
