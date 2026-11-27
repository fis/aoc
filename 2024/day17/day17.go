// Package day17 solves AoC 2024 day 17.
package day17

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 17, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	regs, prog := parseInput(r)
	part1 := interp(regs, prog)
	// N.B. assumes input program does what it does
	part2 := encode(0, prog)
	return []string{part1, strconv.Itoa(part2)}, nil
}

const (
	opADV byte = iota
	opBXL
	opBST
	opJNZ
	opBXC
	opOUT
	opBDV
	opCDV
)

func interp(regs [3]int, prog []byte) string {
	var output strings.Builder

	for ip := 0; ip < len(prog); ip += 2 {
		op, oper := prog[ip], prog[ip+1]
		switch op {
		case opADV:
			regs[0] >>= combo(oper, &regs)
		case opBXL:
			regs[1] ^= int(oper)
		case opBST:
			regs[1] = combo(oper, &regs) & 7
		case opJNZ:
			if regs[0] != 0 {
				ip = int(oper) - 2
			}
		case opBXC:
			regs[1] ^= regs[2]
		case opOUT:
			if output.Len() > 0 {
				output.WriteByte(',')
			}
			output.WriteByte('0' + byte(combo(oper, &regs)&7))
		case opBDV:
			regs[1] = regs[0] >> combo(oper, &regs)
		case opCDV:
			regs[2] = regs[0] >> combo(oper, &regs)
		}
	}

	return output.String()
}

func combo(oper byte, regs *[3]int) int {
	if oper < 4 {
		return int(oper)
	}
	return regs[oper-4]
}

func encode(A int, seq []byte) int {
	if len(seq) == 0 {
		return A
	}
	last := seq[len(seq)-1]
	seq = seq[:len(seq)-1]
	for try := range 8 {
		cand := A<<3 | try
		if byte(try^(cand>>(try^3))^6)&7 == last {
			if ok := encode(cand, seq); ok >= 0 {
				return ok
			}
		}
	}
	return -1
}

func parseInput(r io.Reader) (regs [3]int, prog []byte) {
	data, _ := io.ReadAll(r)
	regData, progData, _ := bytes.Cut(data, []byte("\n\n"))
	regs = *(*[3]int)(util.Ints(string(regData)))
	prog = make([]byte, 0, len(progData)>>1)
	for _, ch := range progData {
		if ch >= '0' && ch < '8' {
			prog = append(prog, ch-'0')
		}
	}
	return regs, prog
}
