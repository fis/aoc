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

package obs

import (
	"github.com/fis/aoc/2022/day16"
	"github.com/fis/aoc/util/ix"
)

func greedyBound(sum day16.ValveSummary, maxT int) (pressure int) {
	n := len(sum.FlowRates)
	at1, at2 := -1, -1
	wait1, wait2 := 0, 0
	open := 0
	for t := 0; t < maxT; t++ {
		if wait1 > 0 && wait2 > 0 {
			wait1, wait2 = wait1-1, wait2-1
			continue
		}
		if wait1 == 0 && wait2 == 0 {
			bestI, bestJ, bestT1, bestT2, bestP := 0, 0, 0, 0, 0
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if i == j || open&((1<<i)|(1<<j)) != 0 {
						continue
					}
					var d1, d2 int
					if at1 < 0 {
						d1 = sum.InitDist[i]
					} else {
						d1 = sum.Dist[at1][i]
					}
					if at2 < 0 {
						d2 = sum.InitDist[j]
					} else {
						d2 = sum.Dist[at2][j]
					}
					t1 := ix.Min(t+d1+1, maxT)
					t2 := ix.Min(t+d2+1, maxT)
					p := (maxT-t1)*sum.FlowRates[i] + (maxT-t2)*sum.FlowRates[j]
					if p > bestP {
						bestI, bestJ, bestT1, bestT2, bestP = i, j, t1, t2, p
					}
				}
			}
			pressure += bestP
			at1, at2, wait1, wait2 = bestI, bestJ, bestT1-t, bestT2-t
			open |= (1 << at1) | (1 << at2)
			continue
		} else if wait1 > 0 {
			at1, at2, wait2 = at2, at1, wait1-1
		} else {
			wait2 = wait2 - 1
		}
		bestI, bestT1, bestP := 0, 0, 0
		for i := 0; i < n; i++ {
			if open&(1<<i) != 0 {
				continue
			}
			t1 := ix.Min(t+sum.Dist[at1][i]+1, maxT)
			p := (maxT - t1) * sum.FlowRates[i]
			if p > bestP {
				bestI, bestT1, bestP = i, t1, p
			}
		}
		pressure += bestP
		at1, wait1 = bestI, bestT1-t
		open |= 1 << at1
	}
	return pressure
}

func releasePressure2(sum day16.ValveSummary, maxT int) (maxPressure int) {
	if len(sum.FlowRates) > maxValves {
		panic("too many valves")
	}
	n := uint8(len(sum.FlowRates))

	maxPressure = greedyBound(sum, maxT)

	maxFlow := [65536]int{}
	for i := 0; i < 65536; i++ {
		for j := uint8(0); j < n; j++ {
			if i&(1<<j) != 0 {
				maxFlow[i] += sum.FlowRates[j]
			}
		}
	}

	var (
		minD1 [maxValves]int
		minD2 [maxValves][maxValves]int
	)
	for i := uint8(0); i < n; i++ {
		minD1[i] = maxT
		for j := uint8(0); j < n; j++ {
			if i != j && sum.Dist[i][j] < minD1[i] {
				minD1[i] = sum.Dist[i][j]
			}
		}
	}
	for i := uint8(0); i < n; i++ {
		for j := uint8(0); j < n; j++ {
			minD2[i][j] = ix.Min(minD1[i], minD1[j])
		}
	}

	var q bucketQ[move2]
	for i := uint8(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t1 := sum.InitDist[i] + 1
			t2 := sum.InitDist[j] + 1
			switch {
			case t1 < t2:
				q.push(t1, move2{st: state2{at1: i, at2: j, open: 1 << i}, wait1: 0, wait2: uint8(t2 - t1), pressure: uint16((maxT - t1) * sum.FlowRates[i])})
			case t1 == t2:
				q.push(t1, move2{st: state2{at1: i, at2: j, open: (1 << i) | (1 << j)}, wait1: 0, wait2: 0, pressure: uint16((maxT - t1) * (sum.FlowRates[i] + sum.FlowRates[j]))})
			case t1 > t2:
				q.push(t2, move2{st: state2{at1: i, at2: j, open: 1 << j}, wait1: uint8(t1 - t2), wait2: 0, pressure: uint16((maxT - t2) * sum.FlowRates[j])})
			}
		}
	}

	log := [maxValves][maxValves][65536]logEntryList2{}
