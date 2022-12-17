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

func releasePressure2(sum valveSummary, maxT int) (maxPressure int) {
	if len(sum.flowRates) > maxValves {
		panic("too many valves")
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

	maxFlow := [65536]int{}
	for i := 0; i < 65536; i++ {
		for j := uint8(0); j < n; j++ {
			if i&(1<<j) != 0 {
				maxFlow[i] += sum.flowRates[j]
			}
		}
	}
	minD := maxT
	for i := uint8(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if sum.dist[i][j] < minD {
				minD = sum.dist[i][j]
			}
		}
	}

	log := [maxValves][maxValves][65536]logEntryList{}
outerLoop:
	for q.len() > 0 {
		pt, p := q.pop()
		for _, e := range log[p.st.at1][p.st.at2][p.st.open] {
			if (e.time <= pt && e.pressure > p.pressure) || (e.time < pt && e.pressure >= p.pressure) {
				continue outerLoop
			}
		}
		if p.pressure > maxPressure {
			maxPressure = p.pressure
		}
		if p.wait1 == 0 && p.wait2 == 0 {
			if p.pressure+(maxT-pt-minD-1)*maxFlow[^p.st.open] <= maxPressure {
				continue
			}
			for i := uint8(0); i < n; i++ {
				for j := uint8(0); j < n; j++ {
					if p.st.open&((1<<i)|(1<<j)) != 0 || i == p.st.at1 || i == p.st.at2 || j == p.st.at1 || j == p.st.at2 || i == j {
						continue
					}
					t1 := pt + sum.dist[p.st.at1][i] + 1
					t2 := pt + sum.dist[p.st.at2][j] + 1
					if t1 >= maxT && t2 >= maxT {
						continue
					}
					var (
						nextt int
						next  move2
					)
					switch {
					case t1 < t2:
						nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << i)}, wait1: 0, wait2: t2 - t1, pressure: p.pressure + (maxT-t1)*sum.flowRates[i]}
					case t1 == t2:
						nextt, next = t1, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << i) | (1 << j)}, wait1: 0, wait2: 0, pressure: p.pressure + (maxT-t1)*(sum.flowRates[i]+sum.flowRates[j])}
					case t1 > t2:
						nextt, next = t2, move2{st: state2{at1: i, at2: j, open: p.st.open | (1 << j)}, wait1: t1 - t2, wait2: 0, pressure: p.pressure + (maxT-t2)*sum.flowRates[j]}
					}
					if next.st.at1 > next.st.at2 {
						next.st.at1, next.st.at2 = next.st.at2, next.st.at1
						next.wait1, next.wait2 = next.wait2, next.wait1
					}
					if newLog, keep := log[next.st.at1][next.st.at2][next.st.open].merge(nextt, next.pressure); !keep {
						continue
					} else {
						log[next.st.at1][next.st.at2][next.st.open] = newLog
					}
					q.push(nextt, next)
				}
			}
			continue
		} else if p.wait1 > 0 {
			if p.pressure+(maxT-pt-p.wait1)*sum.flowRates[p.st.at1]+(maxT-pt-minD-1)*maxFlow[^(p.st.open|(1<<p.st.at1))] <= maxPressure {
				continue
			}
			p.wait1, p.wait2 = p.wait2, p.wait1
			p.st.at1, p.st.at2 = p.st.at2, p.st.at1
		} else {
			if p.pressure+(maxT-pt-p.wait2)*sum.flowRates[p.st.at2]+(maxT-pt-minD-1)*maxFlow[^(p.st.open|(1<<p.st.at2))] <= maxPressure {
				continue
			}
		}
		for i := uint8(0); i < n; i++ {
			if p.st.open&(1<<i) != 0 || i == p.st.at1 {
				continue
			}
			t1 := pt + sum.dist[p.st.at1][i] + 1
			t2 := pt + p.wait2
			if t1 >= maxT && t2 >= maxT {
				continue
			}
			var (
				nextt int
				next  move2
			)
			switch {
			case t1 < t2:
				nextt, next = t1, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: 0, wait2: t2 - t1, pressure: p.pressure}
			case t1 == t2:
				nextt, next = t1, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: 0, wait2: 0, pressure: p.pressure}
			case t1 > t2:
				nextt, next = t2, move2{st: state2{at1: i, at2: p.st.at2, open: p.st.open}, wait1: t1 - t2, wait2: 0, pressure: p.pressure}
			}
			if next.wait1 == 0 && next.st.open&(1<<i) == 0 {
				next.st.open |= 1 << i
				next.pressure += (maxT - t1) * sum.flowRates[i]
			}
			if next.wait2 == 0 && next.st.open&(1<<p.st.at2) == 0 {
				next.st.open |= 1 << p.st.at2
				next.pressure += (maxT - t2) * sum.flowRates[p.st.at2]
			}
			if next.st.at1 > next.st.at2 {
				next.st.at1, next.st.at2 = next.st.at2, next.st.at1
				next.wait1, next.wait2 = next.wait2, next.wait1
			}
			if newLog, keep := log[next.st.at1][next.st.at2][next.st.open].merge(nextt, next.pressure); !keep {
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
	wait1, wait2 int
	pressure     int
}

type state2 struct {
	at1  uint8
	at2  uint8
	open uint16
}
