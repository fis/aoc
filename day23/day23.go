// Package day23 solves AoC 2019 day 23.
package day23

import (
	"strconv"

	"github.com/fis/aoc2019-go/intcode"
	"github.com/fis/aoc2019-go/util"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	var sw netSwitch
	p1, p2 := sw.run(prog)

	return []string{strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10)}, nil
}

const netSize = 50

type computer struct {
	addr  int
	vm    intcode.VM
	token intcode.WalkToken
	q     recvq
}

type recvq []packet

type packet struct {
	x, y int64
}

type netSwitch struct {
	machines     [netSize]computer
	natPacket    packet
	part1, part2 int64
	part1Set     bool
}

func (sw *netSwitch) run(prog []int64) (int64, int64) {
	for i := range sw.machines {
		m := &sw.machines[i]
		m.addr = i
		m.vm.Load(prog)
		m.vm.Walk(&m.token)
		m.token.ProvideInput(int64(i))
		for m.vm.Walk(&m.token) && m.token.IsOutput() {
			sw.send(m)
		}
	}

	for {
		idle := true
		for i := range sw.machines {
			m := &sw.machines[i]
			if m.q.empty() {
				m.token.ProvideInput(-1)
			} else {
				idle = false
				p := m.q.pop()
				m.token.ProvideInput(p.x)
				m.vm.Walk(&m.token)
				m.token.ProvideInput(p.y)
			}
			for m.vm.Walk(&m.token) && m.token.IsOutput() {
				idle = false
				sw.send(m)
			}
		}
		if idle {
			util.Diagf("NAT -> 0: %v\n", sw.natPacket)
			if sw.natPacket.y == sw.part2 {
				return sw.part1, sw.part2
			}
			sw.part2 = sw.natPacket.y
			sw.machines[0].q.push(sw.natPacket)
		}
	}
}

func (sw *netSwitch) send(m *computer) {
	var p packet
	addr := m.token.ReadOutput()
	m.vm.Walk(&m.token)
	p.x = m.token.ReadOutput()
	m.vm.Walk(&m.token)
	p.y = m.token.ReadOutput()
	util.Diagf("%d -> %d: %v\n", m.addr, addr, p)

	if addr == 255 {
		if !sw.part1Set {
			sw.part1Set = true
			sw.part1 = p.y
		}
		sw.natPacket = p
	} else {
		sw.machines[addr].q.push(p)
	}
}

func (q recvq) empty() bool {
	return len(q) == 0
}

func (q *recvq) push(p packet) {
	*q = append(*q, p)
}

func (q *recvq) pop() packet {
	p := (*q)[0]
	*q = (*q)[1:]
	if len(*q) == 0 {
		*q = nil
	}
	return p
}
