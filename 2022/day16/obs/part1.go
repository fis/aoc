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

import "github.com/fis/aoc/2022/day16"

const maxValves = 15 // using some fixed-size arrays for performance

func releasePressure(sum day16.ValveSummary, maxT int) (maxPressure int) {
	if len(sum.FlowRates) > maxValves {
		panic("too many valves")
	}
	n := uint16(len(sum.FlowRates))

	var q bucketQ[move1]
	for i := uint16(0); i < n; i++ {
		t := sum.InitDist[i] + 1
		q.push(t, move1{st: state1{at: i, open: 1 << i}, pressure: (maxT - t) * sum.FlowRates[i]})
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
			t := pt + sum.Dist[p.st.at][i] + 1
			if t >= 30 {
				continue
			}
			next := move1{st: state1{at: i, open: p.st.open | (1 << i)}, pressure: p.pressure + (maxT-t)*sum.FlowRates[i]}
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

type move1 struct {
	st       state1
	pressure int
}

type state1 struct {
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
