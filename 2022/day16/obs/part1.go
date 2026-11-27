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
