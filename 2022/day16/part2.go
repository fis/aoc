// Copyright 2022 Google LLC
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

package day16

import (
	"github.com/fis/aoc/util/fn"
)

func releasePressure2(sum valveSummary, maxT int) (maxPressure int) {
	if len(sum.flowRates) > 16 {
		panic("only <= 16 non-zero flow rate valves supported")
	}
	n := uint8(len(sum.flowRates))

	var q bucketQ[move2]
	for i := uint8(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t1 := sum.initDist[i] + 1
			t2 := sum.initDist[j] + 1
			switch {
			case t1 < t2:
				q.push(t1, move2{st: state2{at1: i, at2: j, open: 1 << i}, wait1: 0, wait2: t2 - t1, pressure: (maxT - t1) * sum.flowRates[i]})
			case t1 == t2:
				q.push(t1, move2{st: state2{at1: i, at2: j, open: (1 << i) | (1 << j)}, wait1: 0, wait2: 0, pressure: (maxT - t1) * (sum.flowRates[i] + sum.flowRates[j])})
			case t1 > t2:
				q.push(t2, move2{st: state2{at1: i, at2: j, open: 1 << j}, wait1: t1 - t2, wait2: 0, pressure: (maxT - t2) * sum.flowRates[j]})
			}
		}
	}

	log := make(map[state2][]logEntry)
outerLoop:
	for q.len() > 0 {
		pt, p := q.pop()
		for _, e := range log[p.st] {
			if (e.time <= pt && e.pressure > p.pressure) || (e.time < pt && e.pressure >= p.pressure) {
				continue outerLoop
			}
		}
		i0, iN, j0, jN := uint8(0), n, uint8(0), n
		if p.wait1 > 0 {
			i0, iN = p.st.at1, p.st.at1+1
		}
		if p.wait2 > 0 {
			j0, jN = p.st.at2, p.st.at2+1
		}
	moveLoop:
		for i := i0; i < iN; i++ {
			for j := j0; j < jN; j++ {
				if p.st.open&((1<<i)|(1<<j)) != 0 {
					continue
				}
				if p.wait1 == 0 && i == p.st.at1 {
					continue
				}
				if p.wait2 == 0 && j == p.st.at2 {
					continue
				}
				t1 := pt + fn.If(p.wait1 > 0, p.wait1, sum.dist[p.st.at1][i]+1)
				t2 := pt + fn.If(p.wait2 > 0, p.wait2, sum.dist[p.st.at2][j]+1)
				if t1 >= 30 && t2 >= 30 {
					continue
				}
				var (
					nextt int
					next  move2
				)
				switch {
				case t1 < t2:
					nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open}, wait1: 0, wait2: t2 - t1, pressure: p.pressure}
				case t1 == t2:
					nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open}, wait1: 0, wait2: 0, pressure: p.pressure}
				case t1 > t2:
					nextt, next = t2, move2{st: state2{at1: i, at2: j, open: p.st.open}, wait1: t1 - t2, wait2: 0, pressure: p.pressure}
				}
				if next.wait1 == 0 && next.st.open&(1<<i) == 0 {
					next.st.open |= 1 << i
					next.pressure += (maxT - t1) * sum.flowRates[i]
				}
				if next.wait2 == 0 && next.st.open&(1<<j) == 0 {
					next.st.open |= 1 << j
					next.pressure += (maxT - t2) * sum.flowRates[j]
				}
				if next.st.at1 > next.st.at2 {
					next.st.at1, next.st.at2 = next.st.at2, next.st.at1
					next.wait1, next.wait2 = next.wait2, next.wait1
				}
				oldLog := log[next.st]
				// TODO: this probably should consider wait times as well somehow
				for _, e := range oldLog {
					if e.time <= nextt && e.pressure >= next.pressure {
						continue moveLoop
					}
				}
				newLog := oldLog[:0]
				for _, e := range oldLog {
					if e.time < nextt || e.pressure > next.pressure {
						newLog = append(newLog, e)
					}
				}
				newLog = append(newLog, logEntry{time: nextt, pressure: next.pressure})
				log[next.st] = newLog
				q.push(nextt, next)
			}
		}
	}

	for _, entries := range log {
		for _, e := range entries {
			if e.pressure > maxPressure {
				maxPressure = e.pressure
			}
		}
	}

	return maxPressure
}

type move2 struct {
	st           state2
	wait1, wait2 int
	pressure     int
}

type state2 struct {
	at1  uint8
	at2  uint8
	open uint16
}
