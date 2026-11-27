// Package day24 solves AoC 2024 day 24.
package day24

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2024, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	c := parseCircuit(lines)
	part1 := c.evaluate()
	part2 := c.analyze()
	return []string{
		strconv.Itoa(part1),
		part2,
	}, nil
}

type circuit struct {
	gates        map[gateID]*gate
	wires        map[gateID]byte
	swapA, swapB gateID
}

type gate struct {
	op         logicOp
	in1, in2   gateID
	out1, out2 gateID
	nout       int
}

type gateID uint16

const (
	extBase gateID = 31 << 10
	noGate  gateID = math.MaxUint16
)

func (c *circuit) evaluate() (result int) {
	for bit := range 100 {
		gate := extGate('z', bit)
		if _, ok := c.gates[gate]; !ok {
			break
		}
		result |= int(c.compute(gate)) << bit
	}
	return result
}

func (c *circuit) compute(gate gateID) (v byte) {
	if v, ok := c.wires[gate]; ok {
		return v
	}
	g := c.gates[gate]
	a, b := c.compute(g.in1), c.compute(g.in2)
	switch g.op {
	case opAND:
		v = a & b
	case opOR:
		v = a | b
	case opXOR:
		v = a ^ b
	}
	c.wires[gate] = v
	return v
}

func (c *circuit) analyze() string {
	size := 0
	for {
		if _, ok := c.gates[extGate('z', size+1)]; !ok {
			break
		}
		size++
	}

	// expected structure:
	// - the first gates reachable from the input bits (xN, yN) are a full adder
	// - there's 5 gates in the full adder: and1, xor1, and2, xor2, and or
	// - the two input bits are the inputs of both and1 and xor1
	// - xor1 and carry-in are the inputs of both and2 and xor2
	// - and1 and and2 are the two inputs of the or gate
	// - xor2 is the output bit (zN), or is next carry-in

	var swaps []gateID

	x0u := c.outputs(extGate('x', 0))
	carryIn := fn.If(!x0u[0].isExt('z'), x0u[0], x0u[1])
nextBit:
	for bit := 1; bit < size; bit++ {
		c.resetSwap()
		cand := c.outputs(extGate('x', bit))
		for len(cand) < 5 {
			for _, cg := range cand {
				for _, ng := range c.outputs(cg) {
					if !slices.Contains(cand, ng) {
						cand = append(cand, ng)
					}
				}
			}
		}
		for swapA := 0; swapA < 4; swapA++ {
			for swapB := swapA; swapB < 5; swapB++ {
				if swapA > 0 && swapB == swapA {
					continue // already tested the no-swaps case
				}
				if swapA != swapB {
					c.swap(cand[swapA], cand[swapB])
				}
				xu, yu := c.outputs(extGate('x', bit)), c.outputs(extGate('y', bit))
				if len(xu) != 2 || len(yu) != 2 || xu[0] != yu[0] || xu[1] != yu[1] {
					continue // x and y not connected to the same pair of gates
				}
				and1, xor1 := xu[0], xu[1]
				if c.gate(and1).op == opXOR && c.gate(xor1).op == opAND {
					and1, xor1 = xor1, and1
				} else if c.gate(and1).op != opAND || c.gate(xor1).op != opXOR {
					continue // x and y not connected to a plausible half adder
				}
				if and1.isExt('z') || xor1.isExt('z') {
					continue // first half adder output connected directly to output bit
				}
				su, cu := c.outputs(xor1), c.outputs(carryIn)
				if len(su) != 2 || len(cu) != 2 || su[0] != cu[0] || su[1] != cu[1] {
					continue // first sum and carry-in not connected to the same pair of gates
				}
				and2, xor2 := su[0], su[1]
				if c.gate(and2).op == opXOR && c.gate(xor2).op == opAND {
					and2, xor2 = xor2, and2
				} else if c.gate(and1).op != opAND || c.gate(xor1).op != opXOR {
					continue // first sum and carry-in not connected to a plausible half adder
				}
				if !xor2.isExt('z') {
					continue // second sum *not* connected directly to output bit
				}
				au := c.outputs(and2)
				if len(au) != 1 || c.gate(au[0]).op != opOR {
					continue // second half adder carry not going only to an or gate
				}
				carryIn = au[0]
				if swapA != swapB {
					swaps = append(swaps, cand[swapA], cand[swapB])
				}
				continue nextBit
			}
		}
		panic("circuit not fixable by a swap")
	}

	slices.Sort(swaps)
	return strings.Join(fn.Map(swaps, gateName), ",")
}

