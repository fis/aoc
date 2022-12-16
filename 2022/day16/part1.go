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

func releasePressure(sum valveSummary, maxT int) (maxPressure int) {
	if len(sum.flowRates) > 16 {
		panic("only <= 16 non-zero flow rate valves supported")
	}
	n := uint16(len(sum.flowRates))

	var q bucketQ[move]
	for i := uint16(0); i < n; i++ {
		t := sum.initDist[i] + 1
		q.push(t, move{st: state{at: i, open: 1 << i}, pressure: (maxT - t) * sum.flowRates[i]})
	}

	log := make(map[state][]logEntry)
	for q.len() > 0 {
		pt, p := q.pop()
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
			oldLog := log[next.st]
			for _, e := range oldLog {
				if e.time <= t && e.pressure >= next.pressure {
					continue moveLoop
				}
			}
			newLog := oldLog[:0]
			for _, e := range oldLog {
				if e.time < t || e.pressure > next.pressure {
					newLog = append(newLog, e)
				}
			}
			newLog = append(newLog, logEntry{time: t, pressure: next.pressure})
			log[next.st] = newLog
			q.push(t, next)
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
