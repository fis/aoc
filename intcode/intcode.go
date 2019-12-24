// Package intcode implements the AoC 2019 Intcode language.
package intcode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc2019-go/util"
)

// Load will read an Intcode program in the standard format (comma-separated integers) from a text
// file.
func Load(path string) ([]int64, error) {
	lines, err := util.ReadLines(path)
	if err != nil {
		return nil, err
	}
	var p []int64
	for _, line := range lines {
		for _, num := range strings.Split(line, ",") {
			val, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing Intcode from %q: %v", path, err)
			}
			p = append(p, val)
		}
	}
	return p, nil
}

// An Intcode VM represents the state of an Intcode computer.
type VM struct {
	data   []int64
	ip     int
	base   int
	Stdin  Reader
	Stdout Writer
}

// A Reader provides input capabilities to an Intcode computer.
type Reader interface {
	Read() int64
}

// A Writer provides output capabilities to an Intcode computer.
type Writer interface {
	Write(val int64)
}

type opcode struct {
	act    func(vm *VM, args []arg)
	narg   int
	jump   bool
	input  bool
	output bool
}

type arg struct {
	val  int64
	mode argMode
}

type argMode uint

const (
	argPos argMode = iota
	argImm
	argRel
)

const maxArgs = 3

// LoadFile resets the computer and loads a program from the named file to its memory.
func (vm *VM) LoadFile(path string) error {
	p, err := Load(path)
	if err != nil {
		return err
	}
	vm.Use(p)
	return nil
}

// Load resets the computer and initializes its memory to be a copy of the program.
func (vm *VM) Load(p []int64) {
	vm.data = make([]int64, len(p))
	for i, val := range p {
		vm.data[i] = val
	}
	vm.Reset()
}

// Use resets the computer and uses the program directly as its memory.
func (vm *VM) Use(p []int64) {
	vm.data = p
	vm.Reset()
}

// Reset resets the IP and base pointers, but not the memory.
func (vm *VM) Reset() {
	vm.ip = 0
	vm.base = 0
}

// Run executes the current program until it halts, reading input from a slice. All output generated
// by the program will be returned in the output slice.
func (vm *VM) Run(in []int64) []int64 {
	pr, pw := vm.getIO()
	defer vm.setIO(pr, pw)
	var out []int64
	vm.setIO(&sliceReader{&in}, &sliceWriter{&out})
	vm.execute(-1)
	return out
}

// Mem returns a pointer to the specified memory cell of the computer.
func (vm *VM) Mem(offset int) *int64 {
	vm.page(offset)
	return &vm.data[offset]
}

// Dump returns a copy of the computer's memory.
func (vm *VM) Dump() []int64 {
	return append([]int64{}, vm.data...)
}

func (vm *VM) execute(steps int) {
	var args [maxArgs]arg
	for steps != 0 {
		op := vm.fetch(&args)
		if op == nil {
			return
		}
		op.act(vm, args[0:op.narg])
		if !op.jump {
			vm.ip += 1 + op.narg
		}
		if steps > 0 {
			steps--
		}
	}
}

func (vm *VM) fetch(args *[maxArgs]arg) *opcode {
	vm.page(vm.ip)
	inst := uint64(vm.data[vm.ip])
	op, ok := opcodes[inst%100]
	if !ok {
		return nil
	}
	vm.page(vm.ip + op.narg)
	inst /= 100
	for i := 0; i < op.narg; i++ {
		args[i].val = vm.data[vm.ip+1+i]
		args[i].mode = argMode(inst % 10)
		inst /= 10
	}
	return op
}

func (vm *VM) read(a arg) int64 {
	var i int
	switch a.mode {
	case argPos:
		i = 0
	case argImm:
		return a.val
	case argRel:
		i = vm.base
	}
	i += int(a.val)
	vm.page(i)
	return vm.data[i]
}

func (vm *VM) write(a arg, val int64) {
	var i int
	switch a.mode {
	case argPos:
		i = 0
	case argImm:
		panic("Intcode: write in immediate mode")
	case argRel:
		i = vm.base
	}
	i += int(a.val)
	vm.page(i)
	vm.data[i] = val
}

func (vm *VM) getIO() (Reader, Writer) {
	return vm.Stdin, vm.Stdout
}

func (vm *VM) setIO(r Reader, w Writer) {
	vm.Stdin, vm.Stdout = r, w
}

func (vm *VM) reader() Reader {
	r := vm.Stdin
	if r == nil {
		r = &interactive
	}
	return r
}

func (vm *VM) writer() Writer {
	w := vm.Stdout
	if w == nil {
		w = &interactive
	}
	return w
}

func (vm *VM) page(i int) {
	if i < 0 {
		panic(fmt.Sprintf("Intcode: page fault: %d", i))
	}
	if i >= len(vm.data) {
		vm.data = append(vm.data, make([]int64, i-len(vm.data)+1)...)
	}
}

