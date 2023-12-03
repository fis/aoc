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

// Package z80ex contains Cgo bindings to the z80ex Z80 CPU emulation library.
package z80ex

// #cgo LDFLAGS: -lz80ex
// #include "native.h"
import "C"
import (
	"io"
	"math"
	"runtime/cgo"

	"github.com/fis/aoc/util/fn"
)

const haltFlag = 0x8000_0000_0000_0000

type CPU struct {
	native *C.struct_cpu
}

type streams struct {
	r io.Reader
	w io.Writer
}

func NewCPU() *CPU {
	cpu := &CPU{
		native: C.go_z80ex_create(),
	}
	if cpu.native == nil {
		panic("out of memory")
	}
	return cpu
}

func (cpu *CPU) Destroy() {
	C.go_z80ex_destroy(cpu.native)
}

func (cpu *CPU) Reset(clearMem bool) {
	C.go_z80ex_reset(cpu.native, fn.If[C.int](clearMem, 1, 0))
}

func (cpu *CPU) WriteMem(data []byte, at int) {
	for i, v := range data {
		cpu.native.ram[at+i] = C.Z80EX_BYTE(v)
	}
}

func (cpu *CPU) Run(r io.Reader, w io.Writer) (steps uint64) {
	steps, _ = cpu.RunBounded(r, w, math.MaxUint64)
	return steps
}

func (cpu *CPU) RunBounded(r io.Reader, w io.Writer, maxT uint64) (steps uint64, halted bool) {
	h := cgo.NewHandle(streams{r: r, w: w})
	defer h.Delete()
	steps = uint64(C.go_z80ex_run(cpu.native, C.uintptr_t(h), C.uint64_t(maxT)))
	halted = steps&haltFlag != 0
	steps &^= haltFlag
	return steps, halted
}

func (cpu *CPU) Trace(r io.Reader, w io.Writer) (steps uint64) {
	h := cgo.NewHandle(streams{r: r, w: w})
	defer h.Delete()
	return uint64(C.go_z80ex_trace(cpu.native, C.uintptr_t(h)))
}

//export readGoInput
func readGoInput(h C.uintptr_t) C.uint8_t {
	s := cgo.Handle(h).Value().(streams)
	var buf [1]byte
	n, _ := s.r.Read(buf[:])
	if n < 1 {
		return 0
	}
	return C.uint8_t(buf[0])
}

//export writeGoOutput
func writeGoOutput(h C.uintptr_t, val C.uint8_t) {
	s := cgo.Handle(h).Value().(streams)
	buf := [1]byte{byte(val)}
	s.w.Write(buf[:])
}
