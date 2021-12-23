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

// Package day23 solves AoC 2021 day 23.
package day23

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	start := decodeState(lines)
	p1 := shortestPath(start, sortedState)
	startDeep := convertState(start)
	p2 := shortestDeepPath(startDeep, sortedDeepState)
	return glue.Ints(p1, p2), nil
}

func shortestPath(from, to amphiState) (energy int) {
	dist := make(map[amphiState]uint32)
	dist[from] = 0
	var (
		fringe amphiPathQ
		moves  []amphiPath
	)
	fringe.push(0, from)
	for {
		pe, p := fringe.pop()
		if p == to {
			return int(pe)
		}
		if oldBest, ok := dist[p]; ok && oldBest < pe {
			continue
		}
		moves = p.moves(moves)
		for _, q := range moves {
			qe := pe + q.energy
			if oldBest, ok := dist[q.state]; ok && oldBest <= qe {
				continue
			}
			dist[q.state] = qe
			fringe.push(qe, q.state)
		}
	}
}

func shortestDeepPath(from, to amphiDeepState) (energy int) {
	dist := make(map[amphiDeepState]uint32)
	dist[from] = 0
	var (
		fringe amphiDeepPathQ
		moves  []amphiDeepPath
	)
	fringe.push(0, from)
	for {
		pe, p := fringe.pop()
		if p == to {
			return int(pe)
		}
		if oldBest, ok := dist[p]; ok && oldBest < pe {
			continue
		}
		moves = p.moves(moves)
		for _, q := range moves {
			qe := pe + q.energy
			if oldBest, ok := dist[q.state]; ok && oldBest <= qe {
				continue
			}
			dist[q.state] = qe
			fringe.push(qe, q.state)
		}
	}
}

/*
amphiState represents the entire state of the amphipods, encoded as follows:

Let's denote the possible positions of the level an amphipod may be in as follows:

    #############
    #12.3.4.5.67#
    ###A#B#C#D###
      #a#b#c#d#
      #########

The state of each of the 15 labeled positions is encoded in a different nybble of a 60-bit word:

    0x7654321dDcCbBaA

Each nybble is either 0 (for empty), or a value between 4 .. 7 to denote an A-type .. D-type amphipod respectively.
*/
type amphiState uint64

const sortedState amphiState = 0x77665544

func (st amphiState) moves(buf []amphiPath) (moves []amphiPath) {
	stepCost := [4]uint32{1, 10, 100, 1000}
	moves = buf[:0]
	roomData, corrData := uint32(st), uint32(st>>32)
	// Out-moves: amphipod going from room to corridor.
	for room, rd := uint32(0), roomData; room < 4; room, rd = room+1, rd>>8 {
		if owner, occupant := 0x4|room, rd&0xff; occupant == 0 || occupant == (owner<<4) || occupant == (owner<<4)|owner {
			continue // either empty or no need to move
		}
		var occupant, nr, baseSteps uint32
		if rd&0xf == 0 {
			occupant, nr, baseSteps = (rd>>4)&0xf, 0, 1
		} else {
			occupant, nr, baseSteps = rd&0xf, rd&0xf0, 0
		}
		newRooms := (roomData & ^(uint32(0xff) << (8 * room))) | (nr << (8 * room))
		for left := uint32(0); left < room+2; left++ {
			cs := 4 * (room + 1 - left)
			if (corrData>>cs)&0xf != 0 {
				break // path blocked
			}
			steps := baseSteps + (left+1)*2
			if left == room+1 {
				steps--
			}
			newCorr := corrData | (occupant << cs)
			moves = append(moves, amphiPath{
				state:  amphiState(newRooms) | (amphiState(newCorr) << 32),
				energy: steps * stepCost[occupant&3],
			})
		}
		for right := uint32(0); room+right < 5; right++ {
			cs := 4 * (2 + room + right)
			if (corrData>>cs)&0xf != 0 {
				break // path blocked
			}
			steps := baseSteps + (right+1)*2
			if room+right == 4 {
				steps--
			}
			newCorr := corrData | (occupant << cs)
			moves = append(moves, amphiPath{
				state:  amphiState(newRooms) | (amphiState(newCorr) << 32),
				energy: steps * stepCost[occupant&3],
			})
		}
	}
	// In-moves: amphipod going from corridor to room.
	for corr, cd := uint32(0), corrData; cd != 0; corr, cd = corr+1, cd>>4 {
		occupant := cd & 0xf
		if occupant == 0 {
			continue // empty corridor
		}
		room := occupant & 3
		rd := (roomData >> (8 * room)) & 0xff
		if rd&0xf != 0 || ((rd>>4) != 0 && (rd>>4) != occupant) {
			continue // room full or occupied by someone who doesn't belong
		}
		var nr, steps uint32
		if rd&0xf0 == 0 {
			nr, steps = occupant<<4, 1
		} else {
			nr, steps = (occupant<<4)|occupant, 0
		}
		if corr <= room { // move right crossing at least one corridor tile
			path := (uint32(1) << (4 * (room - corr + 1))) - 1
			path <<= 4 * (corr + 1)
			if corrData&path != 0 {
				continue // path blocked
			}
			steps += (room - corr + 2) * 2
		} else if corr > room+1 { // move left crossing at least one corridor tile
			path := (uint32(1) << (4 * (corr - room - 2))) - 1
			path <<= 4 * (room + 2)
			if corrData&path != 0 {
				continue // path blocked
			}
			steps += (corr - room - 1) * 2
		} else { // move to a room we're right next to, no need to check for blockers
			steps += 2
		}
		if corr == 0 || corr == 6 {
			steps--
		}
		newRooms := roomData | (nr << (8 * room))
		newCorr := corrData & ^(uint32(0xf) << (4 * corr))
		moves = append(moves, amphiPath{
			state:  amphiState(newRooms) | (amphiState(newCorr) << 32),
			energy: steps * stepCost[occupant&3],
		})
	}
	return moves
}