var opcodes = map[uint64]*opcode{
	1: &opcode{act: (*VM).opAdd, narg: 3},
	2: &opcode{act: (*VM).opMul, narg: 3},
	3: &opcode{act: (*VM).opIn, narg: 1, input: true},
	4: &opcode{act: (*VM).opOut, narg: 1, output: true},
	5: &opcode{act: (*VM).opJNZ, narg: 2, jump: true},
	6: &opcode{act: (*VM).opJZ, narg: 2, jump: true},
	7: &opcode{act: (*VM).opSetLt, narg: 3},
	8: &opcode{act: (*VM).opSetEq, narg: 3},
	9: &opcode{act: (*VM).opSetB, narg: 1},
}

func (vm *VM) opAdd(args []arg) {
	vm.write(args[2], vm.read(args[0])+vm.read(args[1]))
}

func (vm *VM) opMul(args []arg) {
	vm.write(args[2], vm.read(args[0])*vm.read(args[1]))
}

func (vm *VM) opIn(args []arg) {
	vm.write(args[0], vm.reader().Read())
}

func (vm *VM) opOut(args []arg) {
	vm.writer().Write(vm.read(args[0]))
}

func (vm *VM) opJNZ(args []arg) {
	if vm.read(args[0]) != 0 {
		vm.ip = int(vm.read(args[1]))
	} else {
		vm.ip += 3
	}
}

func (vm *VM) opJZ(args []arg) {
	if vm.read(args[0]) == 0 {
		vm.ip = int(vm.read(args[1]))
	} else {
		vm.ip += 3
	}
}

func (vm *VM) opSetLt(args []arg) {
	val := int64(0)
	if vm.read(args[0]) < vm.read(args[1]) {
		val = int64(1)
	}
	vm.write(args[2], val)
}

func (vm *VM) opSetEq(args []arg) {
	val := int64(0)
	if vm.read(args[0]) == vm.read(args[1]) {
		val = int64(1)
	}
	vm.write(args[2], val)
}

func (vm *VM) opSetB(args []arg) {
	vm.base += int(vm.read(args[0]))
}

type sliceReader struct {
	q *[]int64
}

func (r *sliceReader) Read() int64 {
	if len(*r.q) == 0 {
		return 0
	}
	val := (*r.q)[0]
	*r.q = (*r.q)[1:]
	return val
}

type sliceWriter struct {
	q *[]int64
}

func (w *sliceWriter) Write(val int64) {
	*w.q = append(*w.q, val)
}

type interactiveReaderWriter struct{}

func (*interactiveReaderWriter) Read() int64 {
	// TODO
	return 0
}

func (*interactiveReaderWriter) Write(val int64) {
	// TODO
}

var interactive = interactiveReaderWriter{}

// Walk executes the current program up to the first halt, input or output instruction. For the
// initial call, pass in nil as the walk token. If the return value is nil, the computer has
// halted. Otherwise, the computer is requesting either input or output: inspect the walk token to
// learn which, and to either provide the input or deal with the output, as required. Then call Walk
// again with the token to continue operation. This is intended for coöperative multitasking or
// other interleaving of Intcode operation and surrounding logic.
func (vm *VM) Walk(token WalkToken) WalkToken {
	if it, ok := token.(*inputWalkToken); ok {
		vm.write(it.dst, it.val)
		vm.ip += 2
	} else if _, ok := token.(*outputWalkToken); ok {
		vm.ip += 2
	}

	var args [maxArgs]arg
	for {
		op := vm.fetch(&args)
		switch {
		case op == nil:
			return nil
		case op.input:
			return &inputWalkToken{dst: args[0]}
		case op.output:
			return &outputWalkToken{val: vm.read(args[0])}
		}
		op.act(vm, args[0:op.narg])
		if !op.jump {
			vm.ip += 1 + op.narg
		}
	}
}

type WalkToken interface {
	IsInput() bool
	IsOutput() bool
	ProvideInput(val int64)
	GetOutput() int64
}

type inputWalkToken struct {
	val int64
	dst arg
}

func (t *inputWalkToken) IsInput() bool {
	return true
}

func (t *inputWalkToken) IsOutput() bool {
	return false
}

func (t *inputWalkToken) ProvideInput(val int64) {
	t.val = val
}

func (t *inputWalkToken) GetOutput() int64 {
	return 0
}

func (t *inputWalkToken) String() string {
	return fmt.Sprintf("<in:%d@{%d,%d}>", t.val, t.dst.val, t.dst.mode)
}

type outputWalkToken struct {
	val int64
}

func (t *outputWalkToken) IsInput() bool {
	return false
}

func (t *outputWalkToken) IsOutput() bool {
	return true
}

func (t *outputWalkToken) ProvideInput(val int64) {
	// ignored
}

func (t *outputWalkToken) GetOutput() int64 {
	return t.val
}

func (t *outputWalkToken) String() string {
	return fmt.Sprintf("<out:%d>", t.val)
}