outerLoop:
	for q.len() > 0 {
		pt, p := q.pop()
		for _, e := range log[p.st.at1][p.st.at2][p.st.open] {
			if (int(e.time) <= pt && e.pressure > p.pressure) || (int(e.time) < pt && e.pressure >= p.pressure) {
				continue outerLoop
			}
		}
		if int(p.pressure) > maxPressure {
			maxPressure = int(p.pressure)
		}
		if p.wait1 == 0 && p.wait2 == 0 {
			if int(p.pressure)+(maxT-pt-minD2[p.st.at1][p.st.at2]-1)*maxFlow[^p.st.open] <= maxPressure {
				continue
			}
			for i := uint8(0); i < n; i++ {
				for j := uint8(0); j < n; j++ {
					if p.st.open&((1<<i)|(1<<j)) != 0 || i == p.st.at1 || i == p.st.at2 || j == p.st.at1 || j == p.st.at2 || i == j {
						continue
					}
					t1 := pt + sum.Dist[p.st.at1][i] + 1
					t2 := pt + sum.Dist[p.st.at2][j] + 1
					if t1 >= maxT && t2 >= maxT {
						continue
					}
					var (
						nextt int
						next  move2
					)
					switch {
					case t1 < t2:
						nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << i)}, wait1: 0, wait2: uint8(t2 - t1), pressure: p.pressure + uint16((maxT-t1)*sum.FlowRates[i])}
					case t1 == t2:
						nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << i) | (1 << j)}, wait1: 0, wait2: 0, pressure: p.pressure + uint16((maxT-t1)*(sum.FlowRates[i]+sum.FlowRates[j]))}
					case t1 > t2:
						nextt, next = t2, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << j)}, wait1: uint8(t1 - t2), wait2: 0, pressure: p.pressure + uint16((maxT-t2)*sum.FlowRates[j])}
					}
					if next.st.at1 > next.st.at2 {
						next.st.at1, next.st.at2 = next.st.at2, next.st.at1
						next.wait1, next.wait2 = next.wait2, next.wait1
					}
					if newLog, keep := log[next.st.at1][next.st.at2][next.st.open].merge(nextt, int(next.pressure)); !keep {
						continue
					} else {
						log[next.st.at1][next.st.at2][next.st.open] = newLog
					}
					q.push(nextt, next)
				}
			}
			continue
		} else if p.wait1 > 0 {
			if int(p.pressure)+(maxT-pt-int(p.wait1))*sum.FlowRates[p.st.at1]+(maxT-pt-minD1[p.st.at2]-1)*maxFlow[^(p.st.open|(1<<p.st.at1))] <= maxPressure {
				continue
			}
			p.wait1, p.wait2 = p.wait2, p.wait1
			p.st.at1, p.st.at2 = p.st.at2, p.st.at1
		} else {
			if int(p.pressure)+(maxT-pt-int(p.wait2))*sum.FlowRates[p.st.at2]+(maxT-pt-minD1[p.st.at1]-1)*maxFlow[^(p.st.open|(1<<p.st.at2))] <= maxPressure {
				continue
			}
		}
		for i := uint8(0); i < n; i++ {
			if p.st.open&(1<<i) != 0 || i == p.st.at1 {
				continue
			}
			t1 := pt + sum.Dist[p.st.at1][i] + 1
			t2 := pt + int(p.wait2)
			if t1 >= maxT && t2 >= maxT {
				continue
			}
			var (
				nextt int
				next  move2
			)
			switch {
			case t1 < t2:
				nextt, next = t1, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: 0, wait2: uint8(t2 - t1), pressure: p.pressure}
			case t1 == t2:
				nextt, next = t1, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: 0, wait2: 0, pressure: p.pressure}
			case t1 > t2:
				nextt, next = t2, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: uint8(t1 - t2), wait2: 0, pressure: p.pressure}
			}
			if next.wait1 == 0 && next.st.open&(1<<i) == 0 {
				next.st.open |= 1 << i
				next.pressure += uint16((maxT - t1) * sum.FlowRates[i])
			}
			if next.wait2 == 0 && next.st.open&(1<<p.st.at2) == 0 {
				next.st.open |= 1 << p.st.at2
				next.pressure += uint16((maxT - t2) * sum.FlowRates[p.st.at2])
			}
			if next.st.at1 > next.st.at2 {
				next.st.at1, next.st.at2 = next.st.at2, next.st.at1
				next.wait1, next.wait2 = next.wait2, next.wait1
			}
			if newLog, keep := log[next.st.at1][next.st.at2][next.st.open].merge(nextt, int(next.pressure)); !keep {
				continue
			} else {
				log[next.st.at1][next.st.at2][next.st.open] = newLog
			}
			q.push(nextt, next)
		}
	}

	return maxPressure
}

type move2 struct {
	st           state2
	wait1, wait2 uint8
	pressure     uint16
}

type state2 struct {
	at1  uint8
	at2  uint8
	open uint16
}

type logEntry2 struct {
	time     uint16
	pressure uint16
}

type logEntryList2 []logEntry2

func (log logEntryList2) merge(time, pressure int) (logEntryList2, bool) {
	for _, e := range log {
		if int(e.time) <= time && int(e.pressure) >= pressure {
			return nil, false
		}
	}
	newLog := log[:0]
	for _, e := range log {
		if int(e.time) < time || int(e.pressure) > pressure {
			newLog = append(newLog, e)
		}
	}
	newLog = append(newLog, logEntry2{time: uint16(time), pressure: uint16(pressure)})
	return newLog, true
}
