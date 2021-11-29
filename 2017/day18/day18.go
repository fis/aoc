// Copyright 2021 Google LLC
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

// Package day18 solves AoC 2017 day 18.
package day18

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 18, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(set|add|mul|mod|snd|rcv|jgz) ([a-z]|-?\d+)(?: ([a-z]|-?\d+))?$`,
	})
}

func solve(code [][]string) ([]string, error) {
	prog := parseCode(code)
	p1 := part1(prog)
	p2 := part2(prog)
	return glue.Ints(p1, p2), nil
}

func part1(prog []instruction) int {
	lastSnd := -1
	snd := func(val int) { lastSnd = val }
	rcv := func(dst int) (val int, terminate bool) {
		if dst == 0 {
			return 0, false
		} else {
			return lastSnd, true
		}
	}
	st := state{snd: snd, rcv: rcv}
	for st.ip >= 0 && st.ip < len(prog) {
		if terminate := st.step(prog[st.ip]); terminate {
			break
		}
	}
	return lastSnd
}

func part2(code []instruction) int {
	const (
		reqSnd int = iota
		reqRcv
		reqTrm
	)
	type request struct {
		code int
		prog int
		val  int
	}
	type response struct {
		val       int
		terminate bool
	}
	reqChan := make(chan request)
	respChans := [2]chan response{make(chan response), make(chan response)}

	for i := 0; i < 2; i++ {
		progId := i
		snd := func(val int) {
			reqChan <- request{code: reqSnd, prog: progId, val: val}
		}
		rcv := func(int) (val int, terminate bool) {
			reqChan <- request{code: reqRcv, prog: progId}
			resp := <-respChans[progId]
			return resp.val, resp.terminate
		}
		st := state{snd: snd, rcv: rcv}
		st.regs['p'-'a'] = progId
		go func() {
			for st.ip >= 0 && st.ip < len(code) {
				if terminate := st.step(code[st.ip]); terminate {
					break
				}
			}
			reqChan <- request{code: reqTrm, prog: progId}
		}()
	}

	const (
		progRunning int = iota
		progWaiting
		progTerminated
	)
	prog := [2]struct {
		state     int
		inbox     []int
		sendCount int
	}{}
	for prog[0].state != progTerminated || prog[1].state != progTerminated {
		req := <-reqChan
		other := 1 - req.prog
		switch req.code {
		case reqSnd:
			prog[req.prog].sendCount++
			if prog[other].state == progWaiting {
				respChans[other] <- response{val: req.val}
				prog[other].state = progRunning
			} else {
				prog[other].inbox = append(prog[other].inbox, req.val)
			}
		case reqRcv:
			if len(prog[req.prog].inbox) > 0 {
				respChans[req.prog] <- response{val: prog[req.prog].inbox[0]}
				prog[req.prog].inbox = prog[req.prog].inbox[1:]
			} else if prog[other].state != progRunning {
				respChans[req.prog] <- response{terminate: true}
				if prog[other].state == progWaiting {
					respChans[other] <- response{terminate: true}
					prog[other].state = progRunning
				}
			} else {
				prog[req.prog].state = progWaiting
			}
		case reqTrm:
			prog[req.prog].state = progTerminated
			if prog[other].state == progWaiting {
				respChans[other] <- response{terminate: true}
				prog[other].state = progRunning
			}
		}
	}

	return prog[1].sendCount
}

type opCode int

const (
	opSet opCode = iota
	opAdd
	opMul
	opMod
	opSnd
	opRcv
	opJgz
)

var opMap = map[string]struct {
	code   opCode
	dstVal bool
	oneArg bool
}{
	"set": {code: opSet},
	"add": {code: opAdd},
	"mul": {code: opMul},
	"mod": {code: opMod},
	"snd": {code: opSnd, dstVal: true, oneArg: true},
	"rcv": {code: opRcv, dstVal: true, oneArg: true},
	"jgz": {code: opJgz, dstVal: true},
}

type instruction struct {
	op     opCode
	dst    int
	dstReg bool
	src    int
	srcReg bool
}

type state struct {
	ip   int
	regs [26]int
	snd  func(int)
	rcv  func(dst int) (val int, terminate bool)
}

func (s *state) step(inst instruction) (terminate bool) {
	dst, src := inst.dst, inst.src
	if inst.dstReg {
		dst = s.regs[dst]
	}
	if inst.srcReg {
		src = s.regs[src]
	}
	switch inst.op {
	case opSet:
		s.regs[inst.dst] = src
	case opAdd:
		s.regs[inst.dst] += src
	case opMul:
		s.regs[inst.dst] *= src
	case opMod:
		s.regs[inst.dst] %= src
	case opSnd:
		s.snd(dst)
	case opRcv:
		s.regs[inst.dst], terminate = s.rcv(dst)
	case opJgz:
		if dst > 0 {
			s.ip += src
		} else {
			s.ip++
		}
	}
	if inst.op != opJgz {
		s.ip++
	}
	return terminate
}

func parseCode(code [][]string) (insts []instruction) {
	insts = make([]instruction, len(code))
	for i, row := range code {
		opInfo := opMap[row[0]]
		insts[i] = instruction{op: opInfo.code}
		if len(row[1]) == 1 && row[1][0] >= 'a' && row[1][0] <= 'z' {
			insts[i].dst = int(row[1][0] - 'a')
			insts[i].dstReg = true
		} else if opInfo.dstVal {
			insts[i].dst, _ = strconv.Atoi(row[1])
		} else {
			panic("invalid instruction: dst reg required: " + strings.Join(row, " "))
		}
		if opInfo.oneArg != (len(row[2]) == 0) {
			panic("invalid instruction: arg count mismatch")
		} else if len(row[2]) == 1 && row[2][0] >= 'a' && row[2][0] <= 'z' {
			insts[i].src = int(row[2][0] - 'a')
			insts[i].srcReg = true
		} else if len(row[2]) > 0 {
			insts[i].src, _ = strconv.Atoi(row[2])
		}
	}
	return insts
}
