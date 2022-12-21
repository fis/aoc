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

// Package day19 solves AoC 2022 day 19.
package day19

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2022, 19, glue.LineSolver(glue.WithParser(parseBlueprint, solve)))
}

func parseBlueprint(line string) (bp blueprint, err error) {
	const format = (`Blueprint %d:` +
		` Each ore robot costs %d ore.` +
		` Each clay robot costs %d ore.` +
		` Each obsidian robot costs %d ore and %d clay.` +
		` Each geode robot costs %d ore and %d obsidian.`)
	var ignore int
	if _, err := fmt.Sscanf(line, format, &ignore, &bp.oreCostOre, &bp.clayCostOre, &bp.obsCostOre, &bp.obsCostClay, &bp.geoCostOre, &bp.geoCostObs); err != nil {
		return blueprint{}, err
	}
	return bp, nil
}

func solve(blueprints []blueprint) ([]string, error) {
	p1 := qualityLevels(blueprints, 24)
	p2 := fn.ProdF(blueprints[:3], func(bp blueprint) int { return maxGeodes(bp, 32) })
	return glue.Ints(p1, p2), nil
}

func qualityLevels(blueprints []blueprint, maxT int) (totalQL int) {
	for i, bp := range blueprints {
		totalQL += (i + 1) * maxGeodes(bp, maxT)
	}
	return totalQL
}

func maxGeodes(bp blueprint, maxT int) (maxGeo int) {
	capOreR := uint8(ix.Max(ix.Max(bp.clayCostOre, bp.obsCostOre), bp.geoCostOre))
	capClayR := uint8(bp.obsCostClay)
	capObsR := uint8(bp.geoCostObs)

	q := util.NewBucketQ[state](32)
	q.Push(0, state{0, 0, 0, 0, 1, 0, 0, 0})
	for q.Len() > 0 {
		pt, p := q.Pop()
		left := maxT - pt
		if g := lowerBound(p, left, bp); g > maxGeo {
			maxGeo = g
		} else if g := upperBound(p, left, bp); g <= maxGeo {
			continue
		}
		var next [4]struct {
			t int
			s state
		}
		nn := 0
		if p.oreR < capOreR {
			next[0].t = pt + 1 + ix.Max(ix.CeilDiv(bp.oreCostOre-int(p.ore), int(p.oreR)), 0)
			next[0].s = state{
				p.ore + p.oreR*uint8(next[0].t-pt) - uint8(bp.oreCostOre),
				p.clay + p.clayR*uint8(next[0].t-pt),
				p.obs + p.obsR*uint8(next[0].t-pt),
				p.geo + p.geoR*uint8(next[0].t-pt),
				p.oreR + 1, p.clayR, p.obsR, p.geoR,
			}
			nn++
		}
		if p.clayR < capClayR {
			next[nn].t = pt + 1 + ix.Max(ix.CeilDiv(bp.clayCostOre-int(p.ore), int(p.oreR)), 0)
			next[nn].s = state{
				p.ore + p.oreR*uint8(next[nn].t-pt) - uint8(bp.clayCostOre),
				p.clay + p.clayR*uint8(next[nn].t-pt),
				p.obs + p.obsR*uint8(next[nn].t-pt),
				p.geo + p.geoR*uint8(next[nn].t-pt),
				p.oreR, p.clayR + 1, p.obsR, p.geoR,
			}
			nn++
		}
		if p.clayR > 0 && p.obs < capObsR {
			next[nn].t = pt + 1 + ix.Max(ix.Max(ix.CeilDiv(bp.obsCostOre-int(p.ore), int(p.oreR)), ix.CeilDiv(bp.obsCostClay-int(p.clay), int(p.clayR))), 0)
			next[nn].s = state{
				p.ore + p.oreR*uint8(next[nn].t-pt) - uint8(bp.obsCostOre),
				p.clay + p.clayR*uint8(next[nn].t-pt) - uint8(bp.obsCostClay),
				p.obs + p.obsR*uint8(next[nn].t-pt),
				p.geo + p.geoR*uint8(next[nn].t-pt),
				p.oreR, p.clayR, p.obsR + 1, p.geoR,
			}
			nn++
		}
		if p.obsR > 0 {
			next[nn].t = pt + 1 + ix.Max(ix.Max(ix.CeilDiv(bp.geoCostOre-int(p.ore), int(p.oreR)), ix.CeilDiv(bp.geoCostObs-int(p.obs), int(p.obsR))), 0)
			next[nn].s = state{
				p.ore + p.oreR*uint8(next[nn].t-pt) - uint8(bp.geoCostOre),
				p.clay + p.clayR*uint8(next[nn].t-pt),
				p.obs + p.obsR*uint8(next[nn].t-pt) - uint8(bp.geoCostObs),
				p.geo + p.geoR*uint8(next[nn].t-pt),
				p.oreR, p.clayR, p.obsR, p.geoR + 1,
			}
			nn++
		}
		for ni := 0; ni < nn; ni++ {
			n := next[ni]
			if n.t > maxT {
				continue
			}
			if g := upperBound(n.s, maxT-n.t, bp); g > maxGeo {
				q.Push(n.t, n.s)
			}
		}
	}

	return maxGeo
}

func lowerBound(st state, left int, bp blueprint) int {
	// Use the strategy of "just build as many geode robots we can".
	ore, obs, geo := int(st.ore), int(st.obs), int(st.geo)
	oreR, obsR, geoR := int(st.oreR), int(st.obsR), int(st.geoR)
	for left > 0 {
		geo += geoR
		if ore >= bp.geoCostOre && obs >= bp.geoCostObs {
			geoR++
			ore -= bp.geoCostOre
			obs -= bp.geoCostObs
		}
		ore += oreR
		obs += obsR
		left--
	}
	return geo
}

func upperBound(st state, left int, bp blueprint) int {
	// Use the (unrealistic) strategy of "build one more of every type of robot every minute if resources allow".
	ore, clay, obs, geo := int(st.ore), int(st.clay), int(st.obs), int(st.geo)
	oreR, clayR, obsR, geoR := int(st.oreR), int(st.clayR), int(st.obsR), int(st.geoR)
	for left > 0 {
		ore0, clay0, obs0 := ore, clay, obs
		ore, clay, obs, geo = ore+oreR, clay+clayR, obs+obsR, geo+geoR
		if ore0 >= bp.oreCostOre {
			oreR++
		}
		if ore0 >= bp.clayCostOre {
			clayR++
		}
		if ore0 >= bp.obsCostOre && clay0 >= bp.obsCostClay {
			obsR++
		}
		if ore0 >= bp.geoCostOre && obs0 >= bp.geoCostObs {
			geoR++
		}
		left--
	}
	return geo
}

type state struct {
	ore, clay, obs, geo     uint8
	oreR, clayR, obsR, geoR uint8
}

type blueprint struct {
	oreCostOre  int
	clayCostOre int
	obsCostOre  int
	obsCostClay int
	geoCostOre  int
	geoCostObs  int
}
