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

const maxValves = 15 // using some fixed-size arrays for performance

func releasePressure(sum valveSummary, maxT int) (maxPressure int) {
	if len(sum.flowRates) > maxValves {
		panic("too many valves")
	}
	n := uint16(len(sum.flowRates))

	var q bucketQ[move]
	for i := uint16(0); i < n; i++ {
		t := sum.initDist[i] + 1
		q.push(t, move{st: state{at: i, open: 1 << i}, pressure: (maxT - t) * sum.flowRates[i]})
	}

	log := [maxValves][65536]logEntryList{}
	for q.len() > 0 {
		pt, p := q.pop()
		if p.pressure > maxPressure {
			maxPressure = p.pressure
		}
	moveLoop:
		for i := uint16(0); i < n; i++ {
			if i == p.st.at || p.st.open&(1<<i) != 0 {
				continue
			}
			t := pt + sum.dist[p.st.at][i] + 1
			if t >= 30 {
				continue
			}
			next := move{st: state{at: i, open: p.st.open | (1 << i)}, pressure: p.pressure + (maxT-t)*sum.flowRates[i]}
			if newLog, keep := log[next.st.at][next.st.open].merge(t, next.pressure); !keep {
				continue moveLoop
			} else {
				log[next.st.at][next.st.open] = newLog
			}
			q.push(t, next)
		}
	}

	return maxPressure
}

type move struct {
	st       state
	pressure int
}

type state struct {
	at   uint16
	open uint16
}

type logEntry struct {
	time     int
	pressure int
}

type logEntryList []logEntry

func (log logEntryList) merge(time, pressure int) (logEntryList, bool) {
	for _, e := range log {
		if e.time <= time && e.pressure >= pressure {
			return nil, false
		}
	}
	newLog := log[:0]
	for _, e := range log {
		if e.time < time || e.pressure > pressure {
			newLog = append(newLog, e)
		}
	}
	newLog = append(newLog, logEntry{time: time, pressure: pressure})
	return newLog, true
}