func (c *circuit) gate(id gateID) *gate {
	return c.gates[c.swapped(id)]
}

func (c *circuit) outputs(id gateID) (out []gateID) {
	g := c.gates[id]
	switch g.nout {
	case 1:
		return []gateID{c.swapped(g.out1)}
	case 2:
		out = []gateID{c.swapped(g.out1), c.swapped(g.out2)}
		if out[0] > out[1] {
			out[0], out[1] = out[1], out[0]
		}
		return out
	default:
		return nil
	}
}

func (c *circuit) swap(a, b gateID) {
	c.swapA, c.swapB = a, b
}

func (c *circuit) resetSwap() {
	c.swapA, c.swapB = noGate, noGate
}

func (c *circuit) swapped(id gateID) gateID {
	switch id {
	case c.swapA:
		return c.swapB
	case c.swapB:
		return c.swapA
	default:
		return id
	}
}

func (g *gate) addOut(out gateID) {
	switch g.nout {
	case 0:
		g.out1, g.nout = out, 1
	case 1:
		g.out2, g.nout = out, 2
	default:
		panic("out-degree exceeded")
	}
}

func intGate(name string) gateID {
	return gateID(name[0]-'a')<<10 | gateID(name[1]-'a')<<5 | gateID(name[2]-'a')
}

func extGate(key byte, bit int) gateID {
	return extBase | gateID(key-'x')<<7 | gateID(bit)
}

func anyGate(name string) gateID {
	if name[0] >= 'x' && name[0] <= 'z' {
		return extGate(name[0], 10*int(name[1]-'0')+int(name[2]-'0'))
	}
	return intGate(name)
}

func (id gateID) isExt(key byte) bool {
	return id >= extGate(key, 0) && id < extGate(key+1, 0)
}

func gateName(id gateID) string {
	var name [3]byte
	if id == noGate {
		return "(nil)"
	}
	if id >= extGate('x', 0) && id < extGate('z'+1, 0) {
		name[0] = 'x' + byte((id>>7)&3)
		name[1] = '0' + byte((id&0x7f)/10)
		name[2] = '0' + byte((id&0x7f)%10)
	} else {
		name[0] = 'a' + byte(id>>10)
		name[1] = 'a' + byte((id>>5)&0x1f)
		name[2] = 'a' + byte(id&0x1f)
	}
	return string(name[:])
}

type logicOp int

const (
	opAND logicOp = iota
	opOR
	opXOR
	opInput
)

func parseCircuit(lines []string) *circuit {
	ops := map[string]logicOp{
		"AND": opAND,
		"OR":  opOR,
		"XOR": opXOR,
	}

	gates := make(map[gateID]*gate)
	wires := make(map[gateID]byte)
	for _, line := range lines {
		switch {
		case len(line) < 6:
			continue
		case line[3] == ':':
			gates[anyGate(line[:3])] = &gate{op: opInput, in1: noGate, in2: noGate}
			wires[anyGate(line[:3])] = line[5] - '0'
		default:
			words := util.Words(line)
			gates[anyGate(words[4])] = &gate{op: ops[words[1]], in1: anyGate(words[0]), in2: anyGate(words[2])}
		}
	}
	for id, g := range gates {
		for _, in := range []gateID{g.in1, g.in2} {
			if in != noGate {
				gates[in].addOut(id)
			}
		}
	}
	return &circuit{gates: gates, wires: wires}
}

/*

circuit structure:

x00 XOR y00 -> z00
x00 AND y00 -> c0 (wrd)

x01 XOR y01 -> a1a (npf) *
x01 AND y01 -> c1a (tvb) *
a1a XOR c0  -> z01
a1a AND c0  -> c1b (prh) *
c1a OR  c1b -> c1 (nkn)

...

x44 XOR y44 -> a44a (bhg)
x44 AND y44 -> c44a (gvw)
a44a XOR c43 -> z44
a44a AND c43 -> c44b (mqf)
c44a OR  c44b -> z45

45 input bits, 46 output bits, should need 44*5 + 2 = 222 gates, checks out

fixing it manually:
- z10/kmb look like a swap?
- z15/tvp look like a swap?
- z25/dpg look like a swap?
- mmf/vdk look like a swap?

=> dpg,kmb,mmf,tvp,vdk,z10,z15,z25

*/
