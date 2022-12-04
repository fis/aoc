# AoC 2022 solution notes

## Background

It's that time of the year again. Back to the grind. See previous years'
introductions for more about what this is all about. In the meanwhile...

## [Day 1](https://adventofcode.com/2022/day/1): Calorie Counting

Typically soft landing. The Go code just sums and sorts to find the top three.

I've played around with extending the quickselect algorithm in the `util`
package to support the unordered partial sorting needed to find the _k_ lowest
(or highest) elements in arbitrary order in linear time. If I get around to
cleaning it up, I'll update this day's solution to use it.

### Burlesque

Part 1:

```
ln{""};;)ri)++>]
```

Part 2:

```
ln{""};;)ri)++<>3co-]++y
```

Combined:

```
1:               >]
C: ln{""};;)ri)++
2:               <>3co-]++y
```

## [Day 2](https://adventofcode.com/2022/day/2): Rock Paper Scissors

Oh, we're not doing
[Added Alliterative Appeal](https://tvtropes.org/pmwiki/pmwiki.php/Main/AddedAlliterativeAppeal)
(like [in 2020](2020-notes.md)) this year either? A shame.

Nothing interesting to say about today's Go solution. The Burlesque one does a
trick where it takes the A/B/C and X/Y/Z characters' code points modulo 3 and
treats them as the digits of a two-digit base-3 number, and uses that to just
look up the score from a map. This means both parts' solution is essentially the
same, except for the mapping.

### Burlesque

Part 1:

```
ln{XXRT[-{**3.%}m[3ug963174528XXj!!}ms
```

Part 2:

```
ln{XXRT[-{**3.%}m[3ug978123564XXj!!}ms
```

Combined:

```
1:                      963174528
C: ln{XXRT[-{**3.%}m[3ug         XXj!!}ms
2:                      978123564
```

## [Day 3](https://adventofcode.com/2022/day/3): Rucksack Reorganization

Oh, maybe we're alliterative on alternating days? I may be reading too much into
this.

The Go version uses an `[2]uint64` as the representation of a rucksack, with the
i'th bit set if an item with priority i exists in the corresponding compartment.
This makes finding the intersections real easy with bitwise operations.

Today's problem is a relatively good fit for Burlesque, what with its built-in
for finding set intersection (`IN`). Item priorities are computed as `b-38 mod
58`, which avoids needing a conditional.

### Burlesque

Part 1:

```
ln{sa2./cop^IN-]**38.-58.%}ms
```

Part 2:

```
ln3co{p^ININ-]**38.-58.%}ms
```

Combined:

```
1:       sa2./co
C: ln   {       p^IN  -]**38.-58.%}ms
2:   3co            IN
```

## [Day 4](https://adventofcode.com/2022/day/4): Camp Cleanup

Oh, great, now the titles are just actively teasing me.

Anyway, not much to say about the solutions. The Burlesque ones feel inelegant,
but they're short enough to let pass this time.

### Burlesque

Part 1:

```
ln{',;;{'-;;ri^pr@}m[Jp^~~j^p~~||}ms
```

Part 2:

```
ln{',;;{'-;;ri^pr@}MPINL[nz}ms
```

Combined:

```
1:                    m[Jp^~~j^p~~||
C: ln{',;;{'-;;ri^pr@}              }ms
2:                    MPINL[nz
```
