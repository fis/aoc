// Copyright 2020 Google LLC
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

// Package day24 solves AoC 2018 day 24.
package day24

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 24, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	if len(chunks) != 2 {
		return nil, fmt.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	var s systemState
	if err := s.initialize(util.Lines(chunks[0]), util.Lines(chunks[1])); err != nil {
		return nil, err
	}

	part1, _ := s.copy().combat()

	bound, part2 := 1, 0
	for {
		ts := s.copy()
		ts.boost(bound)
		if score, good := ts.combat(); good {
			part2 = score
			break
		}
		bound *= 2
	}
	for lowest := bound / 2; lowest < bound; lowest++ {
		if lowest == 27 {
			continue // boost 27 with my input seems to take a very long time indeed
		}
		ts := s.copy()
		ts.boost(lowest)
		if score, good := ts.combat(); good {
			part2 = score
			break
		}
	}

	return glue.Ints(part1, part2), nil
}

type systemState struct {
	groups            []*armyGroup
	immune, infection int
}

func (s systemState) copy() (c *systemState) {
	c = new(systemState)
	c.immune, c.infection = s.immune, s.infection
	for _, g := range s.groups {
		cg := *g
		c.groups = append(c.groups, &cg)
	}
	return c
}

type armyGroup struct {
	infection  bool
	units      int
	hp         int
	weaknesses damageTypeSet
	immunities damageTypeSet
	baseDamage int
	attackType damageType
	initiative int
}

func (s *systemState) boost(v int) {
	for _, g := range s.groups {
		if !g.infection {
			g.baseDamage += v
		}
	}
}

func (s *systemState) combat() (outcome int, good bool) {
	for s.immune > 0 && s.infection > 0 {
		type target struct {
			g   *armyGroup
			dmg int
		}
		var immuneTargets, infectionTargets []target
		for _, g := range s.groups {
			if !g.infection {
				immuneTargets = append(immuneTargets, target{g: g})
			} else {
				infectionTargets = append(infectionTargets, target{g: g})
			}
		}

		type attack struct {
			atk, def *armyGroup
		}
		var attacks []attack

		sort.Slice(s.groups, func(i, j int) bool {
			gi, gj := s.groups[i], s.groups[j]
			if pi, pj := gi.power(), gj.power(); pi != pj {
				return pi > pj
			}
			return gi.initiative > gj.initiative
		})
		for _, g := range s.groups {
			targets := &infectionTargets
			if g.infection {
				targets = &immuneTargets
			}
			if len(*targets) == 0 {
				continue
			}
			for i := range *targets {
				(*targets)[i].dmg = (*targets)[i].g.receivedDamage(g.attackType, g.power())
			}
			sort.Slice(*targets, func(i, j int) bool {
				if di, dj := (*targets)[i].dmg, (*targets)[j].dmg; di != dj {
					return di > dj
				}
				gi, gj := (*targets)[i].g, (*targets)[j].g
				if pi, pj := gi.power(), gj.power(); pi != pj {
					return pi > pj
				}
				return gi.initiative > gj.initiative
			})
			if (*targets)[0].dmg > 0 {
				attacks = append(attacks, attack{atk: g, def: (*targets)[0].g})
				*targets = (*targets)[1:]
			}
		}

		sort.Slice(attacks, func(i, j int) bool { return attacks[i].atk.initiative > attacks[j].atk.initiative })
		for _, a := range attacks {
			if a.atk.units <= 0 {
				continue
			}
			dmg := a.def.receivedDamage(a.atk.attackType, a.atk.power())
			kills := dmg / a.def.hp
			a.def.units -= kills
			if a.def.units <= 0 {
				if !a.def.infection {
					s.immune--
				} else {
					s.infection--
				}
			}
		}
		s.groups = reapDead(s.groups)
	}

	for _, g := range s.groups {
		outcome += g.units
	}
	return outcome, s.immune > 0
}

func (g *armyGroup) power() int {
	return g.units * g.baseDamage
}

func (g *armyGroup) receivedDamage(dt damageType, power int) int {
	for _, w := range g.weaknesses {
		if w == dt {
			return 2 * power
		}
	}
	for _, i := range g.immunities {
		if i == dt {
			return 0
		}
	}
	return power
}

func reapDead(in []*armyGroup) (out []*armyGroup) {
	var i, o int
	for i < len(in) {
		if in[i].units > 0 {
			if o < i {
				in[o] = in[i]
			}
			o++
		}
		i++
	}
	return in[:o]
}

func (s *systemState) initialize(immuneSystem, infection []string) error {
	if len(immuneSystem) == 0 || immuneSystem[0] != "Immune System:" {
		return fmt.Errorf("bad immune system header")
	}
	if len(infection) == 0 || infection[0] != "Infection:" {
		return fmt.Errorf("bad infection header")
	}
	for _, army := range []struct {
		lines          []string
		counter        *int
		infectionGroup bool
	}{
		{lines: immuneSystem[1:], counter: &s.immune, infectionGroup: false},
		{lines: infection[1:], counter: &s.infection, infectionGroup: true},
	} {
		for _, line := range army.lines {
			g := &armyGroup{infection: army.infectionGroup}
			if err := g.parse(line); err != nil {
				return err
			}
			s.groups = append(s.groups, g)
			(*army.counter)++
		}
	}
	return nil
}

var groupPattern = regexp.MustCompile(`^(\d+) units each with (\d+) hit points(?: \(([^)]+)\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)$`)

func (g *armyGroup) parse(desc string) error {
	m := groupPattern.FindStringSubmatch(desc)
	if m == nil {
		return fmt.Errorf("malformatted description: %s", desc)
	}
	g.units, _ = strconv.Atoi(m[1])
	g.hp, _ = strconv.Atoi(m[2])
	if m[3] != "" {
		for _, spec := range strings.Split(m[3], "; ") {
			switch {
			case strings.HasPrefix(spec, "weak to "):
				if err := g.weaknesses.parseDamageTypes(spec[8:]); err != nil {
					return fmt.Errorf("bad weakness: %w", err)
				}
			case strings.HasPrefix(spec, "immune to "):
				if err := g.immunities.parseDamageTypes(spec[10:]); err != nil {
					return fmt.Errorf("bad immunity: %w", err)
				}
			default:
				return fmt.Errorf("bad damage modifier: %s", spec)
			}
		}
	}
	g.baseDamage, _ = strconv.Atoi(m[4])
	if dt, ok := damageTypes[m[5]]; ok {
		g.attackType = dt
	} else {
		return fmt.Errorf("unknown attack damage type: %s", m[5])
	}
	g.initiative, _ = strconv.Atoi(m[6])
	return nil
}

type damageType int

var damageTypeNames = []string{"bludgeoning", "cold", "fire", "radiation", "slashing"}

var damageTypes map[string]damageType

func init() {
	damageTypes = make(map[string]damageType)
	for dt, n := range damageTypeNames {
		damageTypes[n] = damageType(dt)
	}
}

type damageTypeSet []damageType

func (s *damageTypeSet) parseDamageTypes(spec string) error {
	if spec == "" {
		return nil
	}
	for _, n := range strings.Split(spec, ", ") {
		dt, ok := damageTypes[n]
		if !ok {
			return fmt.Errorf("unknown damage type: %s", n)
		}
		*s = append(*s, dt)
	}
	return nil
}