func decodeState(lines []string) (st amphiState) {
	cells := [15]byte{
		lines[1][11], lines[1][10], lines[1][8], lines[1][6], lines[1][4], lines[1][2], lines[1][1],
		lines[3][9], lines[2][9], lines[3][7], lines[2][7], lines[3][5], lines[2][5], lines[3][3], lines[2][3],
	}
	for _, c := range cells {
		st <<= 4
		if c != '.' {
			st |= amphiState(4 + (c - 'A'))
		}
	}
	return st
}

func (st amphiState) String() string {
	return fmt.Sprintf("%015x", uint64(st))
}

type amphiPath struct {
	state  amphiState
	energy uint32
}

func (p amphiPath) String() string {
	return fmt.Sprintf("%v/%d", p.state, p.energy)
}

const bucketSpan = 16384 // power of 2 that's > that maximum edge length

type amphiPathQ struct {
	at      uint32
	buckets [bucketSpan]struct {
		energy uint32
		states []amphiState
	}
}

func (q *amphiPathQ) pop() (e uint32, st amphiState) {
	for len(q.buckets[q.at].states) == 0 {
		q.at = (q.at + 1) & (bucketSpan - 1)
	}
	states := q.buckets[q.at].states
	st = states[len(states)-1]
	q.buckets[q.at].states = states[:len(states)-1]
	return q.buckets[q.at].energy, st
}

func (q *amphiPathQ) push(e uint32, st amphiState) {
	i := e & (bucketSpan - 1)
	q.buckets[i].energy = e
	q.buckets[i].states = append(q.buckets[i].states, st)
}

/*
amphiDeepState represents the entire state of the amphipods for part 2, encoded as follows:

Let's denote the possible positions of the level an amphipod may be in as follows:

    #############
    #12.3.4.5.67#
    ###A#B#C#D###
      #X#Y#Z#W#
      #x#y#z#w#
      #a#b#c#d#
      #########

The state of each of the 7 labeled corridor positions is encoded in a nybble of a 28-bit word:

    0x7654321

And the state of each of the 16 labeled room positions is encoded in a nybble of a 64-bit word:

    0xdwWDczZCbyYBaxXA

Each nybble is either 0 (for empty), or a value between 4 .. 7 to denote an A-type .. D-type amphipod respectively.
*/
type amphiDeepState struct {
	corridor uint32
	rooms    uint64
}

var sortedDeepState = amphiDeepState{corridor: 0, rooms: 0x7777666655554444}

func (st amphiDeepState) moves(buf []amphiDeepPath) (moves []amphiDeepPath) {
	stepCost := [4]uint32{1, 10, 100, 1000}
	moves = buf[:0]
	roomData, corrData := st.rooms, st.corridor
	// Out-moves: amphipod going from room to corridor.
	for room, rd := uint32(0), roomData; room < 4; room, rd = room+1, rd>>16 {
		if ow, occ := 0x4|room, uint32(rd&0xffff); occ == 0 || occ == ow<<12 || occ == (ow<<12)|(ow<<8) || occ == (ow<<12)|(ow<<8)|(ow<<4) || occ == (ow<<12)|(ow<<8)|(ow<<4)|ow {
			continue // either empty or no need to move
		}
		var occupant, baseSteps uint32
		var nr uint64
		switch {
		case rd&0xfff == 0:
			occupant, nr, baseSteps = uint32((rd>>12)&0xf), 0, 3
		case rd&0xff == 0:
			occupant, nr, baseSteps = uint32((rd>>8)&0xf), rd&0xf000, 2
		case rd&0xf == 0:
			occupant, nr, baseSteps = uint32((rd>>4)&0xf), rd&0xff00, 1
		default:
			occupant, nr, baseSteps = uint32(rd&0xf), rd&0xfff0, 0
		}
		newRooms := (roomData & ^(uint64(0xffff) << (16 * room))) | (nr << (16 * room))
		for left := uint32(0); left < room+2; left++ {
			cs := 4 * (room + 1 - left)
			if (corrData>>cs)&0xf != 0 {
				break // path blocked
			}
			steps := baseSteps + (left+1)*2
			if left == room+1 {
				steps--
			}
			newCorr := corrData | (occupant << cs)
			moves = append(moves, amphiDeepPath{
				state:  amphiDeepState{corridor: newCorr, rooms: newRooms},
				energy: steps * stepCost[occupant&3],
			})
		}
		for right := uint32(0); room+right < 5; right++ {
			cs := 4 * (2 + room + right)
			if (corrData>>cs)&0xf != 0 {
				break // path blocked
			}
			steps := baseSteps + (right+1)*2
			if room+right == 4 {
				steps--
			}
			newCorr := corrData | (occupant << cs)
			moves = append(moves, amphiDeepPath{
				state:  amphiDeepState{corridor: newCorr, rooms: newRooms},
				energy: steps * stepCost[occupant&3],
			})
		}
	}
	// In-moves: amphipod going from corridor to room.
	for corr, cd := uint32(0), corrData; cd != 0; corr, cd = corr+1, cd>>4 {
		occupant := cd & 0xf
		if occupant == 0 {
			continue // empty corridor
		}
		room := occupant & 3
		rd := uint32((roomData >> (16 * room)) & 0xffff)
		if rd&0xf != 0 || ((rd>>4)&0xf != 0 && (rd>>4)&0xf != occupant) || ((rd>>8)&0xf != 0 && (rd>>8)&0xf != occupant) || ((rd>>12) != 0 && (rd>>12) != occupant) {
			continue // room full or occupied by someone who doesn't belong
		}
		var nr, steps uint32
		switch {
		case rd&0xf000 == 0:
			nr, steps = occupant<<12, 3
		case rd&0xf00 == 0:
			nr, steps = (rd&0xf000)|(occupant<<8), 2
		case rd&0xf0 == 0:
			nr, steps = (rd&0xff00)|(occupant<<4), 1
		default:
			nr, steps = (rd&0xfff0)|occupant, 0
		}
		if corr <= room { // move right crossing at least one corridor tile
			path := (uint32(1) << (4 * (room - corr + 1))) - 1
			path <<= 4 * (corr + 1)
			if corrData&path != 0 {
				continue // path blocked
			}
			steps += (room - corr + 2) * 2
		} else if corr > room+1 { // move left crossing at least one corridor tile
			path := (uint32(1) << (4 * (corr - room - 2))) - 1
			path <<= 4 * (room + 2)
			if corrData&path != 0 {
				continue // path blocked
			}
			steps += (corr - room - 1) * 2
		} else { // move to a room we're right next to, no need to check for blockers
			steps += 2
		}
		if corr == 0 || corr == 6 {
			steps--
		}
		newRooms := roomData | (uint64(nr) << (16 * room))
		newCorr := corrData & ^(uint32(0xf) << (4 * corr))
		moves = append(moves, amphiDeepPath{
			state:  amphiDeepState{corridor: newCorr, rooms: newRooms},
			energy: steps * stepCost[occupant&3],
		})
	}
	return moves
}

func convertState(shallow amphiState) (deep amphiDeepState) {
	rd := uint64(uint32(shallow))
	deep.corridor = uint32(shallow >> 32)
	// 0xdDcCbBaA to 0xd64Dc45Cb56Ba77A
	deep.rooms = (rd & 0xf) | ((rd & 0xff0) << 8) | ((rd & 0xff000) << 16) | ((rd & 0xff00000) << 24) | ((rd & 0xf0000000) << 32)
	deep.rooms |= 0x640045005600770
	return deep
}

func (st amphiDeepState) String() string {
	return fmt.Sprintf("%07x|%016x", st.corridor, st.rooms)
}

type amphiDeepPath struct {
	state  amphiDeepState
	energy uint32
}

func (p amphiDeepPath) String() string {
	return fmt.Sprintf("%v/%d", p.state, p.energy)
}

type amphiDeepPathQ struct {
	at      uint32
	buckets [bucketSpan]struct {
		energy uint32
		states []amphiDeepState
	}
}

func (q *amphiDeepPathQ) pop() (e uint32, st amphiDeepState) {
	for len(q.buckets[q.at].states) == 0 {
		q.at = (q.at + 1) & (bucketSpan - 1)
	}
	states := q.buckets[q.at].states
	st = states[len(states)-1]
	q.buckets[q.at].states = states[:len(states)-1]
	return q.buckets[q.at].energy, st
}

func (q *amphiDeepPathQ) push(e uint32, st amphiDeepState) {
	i := e & (bucketSpan - 1)
	q.buckets[i].energy = e
	q.buckets[i].states = append(q.buckets[i].states, st)
}
